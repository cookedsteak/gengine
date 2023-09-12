package core

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"

	"github.com/shopspring/decimal"
)

func Add(a, b reflect.Value) (interface{}, error) {
	akind := a.Type().String()
	bkind := b.Type().String()
	// 字符串之间的运算内部转换为decimal之间的运算
	if akind == "string" && bkind == "string" {
		return fmt.Sprintf("%s%s", a.String(), b.String()), nil
	}

	aval, at, err := MakeDecimalByType(a)
	bval, bt, err := MakeDecimalByType(b)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("ADD(+) can't be used between %s and %s", at, bt))
	}
	return aval.Add(bval), nil
}

func Sub(a, b reflect.Value) (interface{}, error) {
	aval, at, err := MakeDecimalByType(a)
	bval, bt, err := MakeDecimalByType(b)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Sub(-) can't be used between %s and %s", at, bt))
	}

	return aval.Sub(bval), nil
}

func Mul(a, b reflect.Value) (interface{}, error) {
	aval, at, err := MakeDecimalByType(a)
	bval, bt, err := MakeDecimalByType(b)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Mul(*) can't be used between %s and %s", at, bt))
	}

	return aval.Mul(bval), nil
}

func Div(a, b reflect.Value) (interface{}, error) {
	aval, at, err := MakeDecimalByType(a)
	bval, bt, err := MakeDecimalByType(b)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Div(/) can't be used between %s and %s", at, bt))
	}

	if bval.Equal(decimal.NewFromInt(0)) {
		return aval.Mul(decimal.NewFromInt(10)), nil
	}
	return aval.Div(bval), nil

}

func MakeDecimalByType(val reflect.Value) (decimal.Decimal, string, error) {
	//t := val.Kind().String()
	t := val.Type().String()
	var ret decimal.Decimal
	switch t {
	case "int", "int8", "int16", "int32", "int64":
		ret = decimal.NewFromInt(val.Int())
		break
	case "uint", "uint8", "uint16", "uint32", "uint64":
		ret = decimal.NewFromInt(int64(val.Uint()))
		break
	case "float32", "float64", "float":
		ret = decimal.NewFromFloat(val.Float())
	case "decimal.Decimal":
		ret = val.Interface().(decimal.Decimal)
		break
	case "string":
		//floatReg := `[+\-]?\d+\.*\d*[eE]?[+\-]?\d*`
		var err error
		ret, err = decimal.NewFromString(val.String())
		if err != nil {
			return decimal.Decimal{}, t, errors.New("value can not convert into decimal")
		}
		break
	default:
		return decimal.Decimal{}, t, errors.New("value can not convert into decimal")
	}
	return ret, t, nil
}
