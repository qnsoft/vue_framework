package api

import (
	"encoding/json"
	"qnsoft/qn_web_api/controllers/Token"
	"qnsoft/qn_web_api/models/sys"
	"strconv"
	"strings"
	"zhenfangbian/web_api/utils/DbHelper"
	"zhenfangbian/web_api/utils/ErrorHelper"
)

/**
*菜单信息实体
 */
type Menu_Controller struct {
	Token.BaseController
}

/*
系统后台菜单信息
*/
func (this *Menu_Controller) Nav() {
	var _rt_json interface{}
	// var _FormData map[string]string
	// _req := this.Ctx.Input.RequestBody
	// json.Unmarshal([]byte(_req), &_FormData)

	_list := bll_meun_listNav(0)
	if len(_list) > 0 {
		_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "获取数据成功！", "menuList": _list}
	} else {
		_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "获取数据失败！"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
菜单列表信息
*/
func (this *Menu_Controller) List() {
	var _rt_json interface{}
	// var _FormData map[string]string
	// _req := this.Ctx.Input.RequestBody
	// json.Unmarshal([]byte(_req), &_FormData)

	_list := bll_meun_list(0)
	if len(_list) > 0 {
		_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "获取数据成功！", "menuList": _list}
	} else {
		_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "获取数据失败！"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
菜单选择
*/
func (this *Menu_Controller) Select() {
	var _rt_json interface{}
	// var _FormData map[string]string
	// _req := this.Ctx.Input.RequestBody
	// json.Unmarshal([]byte(_req), &_FormData)
	_list := bll_meun_list(0)
	if len(_list) > 0 {
		_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "获取数据成功！", "menuList": _list}
	} else {
		_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "获取数据失败！"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
获取单条菜单信息
*/
func (this *Menu_Controller) Info() {
	var _rt_json interface{}
	_id, _ := this.GetInt64(":menu_id", 1)
	_model_menu := bll_get_model(_id)

	if _model_menu != nil {
		_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "数据获取成功！", "menu": &_model_menu}
	} else {
		_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "数据获取失败！"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
修改角色信息
*/
func (this *Menu_Controller) Edit() {
	var _rt_json interface{}
	var _FormData map[string]string
	_arry_cols := make([]string, 0)
	_req := this.Ctx.Input.RequestBody
	json.Unmarshal([]byte(_req), &_FormData)
	_menu_id, _ := strconv.ParseInt(_FormData["menu_id"], 10, 0)
	_type_edit := _FormData["type_edit"]

	//fmt.Println(_arry_cols, _menu_id)
	// _roleid_list := _FormData["roleid_list"]
	// _status, _ := strconv.Atoi(_FormData["status"])
	_model := new(models.SysMenu)
	if _FormData["parent_id"] != "" {
		_parent_id, _ := strconv.ParseInt(_FormData["parent_id"], 10, 0)
		_model.ParentId = _parent_id
		_arry_cols = append(_arry_cols, "parent_id")
	}
	if _FormData["name"] != "" {
		_model.Name = _FormData["name"]
		_arry_cols = append(_arry_cols, "name")
	}
	if _FormData["url"] != "" {
		_model.Url = _FormData["url"]
		_arry_cols = append(_arry_cols, "url")
	}
	if _FormData["perms"] != "" {
		_model.Perms = _FormData["perms"]
		_arry_cols = append(_arry_cols, "perms")
	}
	if _FormData["type"] != "" {
		_type, _ := strconv.Atoi(_FormData["type"])
		_model.Type = _type
		_arry_cols = append(_arry_cols, "type")
	}
	if _FormData["icon"] != "" {
		_model.Icon = _FormData["icon"]
		_arry_cols = append(_arry_cols, "icon")
	}
	if _FormData["order_num"] != "" {
		_order_num, _ := strconv.Atoi(_FormData["order_num"])
		_model.OrderNum = _order_num
		_arry_cols = append(_arry_cols, "order_num")
	}
	switch _type_edit {
	case "save":
		_model.MenuId = 0
		//_model.CreateTime = time.Now()
		_count, err := new(models.SysMenu).Get_Info_Add(_model)
		ErrorHelper.CheckErr(err)
		if _count > 0 {
			_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "添加成功！", "menu": &_model}
		} else {

			_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "添加失败！"}
		}
	case "update":
		_model.MenuId = 0
		_cols := strings.Join(_arry_cols, ",")
		_count, err := new(models.SysMenu).Get_Info_Update(_menu_id, _cols, _model)
		ErrorHelper.CheckErr(err)
		if _count > 0 || err == nil {
			_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "修改成功！", "menu": &_model}
		} else {
			_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "修改失败！"}
		}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
删除菜单信息
*/
func (this *Menu_Controller) Delete() {
	var _rt_json interface{}
	var _FormData map[string]string
	_req := this.Ctx.Input.RequestBody
	json.Unmarshal([]byte(_req), &_FormData)
	_arry_ids := strings.Split(_FormData["ids"], ",")
	_model := new(models.SysMenu)
	_count, err := DbHelper.MySqlDb().In("menu_id", _arry_ids).Delete(_model)
	ErrorHelper.CheckErr(err)
	if _count > 0 {
		_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "删除成功！"}
	} else {
		_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "删除失败！"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
获取单条菜单信息
*/
func bll_get_model(_menu_id int64) *models.SysMenu {
	_model := models.SysMenu{MenuId: _menu_id}
	_, err := DbHelper.MySqlDb().Get(&_model)
	ErrorHelper.CheckErr(err)
	return &_model
}

/* 获取后台导航菜单列表 */
func bll_meun_listNav(_pid int64) []*Menu_Model {
	_list := make([]*Menu_Model, 0)
	_model := new(models.SysMenu)
	str_pid := strconv.FormatInt(_pid, 10)
	rows, err := DbHelper.MySqlDb().Where(" type<2 and parent_id=" + str_pid).Asc("order_num").Rows(_model)
	ErrorHelper.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		_ = rows.Scan(_model)
		_model_new := new(Menu_Model)
		_model_new.MenuId = _model.MenuId
		_model_new.ParentId = _model.ParentId
		_model_new.Name = _model.Name
		_model_new.Url = _model.Url
		_model_new.Perms = _model.Perms
		_model_new.Type = _model.Type
		_model_new.Icon = _model.Icon
		_model_new.OrderNum = _model.OrderNum
		_model_new.List = bll_meun_listNav(_model.MenuId)
		if _model_new.Type < 2 {
			_list = append(_list, _model_new)
		}
	}
	return _list
}

/* 获取菜单列表 */
func bll_meun_list(_pid int64) []*Menu_Model {
	_list := make([]*Menu_Model, 0)
	_model := new(models.SysMenu)
	rows, err := DbHelper.MySqlDb().Where(" menu_id>0 ").Asc("order_num").Rows(_model)
	ErrorHelper.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		_ = rows.Scan(_model)
		_model_new := new(Menu_Model)
		_model_new.MenuId = _model.MenuId
		_model_new.ParentId = _model.ParentId
		_model_new.ParentName = bll_get_model(_model.ParentId).Name
		_model_new.Name = _model.Name
		_model_new.Url = _model.Url
		_model_new.Perms = _model.Perms
		_model_new.Type = _model.Type
		_model_new.Icon = _model.Icon
		_model_new.OrderNum = _model.OrderNum
		_list = append(_list, _model_new)
	}
	return _list
}

/* 系统菜单实体信息 */
type Menu_Model struct {
	models.SysMenu
	ParentName string        `json:"parentname"`
	Open       string        `json:"open"`
	List       []*Menu_Model `json:"list"`
}
