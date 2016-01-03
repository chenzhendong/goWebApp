package models
import (
	"time"
)

type UserStatus int8
const(
	NEW_STATUS UserStatus = 0
	PENDING_STATUS UserStatus = 1
	ACTIVE_STATUS UserStatus = 2
	INACTIVE_STATUS UserStatus = 3
)

type AddressType int8
const(
	MAILING_ADDRESS UserStatus = 1
	BILLING_ADDRESS UserStatus = 2
)


type UserLogin struct {
	Id int64 `json:"id"`
	Email string `orm:"unique" json:"email"`
	UseName string `orm:"unique;null"`
	MobilePhone string `orm:"unique;null" json:"mobile"`
	Password string `json:"password"`
	Status UserStatus `json:"status"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"createdAt"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updatedAt"`
	Profile *UserProfile `orm:"null;rel(one)"`
	Sessions []*Session `orm:"reverse(many)"`
}


type UserProfile struct {
	Id int64 `json:"id"`
	MailingAddress *Address `orm:"null;rel(one)" json:"mailingAddress"`
	BillingAddress *Address `orm:"null;rel(one)" json:"billingAddress"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	MiddleName string `json:"middleName"`
	BirthDate time.Time `json:"birthDate"`
	HomePhone string `json:"homePhone"`
	WorkPhone string `json:"workPhone"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"createdAt"`
	UpdateAt time.Time `orm:"auto_now;type(datetime)" json:"updatedAt"`
	UserLogin *UserLogin `orm:"reverse(one)"`
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
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"createdAt"`
	UpdateAt time.Time `orm:"auto_now;type(datetime)" json:"updatedAt"`
	UserProfile *UserProfile  `orm:"reverse(one)" json:"userProfile"`
}

type Session struct {
	Id int64 `json:"id"`
	SessionKey string `json:"sessionKey"`
	IpAddress string `json:"ipAddress"`
	ValidThrough time.Time `json:"validThrough"`
	CreatedAt time.Time `orm:"created" json:"createdAt"`
	UpdateAt time.Time `orm:"updated" json:"updatedAt"`
	UserLogin *UserLogin `orm:"rel(fk)" json:"userLogin"`
}




