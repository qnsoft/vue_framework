package PicHelper

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"qnsoft/qn_web_api/utils/ErrorHelper"

	"github.com/BurntSushi/graphics-go/graphics"
	"github.com/astaxie/beego"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

var font *truetype.Font

//画直线
func drawline(x0, y0, x1, y1 int, img *image.RGBA, c color.Color) {
	dx := math.Abs(float64(x1 - x0))
	dy := math.Abs(float64(y1 - y0))
	sx, sy := 1, 1
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}
	err := dx - dy
	for {
		img.Set(x0, y0, c)
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

//填充亮色背景
func Fill(img *image.RGBA, option *DrawOption) {
	c := color.RGBA{}
	if option.RandomBG {
		c.R = uint8(200 + rand.Intn(55))
		c.G = uint8(200 + rand.Intn(55))
		c.B = uint8(200 + rand.Intn(55))
		c.A = uint8(255)
	} else {
		c = *option.BG
	}
	x := img.Rect.Size().X
	y := img.Rect.Size().Y
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			img.Set(i, j, c)
		}
	}
}

//填充暗色背景（供生成字体调用）
func FillDark(img *image.RGBA, option *DrawOption) {
	c := color.RGBA{}
	if option.RandomBG {
		c.R = uint8(rand.Intn(155))
		c.G = uint8(rand.Intn(155))
		c.B = uint8(rand.Intn(155))
		c.A = uint8(255)
	} else {
		c = *option.BG
	}
	x := img.Rect.Size().X
	y := img.Rect.Size().Y
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			img.Set(i, j, c)
		}
	}
}

//产生像素噪点
func Noise(img *image.RGBA, option *DrawOption) {
	h := img.Rect.Size().Y
	w := img.Rect.Size().X
	count := h * w * option.NoisePercent / 100
	for i := 0; i != count; i++ {
		c := color.RGBA{
			R: uint8(rand.Intn(255)),
			G: uint8(rand.Intn(255)),
			B: uint8(rand.Intn(255)),
			A: uint8(255),
		}
		img.Set(rand.Intn(w), rand.Intn(h), c)
	}
}

//产生垃圾线条
func NoiseLine(img *image.RGBA, option *DrawOption) {
	h := img.Rect.Size().Y
	w := img.Rect.Size().X
	for i := 0; i != option.NoiseLineNum; i++ {
		c := color.RGBA{
			R: uint8(rand.Intn(255)),
			G: uint8(rand.Intn(255)),
			B: uint8(rand.Intn(255)),
			A: uint8(255),
		}
		drawline(
			rand.Intn(w),
			rand.Intn(h),
			rand.Intn(w),
			rand.Intn(h),
			img,
			c)
	}
}

//绘画选项结构体
type DrawOption struct {
	//验证码内容
	Str string
	//随机验证码长度
	RandStrLen int
	//验证码图片宽度
	ImgW int
	//验证码图片高度
	ImgH int
	//随机背景(优先)
	RandomBG bool
	//指定背景颜色
	BG *color.RGBA
	//垃圾像素百分比
	NoisePercent int
	//垃圾线条数目
	NoiseLineNum int
	//最小随机字体大小值
	FontSizeMin float64
	//最大随机字体大小值
	FontSizeMax float64
	//随机字体颜色(优先)
	RandomFontColor bool
	//字体颜色
	FontColor *color.RGBA
	//最小随机字体间距值
	MarginMin float64
	//最大随机字体大小值
	MarginMax float64
	//字体旋转角度最小值
	AngleMin float64
	//字体旋转角度最大值
	AngleMax float64
}

//画字体
func DrawString(img *image.RGBA, option *DrawOption) {
	fontcolor := image.NewRGBA(image.Rect(0, 0, option.ImgW, option.ImgH))
	ctx := freetype.NewContext()
	ctx.SetDPI(72)
	ctx.SetFont(font)
	ctx.SetClip(img.Bounds())
	ctx.SetDst(img)
	pad := 0.0
	for _, v := range []rune(option.Str) {
		tmp := image.NewRGBA(img.Rect)
		//在指定范围内波动字体大小
		size := option.FontSizeMin + (option.FontSizeMax-option.FontSizeMin)*rand.Float64()
		//在指定范围内波动字体间距
		margin := option.MarginMin + (option.MarginMax-option.MarginMin)*rand.Float64()

		//建立蒙板供字体使用
		FillDark(fontcolor, &DrawOption{RandomBG: option.RandomFontColor, BG: option.FontColor})
		//设置字体渲染
		ctx.SetSrc(fontcolor)
		ctx.SetFontSize(size)
		ctx.SetDst(tmp)
		ctx.DrawString(string(v), freetype.Pt(
			10+int(pad),
			10+int(ctx.PointToFixed(size)>>6)))
		//取出已生成的字的坐标
		x0, x1, y0, y1 := GetFontRect(tmp)
		//给旋转角度预留额外的空间
		rect := image.Rect(x0-10, y0-10, x1+10, y1+10)
		//取出已生成的字
		char := tmp.SubImage(rect)
		chardst := image.NewRGBA(rect)
		//计算角度
		ang := option.AngleMin + (option.AngleMax-option.AngleMin)*rand.Float64()
		ang = math.Pi / 180 * ang
		//旋转
		graphics.Rotate(chardst, char.(draw.Image), &graphics.RotateOptions{ang})
		//合并
		draw.Draw(img, rect, chardst, chardst.Bounds().Min, draw.Over)
		pad += margin
	}
}

//取出已生成的字的坐标
func GetFontRect(tmp *image.RGBA) (int, int, int, int) {
	minh := -1
	for j := 1; j <= tmp.Rect.Size().Y; j++ {
		for i := 1; i <= tmp.Rect.Size().X; i++ {
			if tmp.At(i, j).(color.RGBA).A != 0 {
				minh = j
				break
			}
		}
		if minh != -1 {
			break
		}
	}
	minw := -1
	for i := 1; i <= tmp.Rect.Size().X; i++ {
		for j := 1; j <= tmp.Rect.Size().Y; j++ {
			if tmp.At(i, j).(color.RGBA).A != 0 {
				minw = i
				break
			}
		}
		if minw != -1 {
			break
		}
	}
	maxh := -1
	for j := tmp.Rect.Size().Y; j >= 1; j-- {
		for i := tmp.Rect.Size().X; i >= 1; i-- {
			if tmp.At(i, j).(color.RGBA).A != 0 {
				maxh = j
				break
			}
		}
		if maxh != -1 {
			break
		}
	}
	maxw := -1
	for i := tmp.Rect.Size().X; i >= 1; i-- {
		for j := 1; j <= tmp.Rect.Size().Y; j++ {
			if tmp.At(i, j).(color.RGBA).A != 0 {
				maxw = i
				break
			}
		}
		if maxw != -1 {
			break
		}
	}
	return minw, maxw, minh, maxh
}

//初始化整形
func DefInt(i *int, def int) {
	if *i <= 0 {
		*i = def
	}
}

//初始化两个浮点型(a理论上比b小)
func Def2Float64(a *float64, b *float64, def float64) {
	if *a < 0 {
		*a = 0
	}
	if *b < 0 {
		*b = 0
	}
	if *a > *b {
		*a, *b = *b, *a
	}
	if *b == 0 && *a == 0 {
		*a = def
		*b = def
	}
}

//新建验证码
func NewCPT(option *DrawOption) (io.Reader, string) {
	//验证码缺省长度为4
	DefInt(&option.RandStrLen, 4)
	//如果验证码为空，则自动生成
	if option.Str == "" {
		randStr := "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for i := 0; i < option.RandStrLen; i++ {
			n := rand.Intn(len(randStr))
			option.Str += randStr[n : n+1]
		}
	}
	//验证码图片宽度
	DefInt(&option.ImgW, 120)
	//验证码图片高度
	DefInt(&option.ImgH, 50)
	//如果指定背景颜色为空，则随机生成
	if option.BG == nil {
		option.RandomBG = true
	}
	//垃圾像素占图片百分比
	DefInt(&option.NoisePercent, 10)
	//垃圾线条数目
	DefInt(&option.NoiseLineNum, 10)
	//默认字体大小
	Def2Float64(&option.FontSizeMin, &option.FontSizeMax, 22)
	//默认字体间距
	Def2Float64(&option.MarginMin, &option.MarginMax, 22)
	//默认旋转角度
	Def2Float64(&option.AngleMin, &option.AngleMax, 30)
	//如果指定字体颜色为空，则随机生成
	if option.FontColor == nil {
		option.RandomFontColor = true
	}
	img := image.NewRGBA(image.Rect(0, 0, option.ImgW, option.ImgH))
	Fill(img, option)
	Noise(img, option)
	DrawString(img, option)
	NoiseLine(img, option)
	b := bytes.NewBuffer(make([]byte, 0))
	png.Encode(b, img.SubImage(img.Rect))
	return b, option.Str
}

/*
获取验证码
*/
func Get_VerifyCode(rw http.ResponseWriter, rq *http.Request) string {
	_font_path := beego.AppConfig.String("server_path::FontPath") + "lucon.ttf"
	ErrorHelper.LogInfo("字体路径", _font_path)
	fontfile, err := ioutil.ReadFile(_font_path)
	ErrorHelper.CheckErr(err)
	font, err = freetype.ParseFont(fontfile)
	ErrorHelper.CheckErr(err)
	r, str := NewCPT(&DrawOption{
		AngleMin:    0,  //最小角度
		AngleMax:    30, //最大角度
		FontSizeMin: 34, //最小字体
		FontSizeMax: 35, //最大字体
		MarginMin:   28, //最小边距
		MarginMax:   29, //最大边距
	})
	io.Copy(rw, r)
	return str
}
