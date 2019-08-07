package DbHelper

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

func MySqlDb() *xorm.Engine {
	var Engine *xorm.Engine
	var dbError error
	//数据库类型
	_type := beego.AppConfig.String("database_mysql::db_type")
	//数据库IP
	_server := beego.AppConfig.String("database_mysql::db_server")
	//数据库端口
	_port := beego.AppConfig.String("database_mysql::db_port")
	////数据库
	_database := beego.AppConfig.String("database_mysql::db_database")
	//数据库用户名
	_user := beego.AppConfig.String("database_mysql::db_user")
	//数据库密码
	_password := beego.AppConfig.String("database_mysql::db_password")
	//数据库密码
	_prefix := beego.AppConfig.String("database_mysql::db_prefix")
	//连接字符串
	_connString := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8", _user, _password, _server, _port, _database)
	Engine, dbError = xorm.NewEngine(_type, _connString)
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, _prefix)
	Engine.SetTableMapper(tbMapper)
	//Engine.SetMaxIdleConns(50)
	//Engine.SetMaxOpenConns(200)
	Engine.ShowSQL(true)
	Engine.Logger().SetLevel(core.LOG_INFO)
	if dbError != nil {
		log.Fatal(dbError)
		panic(dbError)
	}
	return Engine
}
