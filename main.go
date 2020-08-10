package main

import (
	_ "Scheduler/models/db"
	_ "Scheduler/models/sessions"
	_ "Scheduler/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
