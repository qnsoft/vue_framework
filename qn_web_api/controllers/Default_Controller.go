package controllers

import (
	"qnsoft/qn_web_api/controllers/Token"
)

/**
*信息实体
 */
type Default_Controller struct {
	Token.BaseController
}

/*
根目录
*/
func (this *Default_Controller) Get() {
	this.TplName = "web_root/pc_ui/index.html"
}

/*
wap目录
*/
func (this *Default_Controller) Get_Wap() {
	this.TplName = "web_root/wap_ui/index.html"
}
