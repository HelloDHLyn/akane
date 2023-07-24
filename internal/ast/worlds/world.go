package worlds

import (
	"github.com/hellodhlyn/akane/internal/ast/expressions"
	"github.com/hellodhlyn/akane/internal/objects"
)

type Args []objects.Object

type World struct {
	Expressions []expressions.Expression
}

func NewWorld(exprs []expressions.Expression) *World {
	return &World{Expressions: exprs}
}

func (w *World) Eval(_ *Args) objects.Object {
	// TODO - supports real `return` statement
	var returnObj objects.Object
	for idx, expr := range w.Expressions {
		if idx == len(w.Expressions)-1 {
			returnObj = expr.Eval()
		} else {
			expr.Eval()
		}
	}
	return returnObj
}
