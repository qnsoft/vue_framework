package plugs

import (
	"qnsoft/qn_web_api/plugs/Pic"

	"github.com/astaxie/beego"
)

/*
图片相关路由
*/
func Plugs_Pic() {
	//生成普通二维码
	beego.Router("/api/pic/get_qr_pic", &Pic.QrCode_Controller{}, "post,get:Get_QR_Pic")
	//生成带背景二维码
	beego.Router("/api/pic/get_qr_bgpic", &Pic.QrCode_Controller{}, "post,get:Get_QR_BgPic")
}
