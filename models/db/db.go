package db

import (
	"database/sql"
	"errors"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB
var Err error

func init() {
	Conn, Err = sql.Open(beego.AppConfig.String("DBType"), beego.AppConfig.String("DBUsername")+":"+beego.AppConfig.String("DBPassword")+"@tcp("+beego.AppConfig.String("DBHost")+":"+beego.AppConfig.String("DBPort")+")/"+beego.AppConfig.String("DBName"))
	if Err != nil {
		Err = errors.New("ed1001: error opening db")
	}
	Err = Conn.Ping()
	if Err != nil {
		Err = errors.New("ed1002: error connecting to db")
	}
	return
}
