package controllers

import (
	"goWebApp/auth/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"errors"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

var repo = models.RepoInstance

func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
}

func (u *UserController) handleError(code int, err error)  {
	u.Ctx.Output.SetStatus(code)
	u.Data["json"] = err.Error()
}

// @Title createUser
// @Description Create a new users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 400 {object} error
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &user); err != nil {
		u.handleError(400, err)
	} else {
		if user, err = repo.SaveUser(user); err != nil {
			u.handleError(400, err)
		} else{
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Get Single User
// @Description Get a User by ID
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 400 {object} error
// @Failure 403 :id is empty
// @router /:id [get]
func (u *UserController) GetOne() {
	idStr := u.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	user,err := repo.Get(uint64(id))
	if err != nil {
		u.handleError(400, err)
	} else {
		u.Data["json"] = user
	}
	u.ServeJSON()
}

// @Title GetAll
// @Description Get All Users
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.User
// @Failure 400 {object} error
// @Failure 403
// @router / [get]
func (c *UserController) GetAll() {
	var sortby []string
	var order []string
	var limit int64 = 20
	var offset int64 = 0

	db := repo.GetQueryBuilder()

	// limit: 20 (default is 20)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}

	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}

	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}

	db = db.Offset(offset).Limit(limit)
	for idx,col := range sortby {
		col += " "
		col += order[idx]
		db = db.Order(col)
	}

	l, err := repo.FindAll(db)

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// @Title Update
// @Description Update the User
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.User	true		"body for User content"
// @Success 200 {object} models.User
// @Failure 400 {object} error
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.User{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if v.ID >0 && uint64(id) != v.ID {
			c.handleError(400, errors.New("The user ID in URL is not match the ID in body."))
		} else {
			if v, err := repo.SaveUser(v); err == nil {
				c.Data["json"] = v
			} else {
				c.handleError(400, err)
			}
		}
	} else {
		c.handleError(400, err)
	}
	c.ServeJSON()
}

