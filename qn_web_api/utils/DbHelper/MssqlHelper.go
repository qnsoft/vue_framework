package DbHelper

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

func MsSqlDb() *xorm.Engine {
	var Engine *xorm.Engine
	var dbError error
	//数据库类型
	_type := beego.AppConfig.String("database_mssql::db_type")

	//数据库IP
	//数据库IP
	_server := beego.AppConfig.String("database_mssql::db_server")
	//数据库端口
	_port := beego.AppConfig.String("database_mssql::db_port")
	////数据库
	_database := beego.AppConfig.String("database_mssql::db_database")
	//数据库用户名
	_user := beego.AppConfig.String("database_mssql::db_user")
	//数据库密码
	_password := beego.AppConfig.String("database_mssql::db_password")

	if Engine != nil {
		return Engine
	}
	//连接字符串
	_connString := fmt.Sprintf("server=%s;port=%s;database=%s;user id=%s;password=%s", _server, _port, _database, _user, _password)
	Engine, dbError = xorm.NewEngine(_type, _connString)
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "tb_")
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
