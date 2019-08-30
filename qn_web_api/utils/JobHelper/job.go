package JobHelper

import (
	"fmt"
	"html/template"
	"os/exec"
	"qnsoft/qn_web_api/utils/WebHelper"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

var mailTpl *template.Template

func init() {
	mailTpl, _ = template.New("mail_tpl").Parse(`
	你好 {{.username}}，<br/>
	<p>以下是任务执行结果：</p>
	<p>
	任务ID：{{.task_id}}<br/>
	任务名称：{{.task_name}}<br/>       
	执行时间：{{.start_time}}<br />
	执行耗时：{{.process_time}}秒<br />
	执行状态：{{.status}}
	</p>
	<p>-------------以下是任务执行输出-------------</p>
	<p>{{.output}}</p>
	<p>
	--------------------------------------------<br />
	本邮件由系统自动发出，请勿回复<br />
	如果要取消邮件通知，请登录到系统进行设置<br />
	</p>
`)

}

type Job struct {
	id         int         // 计划ID
	logId      int64       // 日志记录ID
	name       string      // 任务名称
	task       *Task_model // 任务对象
	runFunc    func()      // 执行函数
	status     int         // 任务状态，大于0表示正在执行中
	Concurrent bool        // 同一个任务是否允许并行执行
}

/*
创建新任务到计划
@task任务对象
*/
func NewJobFromTask(task *Task_model, _func func()) (*Job, error) {
	if task.Id < 1 {
		return nil, fmt.Errorf("ToJob: 缺少id")
	}
	job := NewCommandJob(task.Id, task.Name, _func, "task.Command")
	job.task = task
	return job, nil
}

/*
创建新计划
@id 任务id
@name 任务名称
@command 任务命令
*/
func NewCommandJob(id int, name string, _func func(), command string) *Job {
	job := &Job{
		id:      id,
		name:    name,
		runFunc: _func,
	}
	// job.runFunc = func(timeout time.Duration) (string, string, error, bool) {
	// 	bufOut := new(bytes.Buffer)
	// 	bufErr := new(bytes.Buffer)
	// 	cmd := exec.Command("/bin/bash", "-c", command)
	// 	cmd.Stdout = bufOut
	// 	cmd.Stderr = bufErr
	// 	cmd.Start()
	// 	err, isTimeout := runCmdWithTimeout(cmd, timeout)
	// 	return bufOut.String(), bufErr.String(), err, isTimeout
	// }
	return job
}

/*
计划状态
*/
func (j *Job) Status() int {
	return j.status
}

/*
计划名称
*/
func (j *Job) GetName() string {
	return j.name
}

/*
计划id
*/
func (j *Job) GetId() int {
	return j.id
}

/*
计划日志id
*/
func (j *Job) GetLogId() int64 {
	return j.logId
}

/*
*运行任务
 */
func (j *Job) Run() {
	if !j.Concurrent && j.status > 0 {
		beego.Warn(fmt.Sprintf("任务[%d]上一次执行尚未结束，本次被忽略。", j.id))
		return
	}

	defer func() {
		if err := recover(); err != nil {
			beego.Error(err, "\n", string(debug.Stack()))
		}
	}()

	if workPool != nil {
		workPool <- true
		defer func() {
			<-workPool
		}()
	}
	//开始执行计划任务 //根据任务id,从数据库表中取出要执行的任务
	fmt.Println("开始执行任务：", strconv.Itoa(j.id)+"["+j.name+"]")
	j.status++
	defer func() {
		j.status--
	}()
	j.runFunc()
	//ZN_DF(j.id)
	//_start_time := php2go.URLEncode(date.FormatDate(time.Now().Add(time.Minute*-10), "yyyy-MM-dd HH:mm:ss"))

	/* _start_time := php2go.URLEncode("2019-08-14 10:48:03")
	_http_url := "http://api.vephp.com/order?vekey=V00002504Y26508322&span=1200&order_scene=2&tk_status=1&start_time=" + _start_time
	_req := Self_Get(_http_url, nil)
	json_all := make(map[string]interface{})
	json.Unmarshal([]byte(_req), &json_all)
	//fmt.Println(_req)
	//fmt.Println(json_all["data"])
	mjson, _ := json.Marshal(json_all["data"])
	mString := string(mjson)
	//	fmt.Println(json_all["error"])
	//	fmt.Println(json_all["data"])
	//if json_all["data"] != nil {
	//fmt.Println(json_all["data"])
	// _req_a := fmt.Sprintf("%s", json_all["data"])
	// fmt.Println(_req_a)
	var json_data []map[string]interface{}
	json.Unmarshal([]byte(mString), &json_data)
	if len(json_data) > 0 {
		for i, v := range json_data {
			_relation_id, _ := json.Marshal(v["relation_id"])
			fmt.Println(i, "buyuser---", string(_relation_id))
			_trade_id, _ := json.Marshal(v["trade_id"])
			fmt.Println(i, "sNo---", string(_trade_id))
			fmt.Println(i, "goods_title---", v["item_title"])
			_num_iid, _ := json.Marshal(v["num_iid"])
			fmt.Println(i, "goods_img---", string(_num_iid))
			fmt.Println(i, "goods_price---", v["price"])
			fmt.Println(i, "num---", v["item_num"])
			fmt.Println(i, "amount---", v["alipay_total_price"])
			fmt.Println(i, "comm_rate---", v["total_commission_rate"])
			fmt.Println(i, "comm_amount---", v["total_commission_fee"])
			fmt.Println(i, "status---", v["tk_status"])
			fmt.Println(i, "create_time---", v["create_time"])
			fmt.Println(i, "o_status---", v["tk_status"])
		}
	}
	fmt.Println("任务执行完毕!") */
}

/*
初始化任务
*/
func InitJobs(_list []*Task_model, _func func()) {
	//_list := make([]*Task_model, 0)
	//遍历信用卡智能代换任务列表
	for _, task := range _list {
		job, err := NewJobFromTask(task, _func)
		if err != nil {
			beego.Error("InitJobs:", err.Error())
			continue
		}
		AddJob(task.Spec, job)
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

/*
*任务实体
 */
type Task_model struct {
	//任务id
	Id int
	//任务名称
	Name string
	//执行时间表达式
	Spec string
}

/*
 get提交
 _map 提交参数
*/
func Self_Get(_http_url string, _map map[string]interface{}) string {
	_HeaderData := map[string]string{"Content-Type": "application/json"}
	_req := WebHelper.HttpGet(_http_url, _HeaderData, _map)
	return _req
}
