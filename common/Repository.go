package common
import "github.com/jinzhu/gorm"


//Empty interface expects input/output have the type of a struct of aggregate root, not the pointer to struct
type Repository interface {
	Create(T interface{})
	Get(id uint64) interface{}
	GetAll() []interface{}
	FindAll(db gorm.DB) []interface{}
	FindOne(db gorm.DB) interface{}
	SaveOrUpdate(T interface{})
}

