package models

import (
	"time"
)

type SysOss struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	Url        string    `xorm:"comment('URL地址') VARCHAR(200)"`
	CreateDate time.Time `xorm:"comment('创建时间') DATETIME"`
}
