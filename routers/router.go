package routers

import (
	"Scheduler/controllers/login"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &login.Login{})
    beego.Router("/Login", &login.Login{})
    beego.Router("/Success", &login.Success{})
}
