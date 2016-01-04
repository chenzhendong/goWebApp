package models
import (
	"github.com/astaxie/beego/orm"
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
	IsChanged bool `orm:"-"`
}

func (m *UserLogin) Insert(o orm.Ormer) error {
	if _, err := o.Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *UserLogin) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *UserLogin) Update(o orm.Ormer, fields ...string) error {
	if _, err := o.Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *UserLogin) Delete(o orm.Ormer) error {
	if _, err := o.Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *UserLogin) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
