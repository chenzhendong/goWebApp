package auth

import (
	"time"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"log"
)

type UserStatus int8
const(
	NEW_STATUS UserStatus = 1
	ACTIVE_STATUS UserStatus = 2
	INACTIVE_STATUS UserStatus = 3
)

type AddressType int8
const(
	MAILING_ADDRESS UserStatus = 1
	BILLING_ADDRESS UserStatus = 2
)

type UserLogin struct {
	Id int64 `xorm:"autoincr" json:"id"`
	Email string `xorm:"unique notnull" json:"email"`
	MobilePhone string `xorm:"unique" json:"mobile"`
	Password string `json:"password"`
	Status UserStatus `json:"status"`
	CreatedAt time.Time `xorm:"created" json:"createAt"`
	UpdateAt time.Time `xorm:"updated" json:"updateAt"`
}

type UserProfile struct {
	Id int64 `json:"id"`
	MailingAddress Address `xorm:"extends" json:"mailingAddress"`
	BillingAddress Address `xorm:"extends" json:"billingAddress"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	MiddleName string `json:"middleName"`
	BirthDate time.Time `json:"birthDate"`
	HomePhone string `json:"homePhone"`
	WorkPhone string `json:"workPhone"`
	CreatedAt time.Time `xorm:"created" json:"createAt"`
	UpdateAt time.Time `xorm:"updated" json:"updateAt"`
}

type Address struct {
	Id int64 `json:"id"`
	Attn string  `json:"attn"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	City string `json:"city"`
	StateProvince string `json:"stateProvince"`
	PostalCode string `json:"postalCode"`
	Phone string `json:"phone"`
	Country string `json:"country"`
	AddressType AddressType `json:"addressType"`
	CreatedAt time.Time `xorm:"created" json:"createAt"`
	UpdateAt time.Time `xorm:"updated" json:"updateAt"`
	UserProfile UserProfile  `json:"userProfile"`
}

type Session struct {
	id int64 `json:"id"`
	SessionKey string `json:"sessionKey"`
	ValidThrough time.Time `json:"validThrough"`
	CreatedAt time.Time `xorm:"created" json:"createAt"`
	UpdateAt time.Time `xorm:"updated" json:"updateAt"`
	UserLogin UserLogin `json:"userLogin"`
}

var X *xorm.Engine

func init(){
	var err error
	X, err = xorm.NewEngine("postgres", "user=postgres password='dyslmt' dbname=nodeframe sslmode=disable")
	if err != nil {
		log.Fatalf("Fail to create engine: %v", err)
	}

	X.DropTables(new(UserLogin), new(UserProfile), new(Address))

	if err = X.Sync(new(UserLogin), new(UserProfile), new(Address)); err!=nil {
		log.Fatalf("Fail to sync database: %v", err)
	}
}



