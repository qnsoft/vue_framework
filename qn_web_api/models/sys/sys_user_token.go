package models

import (
	"time"
)

type SysUserToken struct {
	UserId     int64     `xorm:"not null pk BIGINT(20)"`
	Token      string    `xorm:"not null comment('token') unique VARCHAR(100)"`
	ExpireTime time.Time `xorm:"comment('过期时间') DATETIME"`
	UpdateTime time.Time `xorm:"comment('更新时间') DATETIME"`
}
