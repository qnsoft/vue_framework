package api

import (
	"encoding/json"
	"fmt"
	"qnsoft/qn_web_api/controllers/Token"
	"qnsoft/qn_web_api/models/sys"
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
	// var _FormData map[string]string
	// _req := this.Ctx.Input.RequestBody
	// json.Unmarshal([]byte(_req), &_FormData)

	_model := models.SysUser{UserId: 3}
	_results, err := DbHelper.MySqlDb().Get(&_model)
	ErrorHelper.CheckErr(err)
	if _results {
		_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "数据获取成功！", "user": &_model}
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
	_where := []interface{}{"user_id>?", 0} //查询条件表达式
	_model := new(models.SysUser)
	_totalCount, err1 := DbHelper.MySqlDb().Where(_where).Count(_model)
	ErrorHelper.CheckErr(err1)
	_rows, err2 := DbHelper.MySqlDb().Where(_where).Limit(_pageSize, (_pageIndex-1)*_pageSize).Rows(_model)
	ErrorHelper.CheckErr(err2)
	defer _rows.Close()
	_list := make([]*models.SysUser, 0)
	_totalPage := 0
	if int(_totalCount)%_pageSize == 0 {
		_totalPage = int(_totalCount) / _pageSize
	} else {
		_totalPage = int(_totalCount)/_pageSize + 1
	}
	for _rows.Next() {
		_ = _rows.Scan(_model)
		_model_new := new(models.SysUser)
		_model_new.UserId = _model.UserId
		_model_new.Username = _model.Username
		_model_new.Email = _model.Email
		_model_new.Mobile = _model.Mobile
		_model_new.Status = _model.Status
		_model_new.CreateUserId = _model.CreateUserId
		_model_new.CreateTime = _model.CreateTime
		_list = append(_list, _model_new)
	}
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
