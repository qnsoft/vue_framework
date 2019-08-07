package models

import (
	"time"
)

type SysUser struct {
	UserId       int64     `json:"user_id,omitempty" xorm:"not null pk autoincr BIGINT(20)"`
	Username     string    `json:"username,omitempty" xorm:"not null comment('用户名') unique VARCHAR(50)"`
	Password     string    `json:"password,omitempty" xorm:"comment('密码') VARCHAR(100)"`
	Salt         string    `json:"salt,omitempty" xorm:"comment('盐') VARCHAR(20)"`
	Email        string    `json:"email,omitempty" xorm:"comment('邮箱') VARCHAR(100)"`
	Mobile       string    `json:"mobile,omitempty" xorm:"comment('手机号') VARCHAR(100)"`
	Status       int       `json:"status,omitempty" xorm:"comment('状态  0：禁用   1：正常') TINYINT(4)"`
	CreateUserId int64     `json:"create_user_id,omitempty" xorm:"comment('创建者ID') BIGINT(20)"`
	CreateTime   time.Time `json:"create_time,omitempty" xorm:"comment('创建时间') DATETIME"`
}
