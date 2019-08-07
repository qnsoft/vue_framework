package models

/* 系统设置实体对象 */
type SysConfig struct {
	Id         int64  `json:"id,omitempty" xorm:"pk autoincr BIGINT(20)"`
	ParamKey   string `json:"param_key,omitempty" xorm:"comment('key') unique VARCHAR(50)"`
	ParamValue string `json:"param_value,omitempty" xorm:"comment('value') VARCHAR(2000)"`
	Status     int    `json:"status,omitempty" xorm:"default 1 comment('状态   0：隐藏   1：显示') TINYINT(4)"`
	Remark     string `json:"remark,omitempty" xorm:"comment('备注') VARCHAR(500)"`
}
