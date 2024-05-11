package domain

type Entity interface {
	ToJson() (string, error)
	Equals(id string) bool
}
