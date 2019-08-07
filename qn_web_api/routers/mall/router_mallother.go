package routers

import (
	"qnsoft/mall_online/controllers/mall_other/taobao"

	"github.com/astaxie/beego"
)

func Mallother() {
	//阿里商品列表
	beego.Router("/api/mallother/taobao/goods", &taobao.Mall_other_taobao_Controller{}, "post:Goods")
	//阿里妈妈推广券
	beego.Router("/api/mallother/taobao/Taobao_YouhuiQuan", &taobao.Mall_other_taobao_Controller{}, "post:Taobao_YouhuiQuan")
	//工具接口
	//beego.Router("/api/mallother/taobao/Get_xianjia", &taobao.Mall_other_taobao_Controller{}, "post:Get_xianjia")
}
