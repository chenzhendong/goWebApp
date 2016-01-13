package common

type Repository interface {
	Get(T interface{})
	GetAll() []interface{}
	FindAll(T map[string]interface{}) []interface{}
	FindOne(T map[string]interface{}) interface{}
	SaveOrUpdate(interface{})
}

