package PicHelper

import (
	"net/http"

	"github.com/mojocn/base64Captcha"
)

/*
线上案例地址
在线Demo [Playground Powered by Vuejs+elementUI+Axios](http://captcha.mojotv.cn)
*/
//创建base64数字验证码配置
var config_Digit = base64Captcha.ConfigDigit{
	Height:     80,
	Width:      240,
	MaxSkew:    0.7,
	DotCount:   80,
	CaptchaLen: 4,
}

//创建base64数字验证码配置
var config_Character = base64Captcha.ConfigCharacter{
	Mode:             3,     //样式  base64Captcha.CaptchaModeNumberAlphabet,//CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
	Height:           80,    //高
	Width:            240,   //宽
	IsUseSimpleFont:  true,  //是否使用字体
	IsShowHollowLine: false, //显示空心横线
	IsShowNoiseDot:   true,  // 显示噪声干扰点
	IsShowNoiseText:  false, //显示噪声干扰字符
	IsShowSlimeLine:  false, //显示细线
	IsShowSineLine:   false, //显示曲线
	CaptchaLen:       4,     //显示字符数
}

//声音验证码配置
var config_Audio = base64Captcha.ConfigAudio{
	CaptchaLen: 4,    //显示字符数
	Language:   "zh", //语言
}

/*
数字验证码
*/
func Pic_verifycode_digit(w http.ResponseWriter, r *http.Request) (string, string) {
	idKeyD, capD := base64Captcha.GenerateCaptcha("", config_Digit)
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)
	return idKeyD, base64stringD
}

/*
数字+字母验证码
*/
func Pic_verifycode_character(w http.ResponseWriter, r *http.Request) (string, string) {
	idKeyD, capD := base64Captcha.GenerateCaptcha("", config_Character)
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)
	return idKeyD, base64stringD
}

/*
数字+字母验证码
*/
func Pic_verifycode_audio(w http.ResponseWriter, r *http.Request) (string, string) {
	idKeyD, capD := base64Captcha.GenerateCaptcha("", config_Audio)
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)
	return idKeyD, base64stringD
}
