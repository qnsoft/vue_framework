package routers

import (
	// "qnsoft/qn_web_api/plugs"
	"qnsoft/qn_web_api/plugs/ImageUpload"
	//"qnsoft/qn_web_api/plugs/OnlinePay"
	"qnsoft/qn_web_api/controllers/Token"
	"qnsoft/qn_web_api/plugs/Sms"
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
	beego.SetStaticPath("/upload", "upload") //重定向静态路径
	beego.SetStaticPath("/views", "views")   //重定向静态路径
	//-------------------------基础接口开始-----------------------------------//
	//根目录
	//beego.Router("/", &controllers.Default_Controller{}, "get,post:Get")
	//先获取access_token
	beego.Router("/access_token", &Token.AccesstokenController{}, "post:Access_Token")
	//token测试
	//beego.Router("/testtoken", &controllers.Default_Controller{}, "get,post:TestToken")
	//上传图片
	beego.Router("/img/upload_pic", &ImageUpload.Image_Uplaod_Controller{}, "post:Uplaod_Pic")
	//获取图片
	beego.Router("/img/info_pic", &ImageUpload.Image_Uplaod_Controller{}, "get:Info_Pic")
	//发送短信验证码
	beego.Router("/sms/MsgSend", &Sms.Sms_Controller{}, "post:MsgSend")
	//-------------------------基础接口结束-----------------------------------//
	//-------------------------sys相关接口路由开始-----------------------------------//
	routers.Sys_router()
	//-------------------------sys相关接口路由结束-----------------------------------//
	//-------------------------DEMO接口结束-----------------------------------//
	//与淘宝对接接口
	//Mallother()
}
