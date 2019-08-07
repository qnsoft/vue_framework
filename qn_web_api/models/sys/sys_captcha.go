package models

import (
	"time"
)

type SysCaptcha struct {
	Uuid       string    `xorm:"not null pk comment('uuid') CHAR(36)"`
	Code       string    `xorm:"not null comment('验证码') VARCHAR(6)"`
	ExpireTime time.Time `xorm:"comment('过期时间') DATETIME"`
}
