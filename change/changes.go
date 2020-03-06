package change

type Changes []Change

type Change interface {
	Name() string
	Value() interface{}
}
