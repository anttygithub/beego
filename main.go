package main

import (
	//	"fmt"

	"github.com/astaxie/beego"
	"github.com/beegoForWeDOS/controller"
)

type MainController struct {
	beego.Controller
}

func main() {
	beego.SessionConfig.SessionOn = true
	//	beego.RegisterController("/", &controllers.IndexController{})
	beego.RegisterController("/", &controllers.IndexController{})
	beego.RegisterController("/reboot", &controllers.RebootController{})
	beego.Run()
}
