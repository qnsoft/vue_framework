package api

import (
	"encoding/json"
	"fmt"
	"qnsoft/qn_web_api/controllers/Token"
	"qnsoft/qn_web_api/models/sys"
	"strconv"
	"time"
	"zhenfangbian/web_api/utils/DbHelper"
	"zhenfangbian/web_api/utils/ErrorHelper"
)

/**
*系统用户信息实体
 */
type User_Controller struct {
	Token.BaseController
}

/*
系统账户登录
*/
func (this *User_Controller) Login() {
	var _rt_json interface{}
	var _FormData map[string]string
	_req := this.Ctx.Input.RequestBody
	json.Unmarshal([]byte(_req), &_FormData)
	//_Password := this.GetString("password")
	fmt.Println(_FormData["username"])
	_model := models.SysUser{Username: _FormData["username"], Password: _FormData["password"]}
	_results, err := DbHelper.MySqlDb().Get(&_model)
	ErrorHelper.CheckErr(err)
	if _results {
		_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "登录成功！", "data": &_model}
	} else {
		_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "登录失败！"}
	}
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
		_role_list := new(models.SysUserRole).List_RoleId(_model.UserId)
		_model_a := struct {
			*models.SysUser
			RoleidList []int64 `json:"roleid_list,omitempty"`
		}{_model, _role_list}
		_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "数据获取成功！", "user": &_model_a}
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
	_where := "user_id>0" //查询条件表达式
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
	_req := this.Ctx.Input.RequestBody
	json.Unmarshal([]byte(_req), &_FormData)
	_user_id, _ := strconv.ParseInt(_FormData["user_id"], 10, 0)
	_type := _FormData["type"]
	_roleid_list := _FormData["roleid_list"]
	_status, _ := strconv.Atoi(_FormData["status"])
	_model := models.SysUser{Username: _FormData["username"], Email: _FormData["email"], Mobile: _FormData["mobile"], Status: _status, CreateUserId: 0}
	switch _type {
	case "save":
		_model.CreateTime = time.Now()
		_count, err := new(models.SysUser).Get_Info_Add(&_model)
		ErrorHelper.CheckErr(err)
		if _count > 0 {
			new(models.SysUserRole).Add_user_role(_model.UserId, _roleid_list) //添加用户与角色关系
			_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "添加成功！", "role": &_model}
		} else {

			_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "添加失败！"}
		}
	case "update":
		//	_model.UserId = _user_id
		_count, err := new(models.SysUser).Get_Info_Update(_user_id, &_model)
		ErrorHelper.CheckErr(err)
		if _count > 0 || err == nil {
			new(models.SysUserRole).Update_user_role(_model.UserId, _roleid_list) //修改用户与角色关系
			_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "数据获取成功！", "role": &_model}
		} else {
			_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "数据获取失败！"}
		}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}
