package models
import (
	"github.com/golang/groupcache/lru"
	"strconv"
)

type Repository struct {

}

var cache *lru.Cache

func init()  {
	cache = lru.New(1000)
}

func getRepositoryId(user *UserLogin) string {
	if user.ID > 0 {
		return "UserLogin::" + strconv.FormatInt(user.ID, 19)
	} else {
		return ""
	}
}

func (repo Repository) Get (user *UserLogin){
	key := getRepositoryId(user)
	value, ok := cache.Get(key)

	if ok {
		v,_ := value.(UserLogin)
		user = &v
	} else {

	}

}


/*type UserRepo struct {
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
			profileRef := new(UserProfile)
			if err := profileRef.Query().Filter("user_login_id", user.Id).One(profileRef); err != nil {
				Log.Error("Failed to get just one profile of a user, could be a data integration problem (user_login tablo 1-1 with user_profile table)", err)
				return err
			}
			user.Profile = profileRef

			var addresses []*Address
			adrRef := new(Address)
			if _, err := adrRef.Query().Filter("profile_id", profileRef.Id).All(&addresses); err != nil {
				Log.Error("Failed to get address list for profile", err)
				return err
			}
			profileRef.Addresses = addresses
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
			if err = userRef.Insert(o); err != nil {
				Log.Error("Error on insert user", userRef)
				o.Rollback()
				return err
			}
			profileRef.UserLoginId = userRef.Id
			if err = profileRef.Insert(o); err != nil {
				Log.Error("Error on insert profile", profileRef)
				o.Rollback()
				return err
			}
			for _, addrRef := range addresses {
				addrRef.ProfileId = profileRef.Id
				if err := addrRef.Insert(o); err != nil {
					Log.Error("Error on insert an address", addrRef)
					o.Rollback()
					return err
				}
			}
			o.Commit()


			userRef.Profile = profileRef

			profileRef.Addresses = addresses
			for _, addr := range addresses {
				addr.ProfileId = profileRef.Id
			}
		}
	}


	for _, userRef := range userRepo.QueryEntries {
		profileRef := userRef.Profile
		addresses := userRef.Profile.Addresses

		o := orm.NewOrm()
		err := o.Begin()

		if userRef.IsChanged || profileRef.IsChanged {
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
		}

		for _, addr := range addresses {
			if addr.IsChanged {
				if err = addr.Update(o); err != nil {
					o.Rollback()
					return err
				}
			}
		}

		userRef.Profile = profileRef
		profileRef.Addresses = addresses
		o.Commit()
	}

	return nil
}*/

