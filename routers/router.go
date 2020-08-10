package routers

import (
	"Scheduler/controllers/login"
	"Scheduler/controllers/signup"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router(beego.AppConfig.String("Index"), &login.Login{})
	beego.Router(beego.AppConfig.String("Login"), &login.Login{})
	beego.Router(beego.AppConfig.String("Signup"), &signup.Signup{})
	//beego.Router(beego.AppConfig.String("Home")) TODO need to finish
}
