package models

import (
	"time"
)

/*
系统日志实体对象
*/
type SysLog struct {
	Id         int64     `json:"id,omitempty" xorm:"pk autoincr BIGINT(20)"`
	Username   string    `json:"username,omitempty" xorm:"comment('用户名') VARCHAR(50)"`
	Operation  string    `json:"operation,omitempty" xorm:"comment('用户操作') VARCHAR(50)"`
	Method     string    `json:"method,omitempty" xorm:"comment('请求方法') VARCHAR(200)"`
	Params     string    `json:"params,omitempty" xorm:"comment('请求参数') VARCHAR(5000)"`
	Time       int64     `json:"time,omitempty" xorm:"not null comment('执行时长(毫秒)') BIGINT(20)"`
	Ip         string    `json:"ip,omitempty" xorm:"comment('IP地址') VARCHAR(64)"`
	CreateDate time.Time `json:"create_date,omitempty" xorm:"comment('创建时间') DATETIME"`
}
