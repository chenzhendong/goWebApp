package models
import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type UserRepo struct {
	NewEntries   []*UserLogin
	QueryEntries map[int64]*UserLogin
}

func NewUserRepo() *UserRepo {
	userRepoRef := new(UserRepo)
	userRepoRef.QueryEntries = make(map[int64]*UserLogin)
	return userRepoRef
}

func (userRepo *UserRepo) NewEntry() *UserLogin {
	newUserRef := new(UserLogin)
	profileRef := new(UserProfile)
	newUserRef.Profile = profileRef
	userRepo.NewEntries = append(userRepo.NewEntries, newUserRef)

	return newUserRef
}

func (userRepo *UserRepo) QueryBuilder() orm.QuerySeter {
	return new(UserLogin).Query()
}

func (userRepo *UserRepo) Get(query orm.QuerySeter) error {
	var users []*UserLogin
	if _, err := query.All(&users); err != nil {
		return err
	}

	for _, user := range users {
		if user != nil {

			userRepo.QueryEntries[user.Id] = user
		}
	}
	return nil
}

func (userRepo *UserRepo) Save() error {
	fmt.Println("Starting Persist Repository ...")

	for _, userRef := range userRepo.NewEntries {
		if userRef != nil {
			profileRef := userRef.Profile
			addresses := profileRef.Addresses

			o := orm.NewOrm()
			err := o.Begin()
			if err != nil {
				return err
			}
			for _, addrRef := range addresses {
				if err := addrRef.Insert(o); err != nil {
					Log.Error("Error on insert an address", addrRef)
					o.Rollback()
					return err
				}
			}
			if err = profileRef.Insert(o); err != nil {
				Log.Error("Error on insert profile", profileRef)
				o.Rollback()
				return err
			}
			if err = userRef.Insert(o); err != nil {
				Log.Error("Error on insert user", userRef)
				o.Rollback()
				return err
			}
			o.Commit()

			userRef.ProfileId = profileRef.Id
			profileRef.UserLoginId = userRef.Id
			userRef.Profile = profileRef

			profileRef.Addresses = addresses
			for _, addr := range addresses {
				addr.ProfileId = profileRef.Id
			}
		}
	}


	for _, userRef := range userRepo.QueryEntries {
		profileRef := userRef.Profile

		if userRef.IsChanged || profileRef.IsChanged {
			o := orm.NewOrm()
			err := o.Begin()
			if err != nil {
				return err
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

			for _, addr := range profileRef.Addresses {
				if addr.IsChanged {
					if err = addr.Update(o); err != nil {
						o.Rollback()
						return err
					}
				}
			}
			o.Commit()
		}

	}


	return nil
}

