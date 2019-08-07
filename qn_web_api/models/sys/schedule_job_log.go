package models

import (
	"time"
)

type ScheduleJobLog struct {
	LogId      int64     `xorm:"not null pk autoincr comment('任务日志id') BIGINT(20)"`
	JobId      int64     `xorm:"not null comment('任务id') index BIGINT(20)"`
	BeanName   string    `xorm:"comment('spring bean名称') VARCHAR(200)"`
	Params     string    `xorm:"comment('参数') VARCHAR(2000)"`
	Status     int       `xorm:"not null comment('任务状态    0：成功    1：失败') TINYINT(4)"`
	Error      string    `xorm:"comment('失败信息') VARCHAR(2000)"`
	Times      int       `xorm:"not null comment('耗时(单位：毫秒)') INT(11)"`
	CreateTime time.Time `xorm:"comment('创建时间') DATETIME"`
}
