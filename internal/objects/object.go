package objects

type ObjectType int

const (
	WorldObject ObjectType = iota
	IntegerObject
)

type Object interface {
	Type() ObjectType
	Bytes() []byte
}
