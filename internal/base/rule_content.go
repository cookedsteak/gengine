package base

import (
	"errors"
	"reflect"

	"github.com/cookedsteak/gengine/context"
)

type RuleContent struct {
	Statements *Statements
}

func (t *RuleContent) Execute(dc *context.DataContext, Vars map[string]reflect.Value) (reflect.Value, error, bool) {
	return t.Statements.Evaluate(dc, Vars)
}

func (t *RuleContent) AcceptStatements(stmts *Statements) error {
	if t.Statements == nil {
		t.Statements = stmts
		return nil
	}
	return errors.New("RuleContent's statements set twice.")
}
