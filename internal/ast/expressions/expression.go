package expressions

import "github.com/hellodhlyn/akane/internal/objects"

type Expression interface {
	Eval() objects.Object
}
