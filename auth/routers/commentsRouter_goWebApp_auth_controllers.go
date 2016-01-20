package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["goWebApp/auth/controllers:UserController"] = append(beego.GlobalControllerRouter["goWebApp/auth/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

}
