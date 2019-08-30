package models

import (
	"strconv"
	"strings"
	"qnsoft/qn_web_api/utils/DbHelper"
	"qnsoft/qn_web_api/utils/ErrorHelper"
)

/*
用户与角色关系实体信息
*/
type SysUserRole struct {
	Id     int64 `json:"id,omitempty" xorm:"pk autoincr BIGINT(20)"`
	UserId int64 `json:"user_id,omitempty" xorm:"comment('用户ID') BIGINT(20)"`
	RoleId int64 `json:"role_id,omitempty" xorm:"comment('角色ID') BIGINT(20)"`
}

/*
获取用户与角色关系列表
*/
func (this *SysUserRole) List(_UserId int64) []*SysUserRole {
	_model := new(SysUserRole)
	_rows, err := DbHelper.MySqlDb().Where(" user_id=3 ").Rows(_model)
	ErrorHelper.CheckErr(err)
	defer _rows.Close()
	_list := make([]*SysUserRole, 0)
	for _rows.Next() {
		_ = _rows.Scan(_model)
		_model_new := new(SysUserRole)
		_model_new.Id = _model.Id
		_model_new.UserId = _model.UserId
		_model_new.RoleId = _model.RoleId
		_list = append(_list, _model_new)
	}
	return _list
}

/*
获取用户与角色关系列表
*/
func (this *SysUserRole) List_RoleId(_UserId int64) []int64 {
	_model := new(SysUserRole)
	_rows, err := DbHelper.MySqlDb().Where(" user_id=? ", _UserId).Rows(_model)
	ErrorHelper.CheckErr(err)
	defer _rows.Close()
	_arry := make([]int64, 0)
	for _rows.Next() {
		_ = _rows.Scan(_model)
		_model_new := new(SysUserRole)
		_model_new.RoleId = _model.RoleId
		_arry = append(_arry, _model_new.RoleId)
	}
	return _arry
}

/*
添加用户角色
*/
func (this *SysUserRole) Add_user_role(_user_id int64, _role_ids string) {
	_arry_ids := strings.Split(_role_ids, ",")
	for _, v := range _arry_ids {
		_role_id, _ := strconv.ParseInt(v, 10, 0)
		_model := SysUserRole{UserId: _user_id, RoleId: _role_id}
		_, err := DbHelper.MySqlDb().Insert(&_model)
		ErrorHelper.CheckErr(err)
	}

}

/*
修改用户角色
*/
func (this *SysUserRole) Update_user_role(_user_id int64, _role_ids string) {
	_arry_ids := strings.Split(_role_ids, ",")
	_model := new(SysUserRole)
	_, err := DbHelper.MySqlDb().In("user_id", _user_id).Delete(_model)
	ErrorHelper.CheckErr(err)
	for _, v := range _arry_ids {
		_role_id, _ := strconv.ParseInt(v, 10, 0)
		_model1 := SysUserRole{UserId: _user_id, RoleId: _role_id}
		_, err := DbHelper.MySqlDb().Insert(&_model1)
		ErrorHelper.CheckErr(err)
	}
}
