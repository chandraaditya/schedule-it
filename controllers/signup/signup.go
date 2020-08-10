package signup

import (
	"Scheduler/models/user"
	"github.com/astaxie/beego"
)

type Signup struct {
	beego.Controller
}

func (c *Signup) Get() {
	c.TplName = "signup/signup.html"
	return
}

func (c *Signup) Post() {
	var User user.User

	User.HouseNo = c.Input().Get("houseno")
	User.Name = c.Input().Get("username")
	User.Email = c.Input().Get("email")
	User.PlainPass = c.Input().Get("password")
	repass := c.Input().Get("repassword")
	if repass != User.PlainPass {
		c.Data["HouseNo"] = User.HouseNo
		c.Data["Name"] = User.Name
		c.Data["Email"] = User.Email
		c.Data["Error"] = "Passwords do not match, please try again."
		c.TplName = "signup/signup.html"
		return
	}
	err := User.CreateNewUser()
	if err != nil {
		if err.Error() == "ea1001" {
			c.Data["Error"] = "User with email already exists, please proceed to login."
			c.TplName = "signup/signup.html"
			return
		}
		c.Data["Error"] = "Error, please try again."
		c.TplName = "signup/signup.html"
		return
	}
	//TODO go to dashboard
	c.Redirect("/Success", 302)
	return
}
