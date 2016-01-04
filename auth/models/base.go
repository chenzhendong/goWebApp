package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)


func init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "5432"
	}
	orm.RegisterDataBase("default", "postgres", "host=" + dbhost + " user=" + dbuser + " password='" + dbpassword + "' dbname=" + dbname + " port=" + dbport + " sslmode=disable")
	orm.RegisterModel(new(UserLogin), new(UserProfile), new(Address), new(Session))
	orm.RunSyncdb("default", true, true)
}




