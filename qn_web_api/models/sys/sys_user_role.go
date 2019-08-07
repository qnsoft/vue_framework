package models

type SysUserRole struct {
	Id     int64 `xorm:"pk autoincr BIGINT(20)"`
	UserId int64 `xorm:"comment('用户ID') BIGINT(20)"`
	RoleId int64 `xorm:"comment('角色ID') BIGINT(20)"`
}
