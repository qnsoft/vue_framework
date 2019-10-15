package routers

import (
	"qnsoft/qn_web_api/controllers"
	"qnsoft/qn_web_api/controllers/Token"
	"qnsoft/qn_web_api/plugs/ImageUpload"
	"qnsoft/qn_web_api/routers/plugs"
	"qnsoft/qn_web_api/routers/sys"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	//加入跨域权限
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Token"},
		AllowCredentials: true}))
	beego.SetStaticPath("/upload", "upload")              //重定向静态路径 上传查看
	beego.SetStaticPath("/logs", "logs")                  //重定向静态路径 日志查看
	beego.SetStaticPath("/views", "views")                //重定向静态路径 模板查看
	beego.SetStaticPath("/admin", "views/web_root/admin") //重定向静态路径 模板查看
	//-------------------------基础接口开始-----------------------------------//
	//根目录
	beego.Router("/", &controllers.Default_Controller{}, "*:Get")
	//wap目录
	beego.Router("/m", &controllers.Default_Controller{}, "*:Get_Wap")
	//先获取access_token
	beego.Router("/access_token", &Token.AccesstokenController{}, "post:Access_Token")
	//token测试
	//beego.Router("/testtoken", &controllers.Default_Controller{}, "get,post:TestToken")
	//上传图片
	beego.Router("/img/upload_pic", &ImageUpload.Image_Uplaod_Controller{}, "post:Uplaod_Pic")
	//获取图片
	beego.Router("/img/info_pic", &ImageUpload.Image_Uplaod_Controller{}, "get:Info_Pic")
	//-------------------------基础接口结束-----------------------------------//
	//-------------------------sys相关接口路由开始-----------------------------------//
	sys.Sys_router()
	//-------------------------sys相关接口路由结束-----------------------------------//
	//-------------------------DEMO接口结束-----------------------------------//
	//与淘宝对接接口
	//Mallother()
	//短信相关接口
	plugs.Plugs_Sms()
	//图片插件相关接口
	plugs.Plugs_Pic()
}
