package ImageUpload

import (
	"bufio"
	"encoding/json"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"math/rand"
	"os"
	"strings"
	"time"

	"qnsoft/qn_web_api/controllers/Token"
	"qnsoft/qn_web_api/plugs/ImageUpload/imghand"

	"github.com/astaxie/beego"
)

/**
*信息实体
 */
type Image_Uplaod_Controller struct {
	Token.BaseController
}

/**
*@post上传图片
 */
func (u *Image_Uplaod_Controller) Uplaod_Pic() {
	//token检测
	if u.Check_Token() {
		// 响应返回
		res := new(UpdateResponse)
		// 上传表单 --------------------------------------
		// 缓冲的大小 - 4M
		u.Ctx.Request.ParseMultipartForm(1024 << 12)
		//是上传表单域的名字fileHeader
		upfile, upFileInfo, err := u.Ctx.Request.FormFile("userfile")
		if err != nil {
			res.Code = StatusForm
			res.Msg = StatusText(StatusForm)
			u.Ctx.ResponseWriter.Write(ResponseJson(res))
			return
		}
		defer upfile.Close()

		// 图片解码 --------------------------------------

		// 读入缓存
		bufUpFile := bufio.NewReader(upfile)
		// 进行图片的解码
		img, imgtype, err := image.Decode(bufUpFile)
		if err != nil {

			res.Code = StatusImgDecode
			res.Msg = StatusText(StatusImgDecode)
			u.Ctx.ResponseWriter.Write(ResponseJson(res))

			return
		}

		// 判断是否有这个图片类型
		if !imghand.IsType(imgtype) {

			res.Code = StatusImgIsType
			res.Msg = StatusText(StatusImgIsType)
			u.Ctx.ResponseWriter.Write(ResponseJson(res))

			return
		}

		// 设置文件读写下标 --------------------------------

		// 设置下次读写位置（移动文件指针位置）
		_, err = upfile.Seek(0, 0)
		if err != nil {

			res.Code = StatusFileSeek
			res.Msg = StatusText(StatusFileSeek)
			u.Ctx.ResponseWriter.Write(ResponseJson(res))

			return
		}

		//生成本地路径
		t := time.Now()
		file_path := fmt.Sprintf("%d/%d/%d/", t.Year(), t.Month(), t.Day())
		index := strings.LastIndex(upFileInfo.Filename, ".")
		_file := fmt.Sprintf("%d%d%d%d%d%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()) + fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)) + upFileInfo.Filename[index:] //连接后缀名
		fmt.Println(_file)
		// 组合文件完整路径
		dirPath := imghand.JoinPath(file_path) //创建文件目录
		filePath := dirPath + _file            // 文件路径
		// 获取目录信息，并创建目录
		dirInfo, err := os.Stat(dirPath)
		if err != nil {
			err = os.MkdirAll(dirPath, 0666)
			if err != nil {

				res.Code = StatusMkdir
				res.Msg = StatusText(StatusMkdir)
				u.Ctx.ResponseWriter.Write(ResponseJson(res))

				return
			}
		} else {
			if !dirInfo.IsDir() {
				err = os.MkdirAll(dirPath, 0666)
				if err != nil {

					res.Code = StatusMkdir
					res.Msg = StatusText(StatusMkdir)
					u.Ctx.ResponseWriter.Write(ResponseJson(res))

					return
				}
			}
		}
		// 存入文件 --------------------------------------
		_, err = os.Stat(filePath)
		if err != nil {

			// 打开一个文件,文件不存在就会创建
			file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				res.Code = StatusOpenFile
				res.Msg = StatusText(StatusOpenFile)
				u.Ctx.ResponseWriter.Write(ResponseJson(res))
				return
			}
			defer file.Close()

			if imgtype == imghand.PNG {
				err = png.Encode(file, img)

			} else if imgtype == imghand.JPG || imgtype == imghand.JPEG {
				err = jpeg.Encode(file, img, nil)

			} else if imgtype == imghand.GIF {
				// 重新对 gif 格式进行解码
				// image.Decode 只能读取 gif 的第一帧
				// 设置下次读写位置（移动文件指针位置）
				_, err = upfile.Seek(0, 0)
				if err != nil {

					res.Code = StatusFileSeek
					res.Msg = StatusText(StatusFileSeek)
					u.Ctx.ResponseWriter.Write(ResponseJson(res))
					return
				}
				gifimg, giferr := gif.DecodeAll(upfile)
				if giferr != nil {
					res.Code = StatusImgDecode
					res.Msg = StatusText(StatusImgDecode)
					u.Ctx.ResponseWriter.Write(ResponseJson(res))

					return
				}
				err = gif.EncodeAll(file, gifimg)
			}
			if err != nil {
				res.Code = StatusImgEncode
				res.Msg = StatusText(StatusImgEncode)
				u.Ctx.ResponseWriter.Write(ResponseJson(res))
				return
			}
		}
		res.Success = true
		res.Code = StatusOK
		res.Msg = StatusText(StatusOK)
		res.Data.Imgid = beego.AppConfig.String("ServerUrl") + "/" + filePath // fileMd5
		res.Data.Mime = imgtype
		res.Data.Size = upFileInfo.Size

		u.Ctx.ResponseWriter.Write(ResponseJson(res))

	}
}

// 获取图片信息
func (u *Image_Uplaod_Controller) Info_Pic() {

	// 响应返回
	res := new(UpdateResponse)

	// 获取要图片id
	imgid := u.Ctx.Request.FormValue("imgid")

	// 获取裁剪后图像的宽度、高度
	width := imghand.StringToInt(u.Ctx.Request.FormValue("w"))  // 宽度
	height := imghand.StringToInt(u.Ctx.Request.FormValue("h")) // 高度

	// 组合文件完整路径
	filePath := imghand.UrlParse(imgid)
	if filePath == "" {

		res.Code = StatusUrlNotFound
		res.Msg = StatusText(StatusUrlNotFound)
		u.Ctx.ResponseWriter.Write(ResponseJson(res))

		return
	}

	if width != 0 || height != 0 {
		filePath = fmt.Sprintf("%s_%d_%d", filePath, width, height)
	}

	fimg, err := os.Open(filePath)
	if err != nil {

		res.Code = StatusImgNotFound
		res.Msg = StatusText(StatusImgNotFound)
		u.Ctx.ResponseWriter.Write(ResponseJson(res))

		return
	}
	defer fimg.Close()

	bufimg := bufio.NewReader(fimg)
	_, imgtype, err := image.Decode(bufimg)
	if err != nil {

		res.Code = StatusImgNotFound
		res.Msg = StatusText(StatusImgNotFound)
		u.Ctx.ResponseWriter.Write(ResponseJson(res))

		return
	}

	finfo, _ := fimg.Stat()
	res.Success = true
	res.Code = StatusOK
	res.Msg = StatusText(StatusOK)
	res.Data.Imgid = imgid
	res.Data.Mime = imgtype
	res.Data.Size = finfo.Size()
	u.Ctx.ResponseWriter.Write(ResponseJson(res))

}

// 状态码
func (u *Image_Uplaod_Controller) StatusCode() {
	data, _ := json.Marshal(GetStatusText())
	u.Ctx.ResponseWriter.Write(data)
}
