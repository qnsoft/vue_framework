package models

type SysMenu struct {
	MenuId   int64  `json:"menu_id,omitempty" xorm:"not null pk autoincr BIGINT(20)"`
	ParentId int64  `json:"parent_id,omitempty" xorm:"comment('父菜单ID，一级菜单为0') BIGINT(20)"`
	Name     string `json:"name,omitempty" xorm:"comment('菜单名称') VARCHAR(50)"`
	Url      string `json:"url,omitempty" xorm:"comment('菜单URL') VARCHAR(200)"`
	Perms    string `json:"perms,omitempty" xorm:"comment('授权(多个用逗号分隔，如：user:list,user:create)') VARCHAR(500)"`
	Type     int    `json:"type,omitempty" xorm:"comment('类型   0：目录   1：菜单   2：按钮') INT(11)"`
	Icon     string `json:"icon,omitempty" xorm:"comment('菜单图标') VARCHAR(50)"`
	OrderNum int    `json:"order_num,omitempty" xorm:"comment('排序') INT(11)"`
}
