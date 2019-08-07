package FileHelper

import (
	"bytes"
	"strings"

	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"math"
	"os"

	"io/ioutil"
	"log"

	"github.com/golang/freetype"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/vgimg"
)

type MarkText struct {
	Text      string    //文字内容
	FontFile  string    //字体路径
	Color     [4]uint8  //文字颜色RGBA 例子: [0,0,0,122]
	Size      vg.Length //字体大小 单位英寸
	Linewidth vg.Length //行高
	Angle     float64   //角度 0至1 的范围 即0到90度之间
	Space     string    //空格间隔
}

// WaterMark用于在图像上添加水印
func WaterMark(img image.Image, markText MarkText) (image.Image, error) {
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

	// 制作一个 fontstyle ，宽度为英寸,字体 Courier 标准的等宽度字体
	// 读字体数据
	fontBytes, err := ioutil.ReadFile(markText.FontFile)
	if err != nil {
		log.Println("读取字体数据出错")
		log.Println(err)
		return nil, err
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println("转换字体样式出错")
		log.Println(err)
		return nil, err
	}

	vg.AddFont("cn_font", font)
	fontStyle, err := vg.MakeFont("cn_font", vg.Inch*markText.Size)
	if err != nil {
		return nil, err
	}
	// 重复编写水印字体
	marktext := markText.Text
	unitText := marktext

	markTextWidth := fontStyle.Width(marktext)
	for markTextWidth <= diagonal {
		marktext += markText.Space + unitText
		markTextWidth = fontStyle.Width(marktext)
	}
	// 设置水印字体的颜色
	rgba := markText.Color
	c.SetColor(color.RGBA{rgba[0], rgba[1], rgba[2], rgba[3]})
	// 设置 0 到 π/2 之间的随机角度
	c.Rotate(markText.Angle * math.Pi / 2)

	// 设置每行水印的高度并添加水印
	// 一个字体的高度
	lineHeight := fontStyle.Extents().Height * markText.Linewidth
	for offset := -2 * diagonal; offset < 2*diagonal; offset += lineHeight {
		c.FillString(fontStyle, vg.Point{X: 0, Y: offset}, marktext)
	}

	// 画布写入新图片
	// 使用buffer去转换
	jc := vgimg.PngCanvas{Canvas: c}
	buff := new(bytes.Buffer)
	jc.WriteTo(buff)
	img, _, err = image.Decode(buff)
	if err != nil {
		return nil, err
	}

	// 得到图像的中心点
	ctp := int(diagonal * vgimg.DefaultDPI / vg.Inch / 2)

	// 切出打水印的图像
	size := bounds.Size()
	bounds = image.Rect(ctp-size.X/2, ctp-size.Y/2, ctp+size.X/2, ctp+size.Y/2)
	rv := image.NewRGBA(bounds)
	draw.Draw(rv, bounds, img, bounds.Min, draw.Src)
	return rv, nil
}

// WaterMark用于在图像上添加水印
func WaterMark_text_one(img image.Image, markText MarkText) (image.Image, error) {
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

	// 制作一个 fontstyle ，宽度为英寸,字体 Courier 标准的等宽度字体
	// 读字体数据
	fontBytes, err := ioutil.ReadFile(markText.FontFile)
	if err != nil {
		log.Println("读取字体数据出错")
		log.Println(err)
		return nil, err
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println("转换字体样式出错")
		log.Println(err)
		return nil, err
	}

	vg.AddFont("cn_font", font)
	fontStyle, err := vg.MakeFont("cn_font", vg.Inch*markText.Size)
	if err != nil {
		return nil, err
	}
	// 重复编写水印字体
	marktext := markText.Text
	// 设置水印字体的颜色
	rgba := markText.Color
	c.SetColor(color.RGBA{rgba[0], rgba[1], rgba[2], rgba[3]})
	//设置文字坐标间距
	c.FillString(fontStyle, vg.Point{X: 685, Y: 1005}, marktext)
	// 画布写入新图片
	// 使用buffer去转换
	jc := vgimg.PngCanvas{Canvas: c}
	buff := new(bytes.Buffer)
	jc.WriteTo(buff)
	img, _, err = image.Decode(buff)
	if err != nil {
		return nil, err
	}

	// 得到图像的中心点
	ctp := int(diagonal * vgimg.DefaultDPI / vg.Inch / 2)

	// 切出打水印的图像
	size := bounds.Size()
	bounds = image.Rect(ctp-size.X/2, ctp-size.Y/2, ctp+size.X/2, ctp+size.Y/2)
	rv := image.NewRGBA(bounds)
	draw.Draw(rv, bounds, img, bounds.Min, draw.Src)
	return rv, nil
}

// MarkingPicture 给一个文件路径的图片加文字水印
func MarkingPicture(filepath string, text MarkText) (image.Image, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	img, err = WaterMark(img, text)
	if err != nil {
		return nil, err
	}
	return img, nil
}

/*
只加单个水印文字不重复
*/
func MarkingPicture_Text_One(filepath string, text MarkText) (image.Image, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	img, err = WaterMark_text_one(img, text)
	if err != nil {
		return nil, err
	}
	return img, nil
}
func ImageToBuffer(img image.Image, ext string) (rv *bytes.Buffer, err error) {
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
