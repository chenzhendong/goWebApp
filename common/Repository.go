package common

type Repository interface {
	Get(id uint64) interface{}
	GetAll() []interface{}
	FindAll(T map[string]interface{}) []interface{}
	FindOne(T map[string]interface{}) interface{}
	SaveOrUpdate(interface{})
}

