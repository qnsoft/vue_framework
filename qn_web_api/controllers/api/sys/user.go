package api

import (
	"encoding/json"
	"qnsoft/qn_web_api/controllers/Token"
	"qnsoft/qn_web_api/models/sys"
	"qnsoft/qn_web_api/utils/DbHelper"
	"qnsoft/qn_web_api/utils/ErrorHelper"
	"qnsoft/qn_web_api/utils/PicHelper"
	"strconv"
	"strings"
	"time"

	"github.com/mojocn/base64Captcha"
)

/**
*系统用户信息实体
 */
type User_Controller struct {
	Token.BaseController
}

/* 全局变量用来存储验证码 */
var _VerCode string

/*
登录验证码
*/
func (this *User_Controller) VerifyCode() {
	idKeyD, base64stringD := PicHelper.Pic_verifycode_character(this.Ctx.ResponseWriter, this.Ctx.Request)
	//fmt.Println("原始值：", idKeyD+"===>>>", base64stringD)
	_json := map[string]interface{}{"code": 200, "id": idKeyD, "src": base64stringD, "msg": "success", "info": "成功获取验证码!"}
	this.Data["json"] = _json
	this.ServeJSON()

}

/*
系统账户登录
*/
func (this *User_Controller) Login() {
	var _rt_json interface{}
	var _FormData map[string]string
	_req := this.Ctx.Input.RequestBody
	json.Unmarshal([]byte(_req), &_FormData)
	_model := models.SysUser{Username: _FormData["username"], Password: _FormData["password"]}
	_idkeyd := strings.TrimSpace(_FormData["idkeyd"])
	_checkcode := strings.TrimSpace(_FormData["checkcode"])
	//比较图像验证码
	_very_code := base64Captcha.VerifyCaptcha(_idkeyd, _checkcode)
	_results, err := DbHelper.MySqlDb().Get(&_model)
	ErrorHelper.CheckErr(err)
	if _results && _very_code == true {
		_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "登录成功！", "user": &_model}
	} else if _results && _very_code == false {
		_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "验证码不正确！"}
	} else {
		_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "用户名或密码不正确！"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
修改管理员密码
*/
func (this *User_Controller) Password() {
	var _rt_json interface{}
	var _FormData map[string]string
	_req := this.Ctx.Input.RequestBody
	json.Unmarshal([]byte(_req), &_FormData)
	_arry_cols := make([]string, 0)
	_model := new(models.SysUser)
	_user_id, _ := strconv.ParseInt(_FormData["user_id"], 10, 0)
	if _FormData["password"] != "" {
		_model.Password = _FormData["newPassword"]
		_arry_cols = append(_arry_cols, "password")
	}
	_cols := strings.Join(_arry_cols, ",")
	_count, err := new(models.SysUser).Get_Info_Update(_user_id, _cols, _model)
	ErrorHelper.CheckErr(err)
	if _count > 0 || err == nil {
		_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "密码修改成功！"}
	} else {
		_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "密码修改失败！"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
系统账户退出登录
*/
func (this *User_Controller) Logout() {
	var _rt_json interface{}
	_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "成功退出！"}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
系统账户信息
*/
func (this *User_Controller) Info() {
	var _rt_json interface{}
	_user_id, _ := this.GetInt64(":user_id", 3)
	_results, _model := new(models.SysUser).Get_Info(_user_id)
	if _results {
		_role_list := new(models.SysUserRole).List_RoleId(_model.UserId) //获取与userid对应的角色
		_model_a := struct {
			*models.SysUser
			RoleidList []int64 `json:"roleid_list,omitempty"`
		}{_model, _role_list}
		_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "数据获取成功！", "user": _model_a}
	} else {
		_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "数据获取失败！"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
管理员信息列表
*/
func (this *User_Controller) List() {
	var _rt_json interface{}
	_pageIndex, _ := this.GetInt("pageIndex", 1)
	_pageSize, _ := this.GetInt("pageSize", 10)
	_keyWord := this.GetString("keyWord")
	_where := ""
	if len(_keyWord) > 0 {
		_where = "user_id>0 and username like '%" + _keyWord + "%' " //查询条件表达式
	} else {
		_where = "user_id>0" //查询条件表达式
	}
	_totalCount, _totalPage, _list := new(models.SysUser).Get_List_Info_Page(_where, _pageIndex, _pageSize)
	if len(_list) > 0 {
		_rt_json = map[string]interface{}{
			"code": 200,
			"msg":  "success",
			"info": "获取数据列表成功！",
			"page": map[string]interface{}{"totalCount": _totalCount, "pageSize": _pageSize, "totalPage": _totalPage, "currPage": _pageIndex, "list": _list}}
	} else {
		_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "获取数据列表失败！"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
修改角色信息
*/
func (this *User_Controller) Edit() {
	var _rt_json interface{}
	var _FormData map[string]string
	_arry_cols := make([]string, 0)
	_req := this.Ctx.Input.RequestBody
	json.Unmarshal([]byte(_req), &_FormData)
	_user_id, _ := strconv.ParseInt(_FormData["user_id"], 10, 0)
	_type := _FormData["type"]
	_roleid_list := _FormData["roleid_list"]
	_status, _ := strconv.Atoi(_FormData["status"])
	_model := new(models.SysUser)
	if _FormData["username"] != "" {
		_model.Username = _FormData["username"]
		_arry_cols = append(_arry_cols, "username")
	}
	if _FormData["email"] != "" {
		_model.Email = _FormData["email"]
		_arry_cols = append(_arry_cols, "email")
	}
	if _FormData["mobile"] != "" {
		_model.Mobile = _FormData["mobile"]
		_arry_cols = append(_arry_cols, "mobile")
	}
	if _model.Status > -1 {
		_model.Status = _status
		_arry_cols = append(_arry_cols, "status")
	}
	_model.CreateUserId = 0
	switch _type {
	case "save":
		_model.CreateTime = time.Now()
		_count, err := new(models.SysUser).Get_Info_Add(_model)
		ErrorHelper.CheckErr(err)
		if _count > 0 {

			new(models.SysUserRole).Add_user_role(_model.UserId, _roleid_list) //添加用户与角色关系
			_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "添加成功！", "role": &_model}
		} else {

			_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "添加失败！"}
		}
	case "update":
		//	_model.UserId = 0
		_cols := strings.Join(_arry_cols, ",")
		_count, err := new(models.SysUser).Get_Info_Update(_user_id, _cols, _model)
		ErrorHelper.CheckErr(err)
		if _count > 0 || err == nil {
			new(models.SysUserRole).Update_user_role(_user_id, _roleid_list) //修改用户与角色关系
			_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "修改成功！", "role": &_model}
		} else {
			_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "修改失败！"}
		}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
删除用户信息
*/
func (this *User_Controller) Delete() {
	var _rt_json interface{}
	var _FormData map[string]string
	_req := this.Ctx.Input.RequestBody
	json.Unmarshal([]byte(_req), &_FormData)
	_arry_ids := strings.Split(_FormData["ids"], ",")
	_model := new(models.SysUser)
	_count, err := DbHelper.MySqlDb().In("user_id", _arry_ids).Delete(_model)
	ErrorHelper.CheckErr(err)
	if _count > 0 {
		_model1 := new(models.SysUserRole)
		_, err := DbHelper.MySqlDb().In("user_id", _arry_ids).Delete(_model1)
		ErrorHelper.CheckErr(err)
		_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "删除成功！"}
	} else {
		_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "删除失败！"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}
