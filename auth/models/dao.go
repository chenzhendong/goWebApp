package models
import "github.com/astaxie/beego/orm"

func Insert(m interface{}) error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func Read(m interface{}, fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func Update(m interface{}, fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func Delete(m interface{}) error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func Query(m interface{}) orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}