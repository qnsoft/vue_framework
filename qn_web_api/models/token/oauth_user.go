package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

/*
*@token用户信息实体
 */
type User struct {
	Id              int       `json:"id" pk:"auto" orm:"column(id)"`
	Name            string    `json:"name" orm:"column(name)"`
	Passwd          string    `json:"passwd" orm:"column(passwd)"`
	Email           string    `json:"email" orm:"column(email)"`
	Status          int       `json:"status" orm:"column(status)"`
	Create_time     time.Time `json:"create_time" orm:"column(create_time)"`
	Last_login_time time.Time `json:"last_login_time" orm:"column(last_login_time)"`
	Role_id         string    `json:"role_id" orm:"column(role_id)"`
	Appid           string    `json:"appid" orm:"column(appid)"`
	Appsecret       string    `json:"appsecret" orm:"column(appsecret)"`
	Salt            string    `json:"salt" orm:"column(salt)"`
}

func init() {
	orm.RegisterModelWithPrefix(beego.AppConfig.String("database::db_prefix"), new(User))
}

/*
*@检查登录
*@_Appid 用户id
*@_Appsecret 用户密码
 */
func (s *User) Check_Oauth(_Appid, _Appsecret string) (bool, error) {
	var _rt bool = false
	var err error
	// table := TableName("user")
	// query := orm.NewOrm().QueryTable(table)
	// cond := orm.NewCondition()
	// cond1 := cond.And("appid__exact", _Appid).And("appsecret__exact", _Appsecret) //查询分组1
	// query = query.SetCond(cond1)
	// total, err := query.Count()
	// if err == nil && total > 0 {
	// 	_rt = true
	// } else {
	// 	_rt = false
	// }
	return _rt, err
}
