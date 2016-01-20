package controllers

import (
	"goWebApp/auth/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

var repo = models.Repository{}

func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	repo.Save(&user)
	u.Data["json"] = user
	u.ServeJSON()
}


