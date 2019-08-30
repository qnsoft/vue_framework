package PicHelper

import (
	"bytes"
	"errors"
	"net/http"
	"path"
	"strings"

	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"math"
	"os"

	"qnsoft/qn_web_api/utils/ErrorHelper"

	"github.com/golang/freetype"
	"github.com/nfnt/resize"
	"golang.org/x/image/bmp"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/vgimg"
)

/*
图片处理对象
*/
type Pic_Dispose struct {
}

/*
图片对象
*/
type Pic_Model struct {
	//图片路径
	Path string
	//图片定位
	P image.Point
	//是否缩略
	IsScale bool
	//缩略图宽
	Width int
	//缩略图高
	Height int
}

/*
文本对象
*/
type Pic_Text struct {
	Text      string    //文字内容
	FontFile  string    //字体路径
	Color     [4]uint8  //文字颜色RGBA 例子: [0,0,0,122]
	Size      vg.Length //字体大小 单位英寸
	Linewidth vg.Length //行高
	Angle     float64   //角度 0至1 的范围 即0到90度之间
	Space     string    //空格间隔
	Px        float64   //横向(距离右边距离)
	Py        float64   //纵向(距离下边距离)
}

/*
图片与图片合成
*/
func (_pic Pic_Dispose) Pic_pic_ompose(Bg_pic, Ft_pic Pic_Model, New_pic string) string {
	// ErrorHelper.CheckErr(err2) //把图片解码为结构体时出错
	bgimg, err2 := _pic.Get_pic(Bg_pic)
	ErrorHelper.CheckErr(err2) //把图片解码为结构体时出错
	wmb_img, err4 := _pic.Get_pic(Ft_pic)
	ErrorHelper.CheckErr(err4) //把水印图片解码为结构体时出错
	b := bgimg.Bounds()
	//根据b画布的大小新建一个新图像
	m := image.NewRGBA(b)
	draw.Draw(m, b, bgimg, image.ZP, draw.Src)
	draw.Draw(m, wmb_img.Bounds().Add(Ft_pic.P), wmb_img, image.ZP, draw.Over)
	//生成新图片new.jpg,并设置图片质量
	imgw, err5 := os.Create(New_pic)
	ErrorHelper.CheckErr(err5)
	//ErrorHelper.LogInfo("合成出错了", err5)
	jpeg.Encode(imgw, m, &jpeg.Options{100}) //生成jpeg质量
	defer imgw.Close()
	return New_pic
}

/*
图片与文字合成
*/
func (_pic *Pic_Dispose) Pic_text_ompose(Bg_pic Pic_Model, Ft_text Pic_Text, New_pic string) error {
	f, _ := os.Create(New_pic)
	img2, err2 := _pic.Get_pic(Bg_pic)
	ErrorHelper.CheckErr(err2)
	// 打文字
	img3, err3 := _pic.Pic_text_ompose_one(img2, Ft_text)
	ErrorHelper.CheckErr(err3)
	//获取文件后缀名
	ext := path.Ext(Bg_pic.Path)
	// 将 image 写入 buffur
	buff, _ := _pic.ImageToBuffer(img3, ext)
	// buffer 写入水印文件
	buff.WriteTo(f)
	//defer f.Close()
	return err2
}

/*
用于在图像上添加文字
*/
func (_pic *Pic_Dispose) Pic_text_ompose_one(img image.Image, markText Pic_Text) (image.Image, error) {
	// 图片的长度设置画布的长度
	bounds := img.Bounds()
	w := vg.Length(bounds.Max.X) * vg.Inch / vgimg.DefaultDPI
	h := vg.Length(bounds.Max.Y) * vg.Inch / vgimg.DefaultDPI
	// 通过高和宽计算对角线
	diagonal := vg.Length(math.Sqrt(float64(w*w + h*h)))

	// 创建一个画布，宽度和高度是对角线
	c := vgimg.New(diagonal, diagonal)

	// 在画布中心绘制图像
	rect := vg.Rectangle{}
	// 计算中心位置,宽为w,高为h
	rect.Min.X = diagonal/2 - w/2
	rect.Min.Y = diagonal/2 - h/2
	rect.Max.X = diagonal/2 + w/2
	rect.Max.Y = diagonal/2 + h/2
	c.DrawImage(rect, img)

	// 读字体数据
	fontBytes, err1 := ioutil.ReadFile("fonts/zh/simhei.ttf")
	ErrorHelper.CheckErr(err1)
	// //加载字体库文件
	font, err2 := freetype.ParseFont(fontBytes)
	ErrorHelper.CheckErr(err2)
	//加载字库文件流
	//font := trueTypeFontFamilys[0]

	vg.AddFont("cn_font", font)
	fontStyle, err3 := vg.MakeFont("cn_font", vg.Inch*markText.Size)
	ErrorHelper.CheckErr(err3)
	//ErrorHelper.LogInfo("字体大小出错了", err3)
	// 重复编写水印字体
	marktext := markText.Text
	// 设置水印字体的颜色
	rgba := markText.Color
	c.SetColor(color.RGBA{rgba[0], rgba[1], rgba[2], rgba[3]})
	_Point := vg.Point{X: vg.Points(markText.Px), Y: vg.Points(markText.Py)}
	//设置文字坐标间距 新版本
	c.FillString(fontStyle, _Point, marktext)
	// 画布写入新图片
	// 使用buffer去转换
	jc := vgimg.PngCanvas{Canvas: c}
	buff := new(bytes.Buffer)
	jc.WriteTo(buff)
	img, _, _ = image.Decode(buff)
	// 得到图像的中心点
	ctp := int(diagonal * vgimg.DefaultDPI / vg.Inch / 2)
	// 切出打水印的图像
	size := bounds.Size()
	bounds = image.Rect(ctp-size.X/2, ctp-size.Y/2, ctp+size.X/2, ctp+size.Y/2)
	rv := image.NewRGBA(bounds)
	draw.Draw(rv, bounds, img, bounds.Min, draw.Src)
	return rv, nil
}

func (_pic *Pic_Dispose) ImageToBuffer(img image.Image, ext string) (rv *bytes.Buffer, err error) {
	ext = strings.ToLower(ext)
	rv = new(bytes.Buffer)
	switch ext {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(rv, img, &jpeg.Options{Quality: 50})
	case ".png":
		err = png.Encode(rv, img)
	}
	return rv, err
}

/*
获取http文件流
_model 文件路径
*/
func (_pic *Pic_Dispose) Get_pic(_model Pic_Model) (image.Image, error) {
	var _image image.Image
	var err1 error
	if strings.Contains(_model.Path, "http") {
		resp, err := http.Get(_model.Path)
		ErrorHelper.CheckErr(err)
		body, _ := ioutil.ReadAll(resp.Body)
		_image, _, err1 = image.Decode(bytes.NewReader(body))
		ErrorHelper.CheckErr(err1)
	} else {
		f1, err1 := os.Open(_model.Path)
		ErrorHelper.CheckErr(err1)
		defer f1.Close()
		_image, _, err1 = image.Decode(f1)
	}
	if _model.IsScale {
		if _model.Width == 0 || _model.Height == 0 {
			_model.Width = _image.Bounds().Max.X
			_model.Height = _image.Bounds().Max.Y
		}
		_image = resize.Thumbnail(uint(_model.Width), uint(_model.Height), _image, resize.Lanczos3)
	}
	return _image, err1
}

/*
* 图片裁剪
* 入参:
* 规则:如果精度为0则精度保持不变
 */
func (_pic *Pic_Dispose) Clip(in io.Reader, out io.Writer, x0, y0, x1, y1, quality int) error {
	origin, fm, err := image.Decode(in)
	if err != nil {
		return err
	}

	switch fm {
	case "jpeg":
		img := origin.(*image.YCbCr)
		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.YCbCr)
		return jpeg.Encode(out, subImg, &jpeg.Options{quality})
	case "png":
		switch origin.(type) {
		case *image.NRGBA:
			img := origin.(*image.NRGBA)
			subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.NRGBA)
			return png.Encode(out, subImg)
		case *image.RGBA:
			img := origin.(*image.RGBA)
			subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
			return png.Encode(out, subImg)
		}
	case "gif":
		img := origin.(*image.Paletted)
		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.Paletted)
		return gif.Encode(out, subImg, &gif.Options{})
	case "bmp":
		img := origin.(*image.RGBA)
		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
		return bmp.Encode(out, subImg)
	default:
		return errors.New("ERROR FORMAT")
	}
	return nil
}

/*
* 缩略图生成
* 入参:
* 规则: 如果width 或 hight其中有一个为0，则大小不变 如果精度为0则精度保持不变
* 矩形坐标系起点是左上
* 返回:error
 */
func (_pic *Pic_Dispose) Scale(in io.Reader, out io.Writer, width, height, quality int) (image.Image, error) {
	origin, _, err := image.Decode(in)
	ErrorHelper.CheckErr(err)
	if width == 0 || height == 0 {
		width = origin.Bounds().Max.X
		height = origin.Bounds().Max.Y
	}
	if quality == 0 {
		quality = 100
	}
	_image := resize.Thumbnail(uint(width), uint(height), origin, resize.Lanczos3)
	return _image, err
	/* //return jpeg.Encode(out, canvas, &jpeg.Options{quality})

	switch fm {
	case "jpeg":
		return jpeg.Encode(out, canvas, &jpeg.Options{quality})
	case "png":
		return png.Encode(out, canvas)
	case "gif":
		return gif.Encode(out, canvas, &gif.Options{})
	case "bmp":
		return bmp.Encode(out, canvas)
	default:
		return errors.New("ERROR FORMAT")
	}
	return nil */
}
