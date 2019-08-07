package models

/*
角色与菜单信息关系表
*/
type SysRoleMenu struct {
	Id     int64 `json:"id,omitempty" xorm:"pk autoincr BIGINT(20)"`
	RoleId int64 `json:"role_id,omitempty" xorm:"comment('角色ID') BIGINT(20)"`
	MenuId int64 `json:"menu_id,omitempty" xorm:"comment('菜单ID') BIGINT(20)"`
}
