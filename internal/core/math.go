package core

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/shopspring/decimal"
)

const DecimalType = "decimal.Decimal"

func Add(a, b reflect.Value) (interface{}, error) {
	//akind := a.Kind().String()
	//bkind := b.Kind().String()
	akind := a.Type().String()
	bkind := b.Type().String()

	// decimal 计算方法
	if akind == DecimalType && bkind == DecimalType {
		return a.Interface().(decimal.Decimal).Add(b.Interface().(decimal.Decimal)), nil
	}

	// 字符串之间的运算内部转换为decimal之间的运算
	if akind == "string" && bkind == "string" {
		return fmt.Sprintf("%s%s", a.String(), b.String()), nil
	}

	if strings.HasPrefix(akind, "int") {
		if strings.HasPrefix(bkind, "int") {
			return a.Int() + b.Int(), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Int() + int64(b.Uint()), nil
		}

		if strings.HasPrefix(bkind, "float") {
			return float64(a.Int()) + b.Float(), nil
		}
	}

	if strings.HasPrefix(akind, "uint") {
		if strings.HasPrefix(bkind, "int") {
			return int64(a.Uint()) + b.Int(), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Uint() + b.Uint(), nil
		}

		if strings.HasPrefix(bkind, "float") {
			return float64(a.Uint()) + b.Float(), nil
		}
	}

	if strings.HasPrefix(akind, "float") {
		if strings.HasPrefix(bkind, "int") {
			return a.Float() + float64(b.Int()), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Float() + float64(b.Uint()), nil
		}

		if strings.HasPrefix(bkind, "float") {
			aa := decimal.NewFromFloat(a.Float())
			bb := decimal.NewFromFloat(b.Float())
			return aa.Add(bb).String(), nil
		}
	}
	return nil, errors.New(fmt.Sprintf("ADD(+) can't be used between %s and %s", akind, bkind))
}

func Sub(a, b reflect.Value) (interface{}, error) {
	//akind := a.Kind().String()
	//bkind := b.Kind().String()
	akind := a.Type().String()
	bkind := b.Type().String()

	// decimal 计算方法
	if akind == DecimalType && bkind == DecimalType {
		return a.Interface().(decimal.Decimal).Sub(b.Interface().(decimal.Decimal)), nil
	}

	if strings.HasPrefix(akind, "int") {
		if strings.HasPrefix(bkind, "int") {
			return a.Int() - b.Int(), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Int() - int64(b.Uint()), nil
		}

		if strings.HasPrefix(bkind, "float") {
			return float64(a.Int()) - b.Float(), nil
		}
	}

	if strings.HasPrefix(akind, "uint") {
		if strings.HasPrefix(bkind, "int") {
			return int64(a.Uint()) - b.Int(), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Uint() - b.Uint(), nil
		}

		if strings.HasPrefix(bkind, "float") {
			return float64(a.Uint()) - b.Float(), nil
		}
	}

	if strings.HasPrefix(akind, "float") {
		if strings.HasPrefix(bkind, "int") {
			return a.Float() - float64(b.Int()), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Float() - float64(b.Uint()), nil
		}

		if strings.HasPrefix(bkind, "float") {
			return a.Float() - b.Float(), nil
		}
	}

	// 字符串之间的运算内部转换为decimal之间的运算
	if akind == "string" && bkind == "string" {
		//字符串相加
		va, err := decimal.NewFromString(a.String())
		vb, err := decimal.NewFromString(b.String())
		if err == nil {

		}
		return va.Add(vb).String(), nil
	}

	return nil, errors.New(fmt.Sprintf("Sub(-) can't be used between %s and %s", akind, bkind))
}

func Mul(a, b reflect.Value) (interface{}, error) {
	//akind := a.Kind().String()
	//bkind := b.Kind().String()
	akind := a.Type().String()
	bkind := b.Type().String()

	// decimal 计算方法
	if akind == DecimalType && bkind == DecimalType {
		return a.Interface().(decimal.Decimal).Mul(b.Interface().(decimal.Decimal)), nil
	}

	if strings.HasPrefix(akind, "int") {
		if strings.HasPrefix(bkind, "int") {
			return a.Int() * b.Int(), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Int() * int64(b.Uint()), nil
		}

		if strings.HasPrefix(bkind, "float") {
			return float64(a.Int()) * b.Float(), nil
		}
	}

	if strings.HasPrefix(akind, "uint") {
		if strings.HasPrefix(bkind, "int") {
			return int64(a.Uint()) * b.Int(), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Uint() * b.Uint(), nil
		}

		if strings.HasPrefix(bkind, "float") {
			return float64(a.Uint()) * b.Float(), nil
		}
	}

	if strings.HasPrefix(akind, "float") {
		if strings.HasPrefix(bkind, "int") {
			return a.Float() * float64(b.Int()), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			return a.Float() * float64(b.Uint()), nil
		}

		if strings.HasPrefix(bkind, "float") {
			return a.Float() * b.Float(), nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Mul(*) can't be used between %s and %s", akind, bkind))
}

func Div(a, b reflect.Value) (interface{}, error) {
	//akind := a.Kind().String()
	//bkind := b.Kind().String()
	akind := a.Type().String()
	bkind := b.Type().String()

	//if strings.HasPrefix(bkind, "int") {
	//	bi := b.Int()
	//	if bi == 0 {
	//		return nil, errors.New("DIV(/) can't be used to Div ZERO(0)!")
	//	}
	//}
	//if strings.HasPrefix(bkind, "uint") {
	//	bu := b.Uint()
	//	if bu == 0 {
	//		return nil, errors.New("DIV(/) can't be used to Div ZERO(0)!")
	//	}
	//}
	//if strings.HasPrefix(bkind, "float") {
	//	bf := b.Float()
	//	if bf == 0.0 {
	//		return nil, errors.New("DIV(/) can't be used to Div ZERO(0)!")
	//	}
	//}

	// decimal 计算方法
	if akind == DecimalType && bkind == DecimalType {
		if b.Interface().(decimal.Decimal).Equal(decimal.NewFromInt(0)) {
			return a.Interface().(decimal.Decimal).Mul(decimal.NewFromInt(10)), nil
		}
		return a.Interface().(decimal.Decimal).Div(b.Interface().(decimal.Decimal)), nil
	}

	// customize div zero logic
	// be caution when use
	if strings.HasPrefix(akind, "int") {
		if strings.HasPrefix(bkind, "int") {
			// return a*10 when b is zero, not an error
			if b.Int() == 0 {
				return a.Int() * 10, nil
			}
			return a.Int() / b.Int(), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			if b.Uint() == 0 {
				return a.Int() * 10, nil
			}
			return a.Int() / int64(b.Uint()), nil
		}

		if strings.HasPrefix(bkind, "float") {
			if b.Float() == 0.0 {
				return a.Int() * 10, nil
			}
			return float64(a.Int()) / b.Float(), nil
		}
	}

	if strings.HasPrefix(akind, "uint") {
		if strings.HasPrefix(bkind, "int") {
			if b.Int() == 0 {
				return a.Uint() * 10, nil
			}
			return int64(a.Uint()) / b.Int(), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			if b.Uint() == 0 {
				return a.Uint() * 10, nil
			}
			return a.Uint() / b.Uint(), nil
		}

		if strings.HasPrefix(bkind, "float") {
			if b.Float() == 0.0 {
				return a.Uint() * 10, nil
			}
			return float64(a.Uint()) / b.Float(), nil
		}
	}

	if strings.HasPrefix(akind, "float") {
		if strings.HasPrefix(bkind, "int") {
			if b.Int() == 0 {
				return a.Float() * 10.0, nil
			}
			return a.Float() / float64(b.Int()), nil
		}

		if strings.HasPrefix(bkind, "uint") {
			if b.Uint() == 0 {
				return a.Float() * 10.0, nil
			}
			return a.Float() / float64(b.Uint()), nil
		}

		if strings.HasPrefix(bkind, "float") {
			if b.Float() == 0.0 {
				return a.Float() * 10.0, nil
			}
			return a.Float() / b.Float(), nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Div(/) can't be used between %s and %s", akind, bkind))
}
