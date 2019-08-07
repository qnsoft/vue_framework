package WebHelper

import (

	//"fmt"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/memcache"
	_ "github.com/astaxie/beego/cache/redis"
)

var cc cache.Cache

/*
*@初始化缓存引擎通道
 */
func init() {
	_cache_type := beego.AppConfig.String("cache::cache_type")
	switch strings.ToLower(_cache_type) {
	case "memory":
		cc, _ = cache.NewCache("memory", `{"interval":60}`)
	case "redis":
		cc, _ = cache.NewCache("redis", `{"conn":"`+beego.AppConfig.String("cache::cache_host")+`"}`)
	case "memcache":
		cc, _ = cache.NewCache("memcache", `{"conn":"`+beego.AppConfig.String("cache::cache_host")+`"}`)
	default:
		cc, _ = cache.NewCache("memory", `{"interval":60}`)
	}
}

/*
*@写入memory引擎缓存
 */
func Set_Cache(key string, value interface{}, timeout int) error {
	err := cc.Put(key, value, time.Duration(timeout)*time.Second)
	fmt.Println(cc.Get(key))
	return err
}

/*
*@读出memory引擎缓存
 */
func Get_Cache(key string) interface{} {
	_obj := cc.Get(key)
	return _obj
}

/*
*@根据键判断缓存是否存在
 */
func Is_Cache(key string) bool {
	return cc.IsExist(key)
}

/*
*@根据键删除缓存
 */
func Delete_Cache(key string) error {
	err := cc.Delete(key)
	return err
}

/*
*@清空所有缓存
 */
func Clear_Cache() {
	cc.ClearAll()
}
