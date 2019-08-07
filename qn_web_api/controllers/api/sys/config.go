package api

import (
	"encoding/json"
	"fmt"
	"qnsoft/qn_web_api/controllers/Token"
	"qnsoft/qn_web_api/models/sys"
	"strconv"
	"strings"
	"zhenfangbian/web_api/utils/DbHelper"
	"zhenfangbian/web_api/utils/ErrorHelper"
)

/**
*系统设置信息
 */
type Config_Controller struct {
	Token.BaseController
}

/*
菜单列表信息
*/
func (this *Config_Controller) List() {
	var _rt_json interface{}
	_pageIndex, _ := this.GetInt("pageIndex", 1)
	_pageSize, _ := this.GetInt("pageSize", 10)
	_where := []interface{}{"id>?", 0} //查询条件表达式
	_model := new(models.SysConfig)
	_totalCount, err1 := DbHelper.MySqlDb().Where(_where).Count(_model)
	ErrorHelper.CheckErr(err1)
	_rows, err2 := DbHelper.MySqlDb().Where(_where).Limit(_pageSize, (_pageIndex-1)*_pageSize).Rows(_model)
	ErrorHelper.CheckErr(err2)
	defer _rows.Close()
	_list := make([]*models.SysConfig, 0)
	_totalPage := 0
	if int(_totalCount)%_pageSize == 0 {
		_totalPage = int(_totalCount) / _pageSize
	} else {
		_totalPage = int(_totalCount)/_pageSize + 1
	}
	for _rows.Next() {
		_ = _rows.Scan(_model)
		_model_new := new(models.SysConfig)
		_model_new.Id = _model.Id
		_model_new.ParamKey = _model.ParamKey
		_model_new.ParamValue = _model.ParamValue
		_model_new.Status = _model.Status
		_model_new.Remark = _model.Remark
		_list = append(_list, _model_new)
	}
	if len(_list) > 0 {
		_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "获取数据列表成功！", "page": map[string]interface{}{"totalCount": _totalCount, "pageSize": _pageSize, "totalPage": _totalPage, "currPage": _pageIndex, "list": _list}}
	} else {
		_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "获取数据列表失败！"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
获取单条配置信息
*/
func (this *Config_Controller) Info() {
	var _rt_json interface{}
	_id, _ := this.GetInt64(":id", 1)
	_model := models.SysConfig{Id: _id}
	_results, err := DbHelper.MySqlDb().Get(&_model)
	ErrorHelper.CheckErr(err)
	if _results {
		_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "数据获取成功！", "config": &_model}
	} else {
		_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "数据获取失败！"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
修改配置信息
*/
func (this *Config_Controller) Edit() {
	var _rt_json interface{}
	var _FormData map[string]string
	_req := this.Ctx.Input.RequestBody
	json.Unmarshal([]byte(_req), &_FormData)
	fmt.Println(_req)
	_id, _ := strconv.ParseInt(_FormData["id"], 10, 0)
	_type := _FormData["type"]
	_model := models.SysConfig{Id: _id, ParamKey: _FormData["param_key"], ParamValue: _FormData["param_value"], Status: 0, Remark: _FormData["remark"]}
	switch _type {
	case "save":
		_count, err := DbHelper.MySqlDb().Insert(&_model)
		ErrorHelper.CheckErr(err)
		if _count > 0 {
			_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "添加成功！", "config": &_model}
		} else {
			_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "添加失败！"}
		}
	case "update":
		_count, err := DbHelper.MySqlDb().Id(_id).Update(&_model)
		ErrorHelper.CheckErr(err)
		if _count > 0 {
			_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "数据获取成功！", "config": &_model}
		} else {
			_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "数据获取失败！"}
		}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}

/*
删除设置信息
*/
func (this *Config_Controller) Delete() {
	var _rt_json interface{}
	var _FormData map[string]string
	_req := this.Ctx.Input.RequestBody
	json.Unmarshal([]byte(_req), &_FormData)
	_arry_ids := strings.Split(_FormData["ids"], ",")
	_model := new(models.SysConfig)
	_count, err := DbHelper.MySqlDb().In("id", _arry_ids).Delete(_model)
	ErrorHelper.CheckErr(err)
	if _count > 0 {
		_rt_json = map[string]interface{}{"code": 200, "msg": "success", "info": "删除成功！"}
	} else {
		_rt_json = map[string]interface{}{"code": 0, "msg": "fail", "info": "删除失败！"}
	}
	this.Data["json"] = _rt_json
	this.ServeJSON()
}
