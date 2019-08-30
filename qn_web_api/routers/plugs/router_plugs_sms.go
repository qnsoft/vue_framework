package plugs

import (
	"qnsoft/qn_web_api/plugs/Sms"

	"github.com/astaxie/beego"
)

/*
短信相关路由
*/
func Plugs_Sms() {
	//发送短信验证码
	beego.Router("/sms/MsgSend", &Sms.Sms_Controller{}, "post:MsgSend")
}
