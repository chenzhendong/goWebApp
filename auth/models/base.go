package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB gorm.DB

func init() {
	dbhost := "localhost"
	dbport := "5432"
	dbuser := "nodeframe"
	dbpassword := "nodeframe"
	dbname := "nodeframe"
	if dbport == "" {
		dbport = "5432"
	}

	db, err := gorm.Open("postgres", "postgres", "host=" + dbhost + " user=" + dbuser + " password='" + dbpassword + "' dbname=" + dbname + " port=" + dbport + " sslmode=disable")
	if err != nil {

	}
	
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.SingularTable(true)

	db.DropTableIfExists(&UserLogin{}, &UserProfile{}, &Address{})
	db.CreateTable(&UserLogin{}, &UserProfile{}, &Address{})
	db.LogMode(true)

	DB = db
}

