package signup

import "github.com/astaxie/beego"

type Signup struct {
	beego.Controller
}

func (c *Signup) Get() {
	c.TplName = "signup/signup.html"
	return
}

func (c *Signup) Post() {
	return
}
