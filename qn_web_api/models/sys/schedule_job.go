package models

import (
	"time"
)

type ScheduleJob struct {
	JobId          int64     `xorm:"not null pk autoincr comment('任务id') BIGINT(20)"`
	BeanName       string    `xorm:"comment('spring bean名称') VARCHAR(200)"`
	Params         string    `xorm:"comment('参数') VARCHAR(2000)"`
	CronExpression string    `xorm:"comment('cron表达式') VARCHAR(100)"`
	Status         int       `xorm:"comment('任务状态  0：正常  1：暂停') TINYINT(4)"`
	Remark         string    `xorm:"comment('备注') VARCHAR(255)"`
	CreateTime     time.Time `xorm:"comment('创建时间') DATETIME"`
}
