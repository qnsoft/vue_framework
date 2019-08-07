package Sms

import (
	"encoding/json"
	"fmt"

	"qnsoft/qn_web_api/controllers/Token"

	"github.com/astaxie/beego"
)

/**
*信息实体
 */
type Sms_Controller struct {
	Token.BaseController
}

//短信发送
func (this *Sms_Controller) MsgSend() {
	_mobile := this.GetString("mobile")
	fmt.Println(_mobile)
	_content := this.GetString("content")
	fmt.Println(_content)
	//token检测
	if this.Check_Token() {
		var _sms_model Sms_Model
		//json.Unmarshal(this.Ctx.Input.RequestBody, &_sms_model)
		_sms_model.Account = beego.AppConfig.String("smsinfo::account")
		_sms_model.Password = beego.AppConfig.String("smsinfo::password")
		_sms_model.Mobile = _mobile
		_sms_model.Content = _content
		var _model HyWx
		_json := SendMsg(&_model, _sms_model)
		_obj := Sms_rtjson_Model{}
		err := json.Unmarshal([]byte(_json), &_obj)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"code": 0, "msg": "发送失败!", "smsid": 0}
		}
		this.Data["json"] = _obj
		this.ServeJSON()
	}
}
