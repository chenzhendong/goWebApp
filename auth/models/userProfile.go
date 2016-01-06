package models
import (
	"time"
	"github.com/astaxie/beego/orm"
)

type UserProfile struct {
	Id int64 `json:"id"`
	MailingAddress *Address `orm:"null;rel(one)" json:"mailingAddress"`
	BillingAddress *Address `orm:"null;rel(one)" json:"billingAddress"`
	FirstName string `orm:"null" json:"firstName"`
	LastName string `orm:"null" json:"lastName"`
	MiddleName string `orm:"null" json:"middleName"`
	BirthDate time.Time `orm:"null;type(date)" json:"birthDate"`
	HomePhone string `orm:"null" json:"homePhone"`
	WorkPhone string `orm:"null" json:"workPhone"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"createdAt"`
	UpdateAt time.Time `orm:"auto_now;type(datetime)" json:"updatedAt"`
	UserLogin *UserLogin `orm:"reverse(one)"`
	IsChanged bool `orm:"-"`
}


func (m *UserProfile) Insert(o orm.Ormer) error {
	if id, err := o.Insert(m); err != nil {
		return err
	}else {
		m.Id = id
		o.Read(m)
		return nil
	}
}

func (m *UserProfile) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *UserProfile) Update(o orm.Ormer, fields ...string) error {
	if _, err := o.Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *UserProfile) Delete(o orm.Ormer) error {
	if _, err := o.Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *UserProfile) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
