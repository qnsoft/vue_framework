package FileHelper

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"qnsoft/qn_web_api/utils/ErrorHelper"
)

/*
生成图片与图片层叠合成图
*/
func Pic_layer(Bg_pic, Ft_pic, New_pic string) string {
	//图片，网上随便找了一张
	img_Bg_pic, err1 := os.Open(Bg_pic) //背景图
	ErrorHelper.CheckErr(err1)          //打开图片出错
	ErrorHelper.LogInfo("打开背景图出错", err1)
	defer img_Bg_pic.Close()
	img, err2 := jpeg.Decode(img_Bg_pic)
	ErrorHelper.CheckErr(err2) //把图片解码为结构体时出错
	//水印,用的是我自己支付宝的二维码
	img_Ft_pic, err3 := os.Open(Ft_pic) //前景图或水印图
	ErrorHelper.CheckErr(err3)          //打开水印图片出错
	ErrorHelper.LogInfo("打开前景图出错", err3)

	// 获取前景图的类型
	// datatype, err2 := imgtype.Get(Ft_pic)
	// if err2 != nil {
	// 	ErrorHelper.LogInfo("不是图片文件")
	// } else {
	// 	// 根据文件类型执行响应的操作
	// 	switch datatype {
	// 	case `image/jpeg`:
	// 		ErrorHelper.LogInfo("这是JPG文件")
	// 	case `image/png`:
	// 		ErrorHelper.LogInfo("这是PNG文件")
	// 	}
	// }

	//ErrorHelper.LogInfo("前景图类型", &img_Ft_pic)
	defer img_Ft_pic.Close()
	//wmb_img, err4 := jpeg.Decode(img_Ft_pic) //前景图解码
	wmb_img, err4 := png.Decode(img_Ft_pic) //前景图解码
	ErrorHelper.CheckErr(err4)              //把水印图片解码为结构体时出错
	//把水印写在右下角，并向0坐标偏移10个像素
	//offset := image.Pt(img.Bounds().Dx()-wmb_img.Bounds().Dx()-10, img.Bounds().Dy()-wmb_img.Bounds().Dy()-10)
	ErrorHelper.LogInfo("前景图解码出错了", err4)
	//	offset := image.Pt(wmb_img.Bounds().Dx()+90, img.Bounds().Dy()-1330)
	offset := image.Pt(282, 298)
	b := img.Bounds()
	//根据b画布的大小新建一个新图像
	m := image.NewRGBA(b)
	//image.ZP代表Point结构体，目标的源点，即(0,0)
	//draw.Src源图像透过遮罩后，替换掉目标图像
	//draw.Over源图像透过遮罩后，覆盖在目标图像上（类似图层）
	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, wmb_img.Bounds().Add(offset), wmb_img, image.ZP, draw.Over)
	//生成新图片new.jpg,并设置图片质量
	imgw, err5 := os.Create(New_pic)
	ErrorHelper.LogInfo("合成出错了", err5)
	//jpeg.Encode(imgw, m, &jpeg.Options{jpeg.DefaultQuality})
	jpeg.Encode(imgw, m, &jpeg.Options{60}) //生成jpeg质量
	defer imgw.Close()
	return New_pic
}
