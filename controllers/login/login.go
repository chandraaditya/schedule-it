package login

import (
	"Scheduler/models/sessions"
	"Scheduler/models/user"
	"github.com/astaxie/beego"
	"log"
)

type Login struct {
	beego.Controller
}

func (c *Login) Get() {
	c.TplName = "login/login.html"
	return
}

func (c *Login) Post() {
	var User user.User

	User.Email = c.GetString("email")
	User.PlainPass = c.GetString("password")
	log.Println("entered data:", User.Email, User.PlainPass)

	err := User.Authenticate()
	if err != nil {
		if err.Error() == "ea1001: error authenticating User" || err.Error() == "el1004" {
			c.Data["Error"] = "Please check your login details."
			c.TplName = "login/login.html"
			return
		}
		if sessions.Err != nil {
			c.Data["Error"] = "Unexpected error occurred, please try again later."
			log.Println("Error: session error") //TODO need to add error code
			c.TplName = "error/error.html"
			return
		}
		c.Data["Error"] = "Unexpected error occurred, please try again later."
		log.Println("Error: session error") //TODO need to add error code
		c.TplName = "error/error.html"
		return
	}

	log.Println(User.Email) //TODO successful login

	c.Redirect("/Success", 302)
	return
}
