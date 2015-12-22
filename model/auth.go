package model

import (
	"log"
	"time"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type UserLogin struct {
	Id int64
	UserName string
	Email string
	MobilePhone string
	CreatedAt time.Time
	UpdateAt time.Time
}



