package Pic

import (
	"image"
	"qnsoft/qn_web_api/controllers/Token"
	"qnsoft/qn_web_api/utils/ErrorHelper"
	"qnsoft/qn_web_api/utils/FileHelper"
	"qnsoft/qn_web_api/utils/PicHelper"
	"qnsoft/qn_web_api/utils/StringHelper"
	"strings"

	"github.com/astaxie/beego"
	"github.com/boombuler/barcode/qr"
)

/*
二维码插件
*/
/*
二维码申城控制器
*/
type QrCode_Controller struct {
	Token.BaseController
}

/*
根据url地址生成普通二维码
*/
func (this *QrCode_Controller) Get_QR_Pic() {
	var _rt interface{}
	_qr_length, _ := this.GetInt("qr_length", 240)
	_qr_url := this.GetString("qr_url", "http://www.qnsoft.cn")
	qrc := StringHelper.NewQrCode(_qr_url, _qr_length, _qr_length, qr.Q, qr.Auto)
	path := StringHelper.GetQrCodeFullPath()
	_qr_img, _path, err := qrc.Encode(path)
	ErrorHelper.CheckErr(err)
	ErrorHelper.LogInfo("生成二维码执行", _qr_img)
	if len(_qr_img) > 5 {
		_rt = map[string]interface{}{"code": 200, "msg": "success", "info": "二维码已生成!", "src": beego.AppConfig.String("server_path::PrefixUrl") + _path + _qr_img}
	} else {
		_rt = map[string]interface{}{"code": 0, "msg": "fail", "info": "二维码生成失败!"}
	}
	this.Data["json"] = _rt
	this.ServeJSON()
}

/*
根据url地址生成带背景二维码
*/
func (this *QrCode_Controller) Get_QR_BgPic() {
	var _rt interface{}
	_qr_length, _ := this.GetInt("qr_length", 240)
	_qr_url := this.GetString("qr_url", "http://www.qnsoft.cn")
	_Pic_Dis := new(PicHelper.Pic_Dispose)
	//获取生成后的二维码地址
	_Qr_img := ""
	qrc := StringHelper.NewQrCode(_qr_url, _qr_length, _qr_length, qr.Q, qr.Auto)
	path := StringHelper.GetQrCodeFullPath()
	_qr_img, _path, err := qrc.Encode(path)
	ErrorHelper.CheckErr(err)
	ErrorHelper.LogInfo("生成二维码执行", _qr_img)
	bg_pic_path := beego.AppConfig.String("server_path::QrCodebgPic")
	new_pic_path := _path + "66666_" + strings.Replace(_qr_img, ".png", ".jpeg", 1)

	_Bg_Pic_Model := PicHelper.Pic_Model{Path: bg_pic_path}
	_Ft_Pic_Model := PicHelper.Pic_Model{Path: _path + _qr_img, P: image.Pt(255, 658)}
	new_bg_img := _Pic_Dis.Pic_pic_ompose(_Bg_Pic_Model, _Ft_Pic_Model, new_pic_path) //图片与图片合成

	_Ft_Text_Model := PicHelper.Pic_Text{
		Text:  string("测试姓名"), //真实姓名
		Color: [4]uint8{7, 7, 7, 255},
		//	FontFile:  beego.AppConfig.String("server_path::FontPath") + "simhei.ttf",
		Size:      0.36,
		Linewidth: 2,
		Angle:     0.3,
		Space:     "   ",
		Px:        590,
		Py:        786,
	}
	_Bg_Pic_Model1 := PicHelper.Pic_Model{Path: new_bg_img}
	// 图片与文字合成文件
	new_img_qr := beego.AppConfig.String("server_path::RuntimeRootPath") + "QR_" + "13938202388" + ".jpeg"
	err0 := _Pic_Dis.Pic_text_ompose(_Bg_Pic_Model1, _Ft_Text_Model, new_img_qr) //图片与文字合成
	ErrorHelper.CheckErr(err0)
	_Bg_Pic_Model_a := PicHelper.Pic_Model{Path: new_img_qr}
	_touxiang := beego.AppConfig.String("server_path::RuntimeRootPath") + "LOGO.png"
	// if len(_model.Headimgurl) > 5 {
	// 	_touxiang = _model.Headimgurl
	// }
	_Ft_Pic_Model_a := PicHelper.Pic_Model{Path: _touxiang, P: image.Pt(146, 330), IsScale: true, Width: 120, Height: 120}
	//new_pic_path_a := "User_" + new_img_qr
	new_bg_img_a := _Pic_Dis.Pic_pic_ompose(_Bg_Pic_Model_a, _Ft_Pic_Model_a, new_img_qr) //图片与头像合成最终效果图
	//执行完之后删除之前的二维码和不含水印文字的合成图
	if FileHelper.CheckNotExist("/" + _path + _qr_img) {
		err := FileHelper.DeleteFile(_path + _qr_img) //删除二维码
		ErrorHelper.CheckErr(err)
	}
	if FileHelper.CheckNotExist("/" + new_pic_path) {
		err := FileHelper.DeleteFile(new_pic_path) //删除合成图
		ErrorHelper.CheckErr(err)
	}
	//判断二维码是否生成成功
	if FileHelper.CheckNotExist("/" + new_bg_img_a) {
		_Qr_img = new_bg_img_a
	} else {
		_Qr_img = "nopic"
	}
	//	return _Qr_img
	if len(_Qr_img) > 5 {
		_rt = map[string]interface{}{"code": 200, "msg": "success", "info": "二维码已生成!", "src": beego.AppConfig.String("server_path::PrefixUrl") + _Qr_img}
	} else {
		_rt = map[string]interface{}{"code": 0, "msg": "fail", "info": "二维码生成失败!"}
	}
	this.Data["json"] = _rt
	this.ServeJSON()
}
