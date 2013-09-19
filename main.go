package main

import (
	"github.com/astaxie/beego"
	"github.com/ginuerzh/icecar/controllers"
	"github.com/ginuerzh/icecar/models"
	//"icecar/filters"
	"labix.org/v2/mgo"
)

func main() {
	beego.Router("/account/register", &controllers.UserController{}, "get,post:Register")
	beego.Router("/account/login", &controllers.UserController{}, "get,post:Login")
	beego.Router("/user/logout", &controllers.UserController{}, "get:Logout")
	beego.Router("/user/getInfo", &controllers.UserController{}, "get,post:UserInfo")
	beego.Router("/file/upload", &controllers.FileController{}, "post:Upload")

	//beego.FilterPrefixPath("/user/getInfo", filters.AccessFilter)

	session, err := mgo.Dial(beego.AppConfig.String("mongourl"))
	if err != nil {
		panic(err)
		return
	}

	models.DB = session.DB(beego.AppConfig.String("mongodb"))
	defer session.Close()

	beego.Debug("start server")
	beego.Run()
}
