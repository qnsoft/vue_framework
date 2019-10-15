package sys

import (
	"qnsoft/qn_web_api/controllers/api/sys"

	"github.com/astaxie/beego"
)

/* sys相关接口路由 */
func Sys_router() {
	//系统用户登录验证码
	beego.Router("/api/sys/verifyCode", &api.User_Controller{}, "*:VerifyCode")
	//系统用户登录
	beego.Router("/api/sys/login", &api.User_Controller{}, "post:Login")
	//系统用户修改密码
	beego.Router("/api/sys/Password", &api.User_Controller{}, "post:Password")
	//系统用户退出登录
	beego.Router("/api/sys/Logout", &api.User_Controller{}, "post:Logout")
	//系统菜单
	beego.Router("/api/sys/menu/nav", &api.Menu_Controller{}, "*:Nav")
	//菜单管理列表
	beego.Router("/api/sys/menu/list", &api.Menu_Controller{}, "*:List")
	//菜单选择
	beego.Router("/api/sys/menu/select", &api.Menu_Controller{}, "*:Select")
	//获取单条菜单信息
	beego.Router("/api/sys/menu/info/:menu_id", &api.Menu_Controller{}, "*:Info")
	//添加或修改菜单信息
	beego.Router("/api/sys/menu/edit", &api.Menu_Controller{}, "post:Edit")
	//删除菜单
	beego.Router("/api/sys/menu/delete", &api.Menu_Controller{}, "post:Delete")

	//系统用户信息
	beego.Router("/api/sys/user/info/:user_id", &api.User_Controller{}, "*:Info")
	//添加或修改系统用户信息
	beego.Router("/api/sys/user/edit", &api.User_Controller{}, "post:Edit")
	//删除系统用户信息
	beego.Router("/api/sys/user/delete", &api.User_Controller{}, "post:Delete")
	//系统管理员列表信息
	beego.Router("/api/sys/user/list", &api.User_Controller{}, "get,post:List")
	//系统角色列表信息
	beego.Router("/api/sys/role/list", &api.Role_Controller{}, "get,post:List")
	//系统角色列表信息选择
	beego.Router("/api/sys/role/select", &api.Role_Controller{}, "get:Select")
	//获取单条角色列表信息
	beego.Router("/api/sys/role/info/:role_id", &api.Role_Controller{}, "get:Info")
	//添加或修改角色信息
	beego.Router("/api/sys/role/edit", &api.Role_Controller{}, "post:Edit")
	//删除角色信息
	beego.Router("/api/sys/role/delete", &api.Role_Controller{}, "post:Delete")

	//系统设置信息
	beego.Router("/api/sys/config/list", &api.Config_Controller{}, "get,post:List")
	//获取单条系统设置信息
	beego.Router("/api/sys/config/info/:id", &api.Config_Controller{}, "get:Info")
	//添加或修改系统设置信息
	beego.Router("/api/sys/config/edit", &api.Config_Controller{}, "post:Edit")
	//删除系统设置信息
	beego.Router("/api/sys/config/delete", &api.Config_Controller{}, "post:Delete")
	//系统日志列表信息
	beego.Router("/api/sys/log/list", &api.Log_Controller{}, "get,post:List")

}
