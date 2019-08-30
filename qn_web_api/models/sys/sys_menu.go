package models

import (
	"qnsoft/qn_web_api/utils/DbHelper"
	"qnsoft/qn_web_api/utils/ErrorHelper"
)

type SysMenu struct {
	MenuId   int64  `json:"menu_id,omitempty" xorm:"not null pk autoincr BIGINT(20)"`
	ParentId int64  `json:"parent_id,omitempty" xorm:"comment('父菜单ID，一级菜单为0') BIGINT(20)"`
	Name     string `json:"name" xorm:"comment('菜单名称') VARCHAR(50)"`
	Url      string `json:"url" xorm:"comment('菜单URL') VARCHAR(200)"`
	Perms    string `json:"perms" xorm:"comment('授权(多个用逗号分隔，如：user:list,user:create)') VARCHAR(500)"`
	Type     int    `json:"type" xorm:"comment('类型   0：目录   1：菜单   2：按钮') INT(11)"`
	Icon     string `json:"icon" xorm:"comment('菜单图标') VARCHAR(50)"`
	OrderNum int    `json:"order_num" xorm:"comment('排序') INT(11)"`
}

/*
添加单条信息
_model 信息实体
*/
func (this *SysMenu) Get_Info_Add(_model *SysMenu) (int64, error) {
	_count, err := DbHelper.MySqlDb().Insert(_model)
	ErrorHelper.CheckErr(err)
	return _count, err
}

/*
修改单条信息
_model 信息实体
*/
func (this *SysMenu) Get_Info_Update(_id int64, _cols string, _model *SysMenu) (int64, error) {
	_count, err := DbHelper.MySqlDb().Id(_id).Cols(_cols).Update(_model)
	ErrorHelper.CheckErr(err)
	return _count, err
}
