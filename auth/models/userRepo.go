package models
import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type UserRepo struct {
	Collection []*UserLogin
}

func (userRepo *UserRepo) New() *UserLogin {
	newUserRef := new(UserLogin)
	mailingAddrRef := new(Address)
	billingAddrRef := new(Address)
	profileRef := new(UserProfile)

	profileRef.BillingAddress = billingAddrRef
	profileRef.MailingAddress = mailingAddrRef

	newUserRef.Profile = profileRef
	userRepo.Collection = append(userRepo.Collection, newUserRef)

	return newUserRef
}

func (userRepo *UserRepo) LoadChildren() error {
	for _, userRef := range userRepo.Collection {
		if (userRef.Id > 0) {
			profileRef := userRef.Profile
			mailingAddrRef := profileRef.MailingAddress
			billingAddrRef := profileRef.BillingAddress
			fmt.Println(profileRef)
			fmt.Println(mailingAddrRef)
			fmt.Println(billingAddrRef)
			if err := profileRef.Read(); err != nil {
				return err
			}
			if err := mailingAddrRef.Read(); err != nil {
				return err
			}
			if err := billingAddrRef.Read(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (userRepo *UserRepo) Save() error {

	for _, userRef := range userRepo.Collection {
		profileRef := userRef.Profile
		mailingAddrRef := profileRef.MailingAddress
		billingAddrRef := profileRef.BillingAddress

		if (userRef.Id <= 0) {
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
			profileRef.BillingAddress = billingAddrRef
			profileRef.MailingAddress = mailingAddrRef
			userRef.Profile = profileRef
			o.Commit()
		} else if userRef.IsChanged || profileRef.IsChanged || mailingAddrRef.IsChanged || billingAddrRef.IsChanged {
			o := orm.NewOrm()
			err := o.Begin()
			if err != nil {
				return err
			}
			if mailingAddrRef.IsChanged {
				if err = mailingAddrRef.Update(o); err != nil {
					o.Rollback()
					return err
				}
			}
			if billingAddrRef.IsChanged {
				if err = billingAddrRef.Update(o); err != nil {
					o.Rollback()
					return err
				}
			}
			if profileRef.IsChanged {
				if err = profileRef.Update(o); err != nil {
					o.Rollback()
					return err
				}
			}
			if userRef.IsChanged {
				if err = userRef.Update(o); err != nil {
					o.Rollback()
					return err
				}
			}
			o.Commit()
		}
	}
	return nil
}

