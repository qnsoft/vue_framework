package StringHelper

import (
	"image/png"
	"os"
	"qnsoft/qn_web_api/utils/ErrorHelper"
	"qnsoft/qn_web_api/utils/FileHelper"

	"github.com/astaxie/beego"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

/*
二维码
*/
type QrCode struct {
	URL    string
	Width  int
	Height int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

const (
	EXT_JPG = ".png"
)

func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		URL:    url,
		Width:  width,
		Height: height,
		Level:  level,
		Mode:   mode,
		Ext:    EXT_JPG,
	}
}

func GetQrCodePath() string {
	return beego.AppConfig.String("server_path::QrCodeSavePath")
}

func GetQrCodeFullPath() string {
	return beego.AppConfig.String("server_path::RuntimeRootPath")
}

func GetQrCodeFullUrl(name string) string {
	return beego.AppConfig.String("server_path::PrefixUrl") + GetQrCodePath() + name
}

func GetQrCodeFileName(value string) string {
	return Md5(value)
}

func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}

func (q *QrCode) CheckEncode(path string) bool {
	src := path + GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	if FileHelper.CheckNotExist(src) == true {
		return false
	}

	return true
}

/*
生成二维码
*/
func (q *QrCode) Encode(path string) (string, string, error) {
	name := GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	src := path + name
	if FileHelper.CheckNotExist(src) == true {
		code, err := qr.Encode(q.URL, q.Level, q.Mode)
		if err != nil {
			ErrorHelper.LogInfo("生成二维码出错1：", err)
			return "", "", err
		}

		code, err = barcode.Scale(code, q.Width, q.Height)
		if err != nil {
			ErrorHelper.LogInfo("生成二维码出错2：", err)
			return "", "", err
		}
		f, err := os.Create(src)
		if err != nil {
			ErrorHelper.LogInfo("打开已生成二维码文件出错：", err)
			//log.Fatal(err)
		}
		err = png.Encode(f, code) //质量
		// //	err = png.Encode(f, code, &jpeg.Options{10}) //质量
		if err != nil {
			ErrorHelper.LogInfo("生成png二维码出错：", err)
			return "", "", err
		}
		defer f.Close()
	}
	return name, path, nil
}
