package Token

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"qnsoft/qn_web_api/models/token"
	"qnsoft/qn_web_api/utils/ErrorHelper"
	//"../../utils/WebHelper"
	"github.com/astaxie/beego"
)

//连接tokentoken
type AccesstokenController struct {
	beego.Controller
}

/*
*用户token
 */
type Usertoken struct {
	Token      string
	Appid      string
	AppSecret  string
	Express_in int64
}

/**
*用Get方式获取token
 */
func (this *AccesstokenController) Access_Token() {
	/*验证appid 和 secret，下发token*/
	//form := CreateTokenForm{}
	var form CreateTokenForm
	if err := this.ParseForm(&form); err != nil {
		beego.Debug("errParseRegsiterForm:", err)
		this.Data["json"] = ErrInputData
		this.ServeJSON()
		return
	}
	_methord := strings.ToLower(this.Ctx.Request.Method)
	switch _methord {
	case "post":
		json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	case "get":
		form.Appid = this.GetString("Appid")
		form.AppSecret = this.GetString("AppSecret")
	default:
		json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	}
	/*验证appid 和 secret，下发token 要和数据库比对*/
	_Model_User := new(models.User)
	_isok, err := _Model_User.Check_Oauth(form.Appid, form.AppSecret)
	fmt.Println(_isok)
	ErrorHelper.CheckErr(err)
	if _isok {
		/*验证结束*/
		var T Usertoken
		T.Token, T.Express_in = Create_token(form.Appid, form.AppSecret)
		express_in := strconv.FormatInt(T.Express_in, 10)
		token_model, err := NewToken(&form, T.Token, express_in)
		if err != nil {
			beego.Error("NewUser:", err)
			this.Data["json"] = ErrSystem
			this.ServeJSON()
			return
		}
		beego.Debug("NewUser:", token_model)
		//token_model.Insert() //先不写入数据库
		T.Appid = form.Appid
		T.AppSecret = form.AppSecret
		this.Data["json"] = &T
		this.ServeJSON()
	} else {
		this.Data["json"] = ErrNoUser
		this.ServeJSON()
		return
	}
}

/*
*@token控制器基础实体
 */
type BaseController struct {
	beego.Controller
}

/*
*检测token
 */
func (this *BaseController) Check_Token() bool {
	var _is_rt = false
	sss := this.Ctx.Request.Header
	fmt.Println(sss)
	token := this.Ctx.Request.Header.Get("token") //接收头部token
	fmt.Println(token)
	if token == "" {
		this.Data["json"] = ErrPermission
		this.ServeJSON()
		_is_rt = false
	} else {
		_, err := Token_auth(token, "AppSecret")
		if err != nil {
			this.Data["json"] = ErrExpired
			this.ServeJSON()
			_is_rt = false
		} else {
			_is_rt = true
		}

	}
	return _is_rt
}
