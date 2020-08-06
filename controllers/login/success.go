package login

import "github.com/astaxie/beego"

type Success struct {
	beego.Controller
}

func (c *Success) Get() {
	c.TplName = "login/success.html"
	return
}
