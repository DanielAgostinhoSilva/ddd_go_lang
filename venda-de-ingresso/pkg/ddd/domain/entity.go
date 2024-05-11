package domain

type Entity interface {
	Equals(value interface{}) bool
	ToJson() (string, error)
}
