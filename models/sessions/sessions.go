package sessions

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"log"
)

var GlobalSession *session.Manager
var Err error

func init() {
	Manager := &session.ManagerConfig{
		CookieName:     beego.AppConfig.String("appname"),
		Gclifetime:     2147483647,
		Secure:         true,
		CookieLifeTime: 2147483647,
	}
	GlobalSession, Err = session.NewManager("memory", Manager)
	if Err != nil {
		log.Println("GS Error: ", Err)
	}
	return
}
