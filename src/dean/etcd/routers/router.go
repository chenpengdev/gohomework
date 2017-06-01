package routers

import (
	"dean/etcd/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/setkey", &controllers.MainController{}, "*:SetKey")
	beego.Router("/getkey", &controllers.MainController{}, "*:GetKey")
	beego.Router("/delkey", &controllers.MainController{}, "*:DelKey")
}
