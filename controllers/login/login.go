package login

import (
	"github.com/astaxie/beego"
	"log"
)

type Login struct {
	beego.Controller
}

func (c *Login) Get() {
	c.Data["Name"] = "Jesus"
	c.TplName = "login/login.html"
	return
}

func (c *Login) Post() {
	Username := c.GetString("username")
	Password := c.GetString("password")
	log.Println("Data:", Username, Password)
	c.Redirect("/Success", 302)
	return
}
