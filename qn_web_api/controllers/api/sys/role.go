package api

import (
	"encoding/json"
	"qnsoft/qn_web_api/controllers/Token"
	"qnsoft/qn_web_api/models/sys"
	"strconv"
	"strings"
	"time"
	"zhenfangbian/web_api/utils/DbHelper"
	"zhenfangbian/web_api/utils/ErrorHelper"
)

/**
*系统用户信息实体
 */
type Role_Controller struct {
	Token.BaseController
}

/*
角色信息列表
*/
func (this *Role_Controller) List() {
	var _rt_json interface{}
	_pageIndex, _ := this.GetInt("pageIndex", 1)
	_pageSize, _ := this.GetInt("pageSize", 10)
	_keyWord := this.GetString("keyWord")
	_where := ""
	if len(_keyWord) > 0 {
		_where = "role_id>0 and role_name like '%" + _keyWord + "%' " //查询条件表达式
	} else {
		_where = "role_id>0" //查询条件表达式
	}
	_model := new(models.SysRole)
	_totalCount, err1 := DbHelper.MySqlDb().Where(_where).Count(_model)
	ErrorHelper.CheckErr(err1)
	_rows, err2 := DbHelper.MySqlDb().Where(_where).Limit(_pageSize, (_pageIndex-1)*_pageSize).Rows(_model)
	ErrorHelper.CheckErr(err2)
	defer _rows.Close()
	_list := make([]*models.SysRole, 0)
	_totalPage := 0
	if int(_totalCount)%_pageSize == 0 {
		_totalPage = int(_totalCount) / _pageSize
	} else {
		_totalPage = int(_totalCount)/_pageSize + 1
	}
	for _rows.Next() {
		_ = _rows.Scan(_model)
		_model_new := new(models.SysRole)
		_model_new.RoleId = _model.RoleId
		_model_new.RoleName = _model.RoleName
		_model_new.Remark = _model.Remark
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

/*
选择角色
*/
func (this *Role_Controller) Select() {
	var _rt_json interface{}
	_keyWord := this.GetString("keyWord")
	_where := ""
	if len(_keyWord) > 0 {
		_where = "role_id>0 and role_name like '%" + _keyWord + "%' " //查询条件表达式
	} else {
		_where = "role_id>0" //查询条件表达式
	}
	_model := new(models.SysRole)

	_rows, err2 := DbHelper.MySqlDb().Where(_where).Rows(_model)
	ErrorHelper.CheckErr(err2)
	defer _rows.Close()
	_list := make([]*models.SysRole, 0)
	for _rows.Next() {
		_ = _rows.Scan(_model)
		_model_new := new(models.SysRole)
		_model_new.RoleId = _model.RoleId
		_model_new.RoleName = _model.RoleName
		_model_new.Remark = _model.Remark
		_model_new.CreateUserId = _model.CreateUserId
		_model_new.CreateTime = _model.CreateTime
		_list = append(_list, _model_new)
	}
	if len(_list) > 0 {
		_rt_json = map[string]interface{}{
			"code": 200,
			"msg":  "success",
			"info": "获取数据列表成功！",
			"list": _list}
	} else {
		_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "获取数据列表失败！"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
获取单条角色信息
*/
func (this *Role_Controller) Info() {
	var _rt_json interface{}
	_id, _ := this.GetInt64(":role_id", 1)
	_model := models.SysRole{RoleId: _id}
	_results, err := DbHelper.MySqlDb().Get(&_model)
	ErrorHelper.CheckErr(err)
	_model_menu := Role_Menu_List_Model{SysRole: models.SysRole{RoleId: _model.RoleId, RoleName: _model.RoleName, Remark: _model.Remark, CreateUserId: _model.CreateUserId, CreateTime: _model.CreateTime}, MenuIdlist: get_role_menu(_model.RoleId)}
	if _results {
		_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "数据获取成功！", "role": &_model_menu}
	} else {
		_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "数据获取失败！"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
获取角色对应的菜单列表信息
*/
func get_role_menu(_role_id int64) string {
	_model := new(models.SysRoleMenu)
	_rows, err := DbHelper.MySqlDb().Where("role_id=?", _role_id).Rows(_model)
	ErrorHelper.CheckErr(err)
	defer _rows.Close()
	_menu_ids := "["
	for _rows.Next() {
		_ = _rows.Scan(_model)
		_model_new := new(models.SysRoleMenu)
		_model_new.MenuId = _model.MenuId
		_menu_ids += strconv.FormatInt(_model_new.MenuId, 10) + ","
	}
	_menu_ids += "]"
	return strings.Replace(_menu_ids, ",]", "]", -1)
}

/*
修改角色信息
*/
func (this *Role_Controller) Edit() {
	var _rt_json interface{}
	var _FormData map[string]string
	_req := this.Ctx.Input.RequestBody
	json.Unmarshal([]byte(_req), &_FormData)
	//fmt.Println(_req)
	_id, _ := strconv.ParseInt(_FormData["role_id"], 10, 0)
	_type := _FormData["type"]
	_menu_idlist := _FormData["menu_idlist"]
	_model := models.SysRole{RoleName: _FormData["role_name"], Remark: _FormData["remark"]}
	switch _type {
	case "save":
		_model.CreateTime = time.Now()
		_count, err := DbHelper.MySqlDb().Insert(&_model)
		ErrorHelper.CheckErr(err)
		if _count > 0 {
			Add_role_menu(_model.RoleId, _menu_idlist)
			_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "添加成功！", "role": &_model}
		} else {

			_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "添加失败！"}
		}
	case "update":
		_model.RoleId = _id
		_count, err := DbHelper.MySqlDb().Id(_id).Update(&_model)
		ErrorHelper.CheckErr(err)
		if _count > 0 || err == nil {
			Update_role_menu(_id, _menu_idlist)
			_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "数据获取成功！", "role": &_model}
		} else {
			_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "数据获取失败！"}
		}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
删除角色信息
*/
func (this *Role_Controller) Delete() {
	var _rt_json interface{}
	var _FormData map[string]string
	_req := this.Ctx.Input.RequestBody
	json.Unmarshal([]byte(_req), &_FormData)
	_arry_ids := strings.Split(_FormData["ids"], ",")
	_model := new(models.SysRole)
	_count, err := DbHelper.MySqlDb().In("role_id", _arry_ids).Delete(_model)
	ErrorHelper.CheckErr(err)
	if _count > 0 {
		_model1 := new(models.SysRoleMenu)
		_, err := DbHelper.MySqlDb().In("role_id", _arry_ids).Delete(_model1)
		ErrorHelper.CheckErr(err)
		_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "删除成功！"}
	} else {
		_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "删除失败！"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/* 角色菜单扩展 */
type Role_Menu_List_Model struct {
	models.SysRole
	MenuIdlist string `json:"menu_idlist,omitempty"`
}

/*
添加角色菜单
*/
func Add_role_menu(_role_id int64, _menu_ids string) {
	_arry_ids := strings.Split(_menu_ids, ",")
	for _, v := range _arry_ids {
		_menu_id, _ := strconv.ParseInt(v, 10, 0)
		_model := models.SysRoleMenu{RoleId: _role_id, MenuId: _menu_id}
		_, err := DbHelper.MySqlDb().Insert(&_model)
		ErrorHelper.CheckErr(err)
	}

}

/*
修改角色菜单
*/
func Update_role_menu(_role_id int64, _menu_ids string) {
	_arry_ids := strings.Split(_menu_ids, ",")
	_model := new(models.SysRoleMenu)
	_, err := DbHelper.MySqlDb().In("role_id", _role_id).Delete(_model)
	ErrorHelper.CheckErr(err)
	for _, v := range _arry_ids {
		_menu_id, _ := strconv.ParseInt(v, 10, 0)
		_model1 := models.SysRoleMenu{RoleId: _role_id, MenuId: _menu_id}
		_, err := DbHelper.MySqlDb().Insert(&_model1)
		ErrorHelper.CheckErr(err)
	}
}
