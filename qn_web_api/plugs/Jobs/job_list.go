package Jobs

import (
	"fmt"
	"qnsoft/qn_web_api/utils/ErrorHelper"
	"qnsoft/qn_web_api/utils/JobHelper"
	"time"
)

/*
测试计划
*/
func Test_joblist() {
	//开始采两小时前集所有订单
	_fasks := make([]*JobHelper.Task_model, 0)
	_model := JobHelper.Task_model{Id: 1001, Name: "测试计划", Spec: "0 0/2 1-23 * * ?"} //每天早上1点到23点 每2分钟执行一次
	_fasks = append(_fasks, &_model)
	JobHelper.InitJobs(_fasks, job_qudao_all_Yesterday) //开始采集昨天全部订单
}

/*
开始采集两小时前全部订单(每天全天执行，采集前两小时前的所有订单，顺便将漏采的补录到数据库)
*/
func job_qudao_all_Yesterday() {
	fmt.Println("执行订单【?】成功！", time.Now)
	ErrorHelper.LogInfo("执行订单【?】成功！", time.Now)
}
