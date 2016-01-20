package models

import (
	"github.com/jinzhu/gorm"
	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
	"log"
)

var DB gorm.DB

func init() {
	DB = initPostgres()

	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	DB.DropTableIfExists(new(User), new(Profile), new(Address))
	DB.CreateTable(new(User), new(Profile), new(Address))
	DB.LogMode(true)
}

func initSqlite() gorm.DB  {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Failed to connect to Sqlite DB.", err)
	}
	return db
}

func initPostgres()  gorm.DB {

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
		log.Fatal("Failed to connect to Postgres DB.", err)
	}

	return db
}

