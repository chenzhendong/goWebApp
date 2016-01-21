package models
import (
	"github.com/golang/groupcache/lru"
	"strconv"
	"github.com/jinzhu/gorm"
	"fmt"
)

type Repository struct {

}

var cache *lru.Cache

func init() {
	fmt.Println("Init Repository of Users ....")
	cache = lru.New(1000)
}

func getRepositoryId(user User) string {
	if user.ID > 0 {
		return "User::" + strconv.FormatUint(user.ID, 19)
	} else {
		return ""
	}
}

func clone(user User) User{
	newUser := User{}
	newUser = user
	newUser.Profile.Addresses = append([]Address(nil), user.Profile.Addresses...)
	return newUser
}

func (repo Repository) Get(id uint64) (User,error) {
	db := DB
	user := User{ID: id}
	key := getRepositoryId(user)
	value, ok := cache.Get(key)

	if ok {
		v, _ := value.(User)
		user = clone(v)
	} else {
		db = db.First(&user)
		if db.Error != nil {
			return user, db.Error
		} else {
			profileRef := &user.Profile
			db = db.Model(&user).Related(profileRef, "Profile").First(profileRef)
			addressesRef := &profileRef.Addresses
			db = db.Model(profileRef).Related(addressesRef).Find(addressesRef)
			cache.Add(key, clone(user))
		}
	}
	return user, nil
}

func (repo Repository) GetAll() []User {
	var users = []User{}
	DB.Model(&User{}).Find(&users)
	return users
}

func (repo Repository) saveOneUser(userRef *User, db *gorm.DB) error {
	if userRef.ID > 0 {
		db = db.Save(userRef)
		profileRef := &userRef.Profile
		db = db.Save(profileRef)
		addresses := profileRef.Addresses
		for _, address := range addresses {
			db = db.Save(&address)
		}
	} else {
		db = db.Create(userRef)
	}

	if( db.Error != nil){
		return db.Error
	} else {
		cacheUser := clone(*userRef)
		key := getRepositoryId(cacheUser)
		cache.Add(key, cacheUser)
		return nil
	}
}


func (repo Repository) SaveUser(user User) (User, error) {
	tx := DB.Begin()
	if err:=repo.saveOneUser(&user, tx); err != nil {
		tx.Rollback()
		return user, err
	} else {
		tx.Commit()
		return user, nil
	}
}

func (repo Repository) SaveUsers(users []User) ([]User, error) {
	tx := DB.Begin()
	for _, user := range users {
		if err:=repo.saveOneUser(&user, tx); err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()
	return users, nil
}



func (repo Repository) FindAll(db *gorm.DB) ([]User, error) {
	var users = make([]User, 0)
	db = db.Find(&users)

	if db.Error != nil {
		return users, db.Error
	} else {
		for idx, user := range users {
			users[idx], _ = repo.Get(user.ID)
		}
		return users, nil
	}
}

func (repo Repository) FindOne(db *gorm.DB) User {
	var user = User{}
	db.First(&user)
	user,_ = repo.Get(user.ID)
	return user
}

func (repo Repository) GetQueryBuilder() *gorm.DB {
	dbRef := DB.Model(&User{})
	return dbRef
}
