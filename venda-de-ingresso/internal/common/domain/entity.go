package domain

type Entity interface {
	ToJson() (string, error)
	Equals(id interface{}) bool
	Id() interface{}
}
