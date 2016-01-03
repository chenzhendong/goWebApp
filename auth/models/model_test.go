package models
import (
	"testing"
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/astaxie/beego"
)

func setup()  {
	beego.AppConfig.Set("dbhost", "localhost")
	beego.AppConfig.Set("dbport", "5432")
	beego.AppConfig.Set("dbuser", "nodeframe")
	beego.AppConfig.Set("dbpassword", "nodeframe")
	beego.AppConfig.Set("dbname", "nodeframe")
}

func TestOrm(t *testing.T)  {
	setup()
	o := orm.NewOrm()
	userLogin := UserLogin{Email: "abc@test.com"}
	id, err := o.Insert(&userLogin)

	fmt.Printf("ID: %d, ERR: %v\n", id, err)
}

