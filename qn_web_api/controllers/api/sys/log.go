package api

import (
	"qnsoft/qn_web_api/controllers/Token"
	"qnsoft/qn_web_api/models/sys"
	"qnsoft/qn_web_api/utils/DbHelper"
	"qnsoft/qn_web_api/utils/ErrorHelper"
)

/**
*菜单信息实体
 */
type Log_Controller struct {
	Token.BaseController
}

/*
菜单列表信息
*/
func (this *Log_Controller) List() {
	var _rt_json interface{}
	_pageIndex, _ := this.GetInt("pageIndex", 1)
	_pageSize, _ := this.GetInt("pageSize", 10)
	_where := []interface{}{"id>?", 0} //查询条件表达式
	_model := new(models.SysLog)
	_totalCount, err1 := DbHelper.MySqlDb().Where(_where).Count(_model)
	ErrorHelper.CheckErr(err1)
	_rows, err2 := DbHelper.MySqlDb().Where(_where).Limit(_pageSize, (_pageIndex-1)*_pageSize).Rows(_model)
	ErrorHelper.CheckErr(err2)
	defer _rows.Close()
	_list := make([]*models.SysLog, 0)
	_totalPage := 0
	if int(_totalCount)%_pageSize == 0 {
		_totalPage = int(_totalCount) / _pageSize
	} else {
		_totalPage = int(_totalCount)/_pageSize + 1
	}
	for _rows.Next() {
		_ = _rows.Scan(_model)
		_model_new := new(models.SysLog)
		_model_new.Id = _model.Id
		_model_new.Username = _model.Username
		_model_new.Operation = _model.Operation
		_model_new.Method = _model.Method
		_model_new.Params = _model.Params
		_model_new.Time = _model.Time
		_model_new.Ip = _model.Ip
		_model_new.CreateDate = _model.CreateDate
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
