package DbHelper

import (
	"fmt"
	"qnsoft/qn_web_api/utils/ErrorHelper"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

/*
*@数据库链接
 */
func ConnectionDb() (*mgo.Database, *mgo.Session) {
	mgo_url := beego.AppConfig.String("database::db_path")
	session, err := mgo.Dial(mgo_url)
	if err != nil {
		fmt.Println(err.Error())
		//os.Exit(-1)
		fmt.Println("------连接数据库失败！------------")
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true) //这里需要写清楚链接模式
	db := session.DB("QN_CMS")           //数据库名称
	return db, session
}

/*
*@数据库关闭
 */
func ColoseDb() {
	_, session := ConnectionDb()
	defer session.Close()
}

/*
*@获取单条信息
*@_query查询条件表达式
 */
func Single(_query map[string]interface{}, _model interface{}, _doc_name string) (interface{}, error) {
	db, _ := ConnectionDb()
	_collec := db.C(_doc_name)
	err := _collec.Find(_query).One(&_model)
	ErrorHelper.CheckErr(err)
	ColoseDb()
	return _model, err
}

/*
*@获取列表信息不分页
*@_query查询条件表达式
*@_doc_name文档名
*@_sort排序表达式
 */
func List(_query interface{}, _doc_name string, _sort ...string) ([]interface{}, int, error) {
	db, _ := ConnectionDb()
	_collec := db.C(_doc_name)
	_info_list := make([]interface{}, 20)
	err := _collec.Find(_query).Sort(_sort...).All(&_info_list)
	ErrorHelper.CheckErr(err)
	_total := len(_info_list)
	ColoseDb()
	return _info_list, _total, err
}

/*
*@获取列表信息分页
*@page_on当前页码
*@page_size每页显示条数
*@_query查询条件表达式
*@_doc_name文档名
*@_sort排序表达式
 */
func List_Page(page_on, page_size int, _query map[string]interface{}, _doc_name string, _sort ...string) ([]interface{}, int, error) {
	db, _ := ConnectionDb()
	_collec := db.C(_doc_name)
	info_list_count := make([]interface{}, 20)
	err1 := _collec.Find(_query).All(&info_list_count)
	ErrorHelper.CheckErr(err1)
	_total := len(info_list_count)
	_info_list := make([]interface{}, 20)
	err2 := _collec.Find(_query).Sort(_sort...).Limit(page_size).Skip(page_size * (page_on - 1)).All(&_info_list)
	ErrorHelper.CheckErr(err2)
	ColoseDb()
	return _info_list, _total, err2
}

/*
*@增加集合
*@doc添加内容对象
 */
func Insert(doc interface{}, _doc_name string) error {
	db, _ := ConnectionDb()
	_collec := db.C(_doc_name)
	err := _collec.Insert(doc)
	ErrorHelper.CheckErr(err)
	ColoseDb()
	return err
}

/*
*@修改集合
*@_query修改条件表达式
*@edit_doc修改内容表达式
 */
func Update(_query interface{}, edit_doc interface{}, _doc_name string) error {
	db, _ := ConnectionDb()
	_collec := db.C(_doc_name)
	err := _collec.Update(_query, edit_doc)
	ErrorHelper.CheckErr(err)
	ColoseDb()
	return err
}

/*
*@删除集合
*@_query删除条件表达式
 */
func Delede(_query interface{}, _doc_name string) error {
	db, _ := ConnectionDb()
	_collec := db.C(_doc_name)
	err := _collec.Remove(_query)
	ErrorHelper.CheckErr(err)
	ColoseDb()
	return err
}
