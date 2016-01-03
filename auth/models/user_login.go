package models
import (
	"time"
)

type UserStatus int8
const(
	NEW_STATUS UserStatus = 1
	ACTIVE_STATUS UserStatus = 2
	INACTIVE_STATUS UserStatus = 3
)

type UserLogin struct {
	Id int64
	Email string `orm:"unique" json:"email"`
	UseName string `orm:"unique;null"`
	MobilePhone string `orm:"unique;null" json:"mobile"`
	Password string `json:"password"`
	Status UserStatus `json:"status"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime" json:"createAt"`
	UpdateAt time.Time `orm:"auto_now;type(datetime" json:"updateAt"`
}




