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

	beego.GlobalControllerRouter["goWebApp/auth/controllers:UserController"] = append(beego.GlobalControllerRouter["goWebApp/auth/controllers:UserController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goWebApp/auth/controllers:UserController"] = append(beego.GlobalControllerRouter["goWebApp/auth/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

}
