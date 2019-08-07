package jobs

import (
	"fmt"
	"os/exec"
	"time"
	"qnsoft/qn_web_api/models/shop"
	"qnsoft/qn_web_api/utils/DbHelper"
	"qnsoft/qn_web_api/utils/ErrorHelper"

	"github.com/astaxie/beego"
)

/*
初始化任务
*/
func InitJobs() {
	_model := new(shop.UserCardJobOrder)
	//从数据库调取任务列表
	rows, err := DbHelper.MySqlDb().SQL("select id,user_id,task_name,cron_spec,implement_time from lkt_user_card_job_order where status=1 and is_finish=0 order by implement_time asc").Rows(_model)
	ErrorHelper.CheckErr(err)
	defer rows.Close()
	_list := make([]*shop.UserCardJobOrder, 0)
	for rows.Next() {
		_ = rows.Scan(_model)
		_model_new := new(shop.UserCardJobOrder)
		_model_new.Id = _model.Id
		_model_new.UserId = _model.UserId
		_model_new.TaskName = _model.TaskName
		_model_new.CronSpec = _model.CronSpec
		_model_new.ImplementTime = _model.ImplementTime
		_list = append(_list, _model_new)
	}
	//遍历任务列表
	for _, task := range _list {
		job, err := NewJobFromTask(task)
		if err != nil {
			beego.Error("InitJobs:", err.Error())
			continue
		}
		AddJob(task.CronSpec, job)
	}
}

/*
*任务执行超时
 */
func runCmdWithTimeout(cmd *exec.Cmd, timeout time.Duration) (error, bool) {
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	var err error
	select {
	case <-time.After(timeout):
		beego.Warn(fmt.Sprintf("任务执行时间超过%d秒，进程将被强制杀掉: %d", int(timeout/time.Second), cmd.Process.Pid))
		go func() {
			<-done // 读出上面的goroutine数据，避免阻塞导致无法退出
		}()
		if err = cmd.Process.Kill(); err != nil {
			beego.Error(fmt.Sprintf("进程无法杀掉: %d, 错误信息: %s", cmd.Process.Pid, err))
		}
		return err, true
	case err = <-done:
		return err, false
	}
}
