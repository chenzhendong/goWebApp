package models
import (
	"github.com/golang/groupcache/lru"
	"strconv"
	"github.com/jinzhu/gorm"
)

type Repository struct {

}

var cache *lru.Cache

func init()  {
	cache = lru.New(1000)
}

func getRepositoryId(user User) string {
	if user.ID > 0 {
		return "User::" + strconv.FormatUint(user.ID, 10)
	} else {
		return ""
	}
}

func (repo Repository) Get (id uint64) User{
	user := User{ID: id}
	key := getRepositoryId(user)
	value, ok := cache.Get(key)

	if ok {
		v,_ := value.(User)
		user = v
	} else {
		DB.First(&user)
		profileRef := &user.Profile
		DB.Model(&user).Related(profileRef, "Profile").First(profileRef)
		addressesRef := &profileRef.Addresses
		DB.Model(profileRef).Related(addressesRef).Find(addressesRef)
	}
	return user
}

func (repo Repository) GetAll () []User {
	var users = []User{}
	DB.Model(&User{}).Find(&users)
	return users
}

func (repo Repository) Save(changedUserRef *User)  {
	key := getRepositoryId(*changedUserRef)

	DB.Save(changedUserRef)
	profileRef := &changedUserRef.Profile
	DB.Save(profileRef)
	addresses := profileRef.Addresses
	for _, address := range addresses {
		DB.Save(&address)
	}
	cacheUser := User{}
	cacheUser = *changedUserRef
	cache.Add(key, cacheUser)
}

func (repo Repository) FindAll(db gorm.DB) []User {
	var users = make([]User, 30)
	db.Find(&users)

	for _, user := range users {
		user = repo.Get(user.ID)
	}

	return users
}

func (repo Repository) FindOne(db gorm.DB) User {
	var user = User{}
	db.First(&user)
	user = repo.Get(user.ID)
	return user
}
