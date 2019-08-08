package models

import (
	"time"
	"zhenfangbian/web_api/utils/DbHelper"
	"zhenfangbian/web_api/utils/ErrorHelper"
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

/*
获取单条信息
_user_id 用户id
*/
func (this *SysUser) Get_Info(_user_id int64) (bool, *SysUser) {
	_model := SysUser{UserId: _user_id}
	_results, err := DbHelper.MySqlDb().Get(&_model)
	ErrorHelper.CheckErr(err)
	return _results, &_model
}

/*
分页列表
*/
func (this *SysUser) Get_List_Info_Page(_where string, _pageIndex, _pageSize int) (int64, int, []*SysUser) {
	_model := new(SysUser)
	_totalCount, err1 := DbHelper.MySqlDb().Where(_where).Count(_model)
	ErrorHelper.CheckErr(err1)
	_rows, err2 := DbHelper.MySqlDb().Where(_where).Limit(_pageSize, (_pageIndex-1)*_pageSize).Rows(_model)
	ErrorHelper.CheckErr(err2)
	defer _rows.Close()
	_list := make([]*SysUser, 0)
	_totalPage := 0
	if int(_totalCount)%_pageSize == 0 {
		_totalPage = int(_totalCount) / _pageSize
	} else {
		_totalPage = int(_totalCount)/_pageSize + 1
	}
	for _rows.Next() {
		_ = _rows.Scan(_model)
		_model_new := new(SysUser)
		_model_new.UserId = _model.UserId
		_model_new.Username = _model.Username
		_model_new.Email = _model.Email
		_model_new.Mobile = _model.Mobile
		_model_new.Status = _model.Status
		_model_new.CreateUserId = _model.CreateUserId
		_model_new.CreateTime = _model.CreateTime
		_list = append(_list, _model_new)
	}
	return _totalCount, _totalPage, _list
}

/*
添加单条信息
_model 信息实体
*/
func (this *SysUser) Get_Info_Add(_model *SysUser) (int64, error) {
	_count, err := DbHelper.MySqlDb().Insert(_model)
	ErrorHelper.CheckErr(err)
	return _count, err
}

/*
修改单条信息
_model 信息实体
*/
func (this *SysUser) Get_Info_Update(_user_id int64, _model *SysUser) (int64, error) {
	_count, err := DbHelper.MySqlDb().Id(_user_id).Update(_model)
	ErrorHelper.CheckErr(err)
	return _count, err
}
