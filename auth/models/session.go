package models
import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Session struct {
	Id int64 `json:"id"`
	SessionKey string `json:"sessionKey"`
	IpAddress string `json:"ipAddress"`
	ValidThrough time.Time `json:"validThrough"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"createdAt"`
	UpdateAt time.Time `orm:"auto_now;type(datetime)" json:"updatedAt"`
}


func (m *Session) Insert(o orm.Ormer) error {
	if id, err := o.Insert(m); err != nil {
		return err
	}else {
		m.Id = id
		o.Read(m)
		return nil
	}
}

func (m *Session) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Session) Update(o orm.Ormer, fields ...string) error {
	if _, err := o.Update(m, fields...); err != nil {
		return err
	}else {
		o.Read(m)
		return nil
	}
}

func (m *Session) Delete(o orm.Ormer) error {
	if _, err := o.Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Session) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}





