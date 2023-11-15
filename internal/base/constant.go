package base

import (
	"reflect"
	"regexp"

	"github.com/shopspring/decimal"

	"github.com/cookedsteak/gengine/context"
)

type Constant struct {
	ConstantValue reflect.Value
}

func (cons *Constant) AcceptString(str string) error {
	cons.ConstantValue = reflect.ValueOf(str)
	return nil
}

func (cons *Constant) Evaluate(dc *context.DataContext, Vars map[string]reflect.Value) (reflect.Value, error) {
	//fmt.Println(cons.ConstantValue.Type().String())
	if cons.ConstantValue.Type().String() == "string" {
		floatReg := `^[+\-]?\d+\.*\d*[eE]?[+\-]?\d*$`
		if ok, _ := regexp.MatchString(floatReg, cons.ConstantValue.String()); ok {
			res, err := decimal.NewFromString(cons.ConstantValue.String())
			if err != nil {
				return cons.ConstantValue, err
			}
			return reflect.ValueOf(res), nil
		}
	}
	return cons.ConstantValue, nil
}

func (cons *Constant) AcceptInteger(i64 int64) error {
	cons.ConstantValue = reflect.ValueOf(i64)
	return nil
}

// receive rule's name
func (cons *Constant) AcceptName(name string) error {
	cons.ConstantValue = reflect.ValueOf(name)
	return nil
}

func (cons *Constant) AcceptId(id int64) error {
	cons.ConstantValue = reflect.ValueOf(id)
	return nil
}

// receive rule's description
func (cons *Constant) AcceptDesc(desc string) error {
	cons.ConstantValue = reflect.ValueOf(desc)
	return nil
}

func (cons *Constant) AcceptSalience(sal int64) error {
	cons.ConstantValue = reflect.ValueOf(sal)
	return nil
}
