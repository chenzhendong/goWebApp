package models
import (
	"github.com/astaxie/beego/orm"
	"github.com/satori/go.uuid"
)

var UserRepo = make(map[string]*UserLogin)

func New() (string, *UserLogin, error) {
	uuid := uuid.NewV4()
	newUserRef := new(UserLogin)
	UserRepo[uuid] = newUserRef

	mailingAddrRef := new(Address)
	billingAddrRef := new(Address)
	profileRef := new(UserProfile)

	profileRef.BillingAddress = billingAddrRef
	profileRef.MailingAddress = mailingAddrRef

	newUserRef.Profile = profileRef

	return uuid.String(), newUserRef
}

func Get(uuid string) (*UserLogin) {
	return UserRepo[uuid]
}

func Persist() error {

	for _, userRef := range UserRepo {
		profileRef := userRef.Profile
		mailingAddrRef := profileRef.MailingAddress
		billingAddrRef := profileRef.BillingAddress

		if(userRef.Id <= 0){
			o := orm.NewOrm()
			err := o.Begin()
			if err != nil {
				return err
			}
			if err = mailingAddrRef.Insert(o); err != nil {
				o.Rollback()
				return err
			}
			if err = billingAddrRef.Insert(o); err != nil {
				o.Rollback()
				return err
			}
			if err = profileRef.Insert(o); err != nil {
				o.Rollback()
				return err
			}
			if err = userRef.Insert(o); err != nil {
				o.Rollback()
				return err
			}
			o.Commit()
		} else if userRef.IsChanged || profileRef.IsChanged || mailingAddrRef.IsChanged || billingAddrRef.IsChanged {
			o := orm.NewOrm()
			err := o.Begin()
			if err != nil {
				return err
			}

		}
	}
	return nil
}

