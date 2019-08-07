package models

import (
	"time"
)

/* 角色实体对象 */
type SysRole struct {
	RoleId       int64     `json:"role_id,omitempty" xorm:"not null pk autoincr BIGINT(20)"`
	RoleName     string    `json:"role_name,omitempty" xorm:"comment('角色名称') VARCHAR(100)"`
	Remark       string    `json:"remark,omitempty" xorm:"comment('备注') VARCHAR(100)"`
	CreateUserId int64     `json:"create_user_id,omitempty" xorm:"comment('创建者ID') BIGINT(20)"`
	CreateTime   time.Time `json:"create_time,omitempty" xorm:"comment('创建时间') DATETIME"`
}
