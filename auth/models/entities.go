package models
import (
	"time"
)

type UserStatus int8
const(
	UNDEFINED_STATUS UserStatus = 0
	NEW_STATUS UserStatus = 1
	PENDING_STATUS UserStatus = 2
	ACTIVE_STATUS UserStatus = 3
	INACTIVE_STATUS UserStatus = 4
)

type AddressType int8
const(
	UNDEFINED_ADDRESS AddressType = 0
	MAILING_ADDRESS AddressType = 1
	BILLING_ADDRESS AddressType = 2
)


type User struct {
	ID        uint64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Email string `sql:"not null;unique;index:idx_user_email" json:"email"`
	UserName string `sql:"index:idx_user_username" json:"userName"`
	MobilePhone string `sql:"index:idx_user_mobile" json:"mobile"`
	Password string `json:"password"`
	Status UserStatus `json:"status;default(0)"`
	Profile Profile
}

type Profile struct {
	ID        uint64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Addresses []Address `json:"addresses"`
	AddressID int64
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	MiddleName string `json:"middleName"`
	BirthDate time.Time `json:"birthDate"`
	Phone string `json:"phone"`
	UserID int64
}

type Address struct {
	ID        uint64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	ProfileID int64 `json:"profileId"`
	Attn string  `json:"attn"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	City string `json:"city"`
	StateProvince string `json:"stateProvince"`
	PostalCode string `json:"postalCode"`
	Phone string `json:"phone"`
	Country string `json:"country"`
	AddressType AddressType `json:"addressType"`
}

