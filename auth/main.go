package main

import (
	_ "goWebApp/auth/docs"
	_ "goWebApp/auth/routers"
	"github.com/astaxie/beego"

)

func init() {

}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

