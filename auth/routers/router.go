// @APIVersion 1.0.0
// @Title User Test API
// @Description beego has a very cool tools to autogenerate documents for your API

package routers

import (
	"goWebApp/auth/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

	)
	beego.AddNamespace(ns)
}
