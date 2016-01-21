package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["goWebApp/users/controllers:AddressController"] = append(beego.GlobalControllerRouter["goWebApp/users/controllers:AddressController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["goWebApp/users/controllers:AddressController"] = append(beego.GlobalControllerRouter["goWebApp/users/controllers:AddressController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goWebApp/users/controllers:AddressController"] = append(beego.GlobalControllerRouter["goWebApp/users/controllers:AddressController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goWebApp/users/controllers:AddressController"] = append(beego.GlobalControllerRouter["goWebApp/users/controllers:AddressController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["goWebApp/users/controllers:AddressController"] = append(beego.GlobalControllerRouter["goWebApp/users/controllers:AddressController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["goWebApp/users/controllers:ProfileController"] = append(beego.GlobalControllerRouter["goWebApp/users/controllers:ProfileController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["goWebApp/users/controllers:ProfileController"] = append(beego.GlobalControllerRouter["goWebApp/users/controllers:ProfileController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goWebApp/users/controllers:ProfileController"] = append(beego.GlobalControllerRouter["goWebApp/users/controllers:ProfileController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goWebApp/users/controllers:ProfileController"] = append(beego.GlobalControllerRouter["goWebApp/users/controllers:ProfileController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["goWebApp/users/controllers:ProfileController"] = append(beego.GlobalControllerRouter["goWebApp/users/controllers:ProfileController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["goWebApp/users/controllers:UserController"] = append(beego.GlobalControllerRouter["goWebApp/users/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["goWebApp/users/controllers:UserController"] = append(beego.GlobalControllerRouter["goWebApp/users/controllers:UserController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goWebApp/users/controllers:UserController"] = append(beego.GlobalControllerRouter["goWebApp/users/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goWebApp/users/controllers:UserController"] = append(beego.GlobalControllerRouter["goWebApp/users/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["goWebApp/users/controllers:UserController"] = append(beego.GlobalControllerRouter["goWebApp/users/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
