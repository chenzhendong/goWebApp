package models
import (
	"time"
	"github.com/astaxie/beego/orm"
)

type AddressType int8
const(
	UNDEFINED_ADDRESS AddressType = 0
	MAILING_ADDRESS AddressType = 1
	BILLING_ADDRESS AddressType = 2
)

type Address struct {
	Id int64 `json:"id"`
	Attn string  `orm:"null" json:"attn"`
	AddressLine1 string `orm:"null" json:"addressLine1"`
	AddressLine2 string `orm:"null" json:"addressLine2"`
	City string `orm:"null" json:"city"`
	StateProvince string `orm:"null" json:"stateProvince"`
	PostalCode string `orm:"null" json:"postalCode"`
	Phone string `orm:"null" json:"phone"`
	Country string `orm:"null" json:"country"`
	AddressType AddressType `orm:"default(0)" json:"addressType"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"createdAt"`
	UpdateAt time.Time `orm:"auto_now;type(datetime)" json:"updatedAt"`
	UserProfile *UserProfile  `orm:"reverse(one)" json:"userProfile"`
	IsChanged bool `orm:"-"`
}

func (m *Address) Insert(o orm.Ormer) error {
	if id, err := o.Insert(m); err != nil {
		return err
	} else {
		m.Id = id
		o.Read(m)
		return nil
	}
}

func (m *Address) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Address) Update(o orm.Ormer, fields ...string) error {
	if _, err := o.Update(m, fields...); err != nil {
		return err
	}else {
		o.Read(m)
		return nil
	}
}

func (m *Address) Delete(o orm.Ormer) error {
	if _, err := o.Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Address) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

