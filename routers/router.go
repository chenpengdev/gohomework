package routers

import (
	"github.com/chenpengdev/gohomework/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
