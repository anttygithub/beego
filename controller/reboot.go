package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/beegoForWeDOS/models"
)

type RebootController struct {
	beego.Controller
}

func (this *RebootController) Get() {
	this.TplName = "reboot.tpl"
}

func (this *RebootController) Post() {
	var user models.User
	inputs := this.Input()
	//fmt.Println(inputs)
	user.Username = inputs.Get("username")
	user.Pwd = inputs.Get("pwd")
	//	err := models.SaveUser(user)
	err := models.Check(user)
	if err == nil {
		this.TplName = "success.tpl"
	} else {
		fmt.Println(err)
		this.TplName = "error.tpl"
	}
}
