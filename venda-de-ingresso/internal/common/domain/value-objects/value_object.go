package value_objects

type ValueObject interface {
	Equals(value interface{}) bool
}
