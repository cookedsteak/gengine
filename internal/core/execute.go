package core

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/shopspring/decimal"
)

func InvokeFunction(obj reflect.Value, methodName string, parameters []reflect.Value) (reflect.Value, error) {
	objVal := obj

	fun := objVal.MethodByName(methodName)
	if !fun.IsValid() {
		return reflect.ValueOf(nil), errors.New(fmt.Sprintf("NOT FOUND Function: %s", methodName))
	}

	//fmt.Println(parameters[0].String())
	//fmt.Println(parameters[1].String())
	params := ParamsTypeChange(fun, parameters)
	//for _, v := range params {
	//	fmt.Println("调用的参数类型为：", v.Type().String(), reflect.ValueOf(v))
	//}
	rs := fun.Call(params)
	raw, e := GetRawTypeValue(rs)
	if e != nil {
		return reflect.ValueOf(nil), e
	}
	return raw, nil
}

/*
*  todo 多返回值修改如下函数
if want to support multi return ,change this method implements
*/
func GetRawTypeValue(rs []reflect.Value) (reflect.Value, error) {
	if len(rs) == 0 {
		return reflect.ValueOf(nil), nil
	} else {
		return rs[0], nil
	}
}

func GetStructAttributeValue(obj reflect.Value, fieldName string) (reflect.Value, error) {
	stru := obj
	var attrVal reflect.Value
	if stru.Kind() == reflect.Ptr {
		attrVal = stru.Elem().FieldByName(fieldName)
	} else {
		attrVal = stru.FieldByName(fieldName)
	}
	return attrVal, nil
}

// SetAttributeValue
// set field value
// value 就是右表达式的值，可能是decimal类型
func SetAttributeValue(obj reflect.Value, fieldName string, value reflect.Value) error {
	field := reflect.ValueOf(nil)
	objType := obj.Type()
	objVal := obj
	if objType.Kind() == reflect.Ptr {
		//it points to struct
		if objType.Elem().Kind() == reflect.Struct {
			field = objVal.Elem().FieldByName(fieldName)
		}
	} else {
		//not a pointer.
		if objType.Kind() == reflect.Struct {
			field = objVal.FieldByName(fieldName)
		}
	}

	if field == reflect.ValueOf(nil) {
		return errors.New(fmt.Sprintf("struct has no this field: %s", fieldName))
	}

	if field.CanSet() {
		typeName := value.Type().String()
		switch field.Type().Kind() {
		case reflect.String:
			if strings.HasPrefix(typeName, "decimal") {
				field.SetString(value.Interface().(decimal.Decimal).String())
				return nil
			}
			field.SetString(value.String())
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if strings.HasPrefix(typeName, "decimal") {
				field.SetInt(value.Interface().(decimal.Decimal).IntPart())
				return nil
			}
			if strings.HasPrefix(typeName, "uint") {
				field.SetInt(int64(value.Uint()))
				return nil
			}
			if strings.HasPrefix(typeName, "float") {
				field.SetInt(int64(value.Float()))
				return nil
			}
			field.SetInt(value.Int())
			break
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if strings.HasPrefix(typeName, "decimal") {
				field.SetUint(uint64(value.Interface().(decimal.Decimal).IntPart()))
				return nil
			}
			if strings.HasPrefix(typeName, "int") && value.Int() >= 0 {
				field.SetUint(uint64(value.Int()))
				return nil
			}
			if strings.HasPrefix(typeName, "float") && value.Float() >= 0 {
				field.SetUint(uint64(value.Float()))
				return nil
			}
			field.SetUint(value.Uint())
			break
		case reflect.Float32, reflect.Float64:
			if strings.HasPrefix(typeName, "decimal") {
				field.SetFloat(value.Interface().(decimal.Decimal).InexactFloat64())
				return nil
			}
			if strings.HasPrefix(typeName, "int") {
				field.SetFloat(float64(value.Int()))
				return nil
			}
			if strings.HasPrefix(typeName, "uint") {
				field.SetFloat(float64(value.Uint()))
				return nil
			}
			field.SetFloat(value.Float())
			break
		case reflect.Bool:
			field.SetBool(value.Bool())
			break
		case reflect.Slice:
			field.Set(value)
			break
		case reflect.Map:
			field.Set(value)
			break
		case reflect.Array:
			field.Set(value)
			break
		case reflect.Struct:
			field.Set(value)
			break
		case reflect.Interface:
			field.Set(value)
			break
		case reflect.Chan:
			field.Set(value)
			break
		case reflect.Complex64:
			field.SetComplex(value.Complex())
			break
		case reflect.Complex128:
			field.SetComplex(value.Complex())
			break
		case reflect.Func:
			field.Set(value)
			break
		default:
			return errors.New(fmt.Sprintf("Not support type:%s", field.Type().Kind().String()))
		}
	} else {
		return errors.New(fmt.Sprintf("%s:must Be Assignable, it should be or be in addressable value!", field.Type().Kind().String()))
	}
	return nil
}

// set single value
func SetSingleValue(obj reflect.Value, fieldName string, value reflect.Value) error {
	if obj.Kind() == reflect.Ptr {
		if value.Kind() == reflect.Ptr {
			//both ptr
			value = value.Elem()
		}

		objKind := obj.Elem().Kind()
		valueKind := value.Kind()
		if objKind == valueKind {
			obj.Elem().Set(value)
			return nil
		} else {
			valueKindStr := valueKind.String()
			switch objKind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if strings.HasPrefix(valueKindStr, "int") {
					obj.Elem().SetInt(value.Int())
					return nil
				}
				if strings.HasPrefix(valueKindStr, "float") {
					obj.Elem().SetInt(int64(value.Float()))
					return nil
				}
				if strings.HasPrefix(valueKindStr, "uint") {
					obj.Elem().SetInt(int64(value.Uint()))
					return nil
				}
				break
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				if strings.HasPrefix(valueKindStr, "int") && value.Int() >= 0 {
					obj.Elem().SetUint(uint64(value.Int()))
					return nil
				}
				if strings.HasPrefix(valueKindStr, "float") && value.Float() >= 0 {
					obj.Elem().SetUint(uint64(value.Float()))
					return nil
				}
				if strings.HasPrefix(valueKindStr, "uint") {
					obj.Elem().SetUint(value.Uint())
					return nil
				}
				break
			case reflect.Float32, reflect.Float64:
				if strings.HasPrefix(valueKindStr, "int") {
					obj.Elem().SetFloat(float64(value.Int()))
					return nil
				}
				if strings.HasPrefix(valueKindStr, "float") {
					obj.Elem().SetFloat(value.Float())
					return nil
				}
				if strings.HasPrefix(valueKindStr, "uint") {
					obj.Elem().SetFloat(float64(value.Uint()))
					return nil
				}
				break
			}
			return errors.New(fmt.Sprintf("\"%s\" value type \"%s\" is different from \"%s\" ", fieldName, obj.Elem().Kind().String(), value.Kind().String()))
		}
	} else {
		return errors.New(fmt.Sprintf("\"%s\" value is unassignable!", fieldName))
	}
}

const (
	_int     = 1
	_uint    = 2
	_float   = 3
	_decimal = 4
)

// ParamsTypeChange
// 自定义算子函数入参的数据类型转换
// @param f是定义的注入函数反射值，可以直接调用
// @param params 是入参
func ParamsTypeChange(f reflect.Value, params []reflect.Value) []reflect.Value {
	// 函数的类型
	tf := f.Type()
	if tf.Kind() == reflect.Ptr {
		tf = tf.Elem()
	}
	pLen := tf.NumIn() // 函数期望参数的数量
	for i := 0; i < pLen; i++ {
		// 这里选择预定义的参数的类型
		switch tf.In(i).Kind() {
		// 参数可能是"123"数字字符，但是底层解析为了decimal，所以需要再转换一下
		case reflect.Slice: // 所期望的是slice
			var err error
			// 这里slice会有两种情况
			// 1、一种是期望参数本身就应该是slice
			// 2、另一种可能，参数是任意数量参数 (...xx)
			if params[i].Kind() != reflect.Slice && params[i].Kind() != reflect.Array { // 任意数量参数
				// 从当前参数开始，后面的参数都要做同类型的转换
				for pk := i; pk < len(params); pk++ {
					params[pk], err = GetWantedValue(params[pk], tf.In(i).Elem()) // In获取参数类型，Elem获取切片参数中的元素类型
					if err != nil {
						fmt.Println(err.Error())
						continue
					}
				}
			} else {
				// 正常slice类型
				params[i], err = GetWantedValue(params[i], tf.In(i).Elem())
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
			}
			break
		case reflect.String:
			if params[i].Type().String() == "decimal.Decimal" {
				params[i] = reflect.ValueOf(params[i].Interface().(decimal.Decimal).String())
			}
			break
		case reflect.Int:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(int(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(int(params[i].Uint()))
			} else if tag == _decimal {
				params[i] = reflect.ValueOf(int(params[i].Interface().(decimal.Decimal).IntPart()))
			} else {
				params[i] = reflect.ValueOf(int(params[i].Float()))
			}
			break
		case reflect.Int8:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(int8(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(int8(params[i].Uint()))
			} else if tag == _decimal {
				params[i] = reflect.ValueOf(int8(params[i].Interface().(decimal.Decimal).IntPart()))
			} else {
				params[i] = reflect.ValueOf(int8(params[i].Float()))
			}
			break
		case reflect.Int16:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(int16(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(int16(params[i].Uint()))
			} else if tag == _decimal {
				params[i] = reflect.ValueOf(int16(params[i].Interface().(decimal.Decimal).IntPart()))
			} else {
				params[i] = reflect.ValueOf(int16(params[i].Float()))
			}
			break
		case reflect.Int32:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(int32(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(int32(params[i].Uint()))
			} else if tag == _decimal {
				params[i] = reflect.ValueOf(int32(params[i].Interface().(decimal.Decimal).IntPart()))
			} else {
				params[i] = reflect.ValueOf(int32(params[i].Float()))
			}
			break
		case reflect.Int64:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(params[i].Int())
			} else if tag == _uint {
				params[i] = reflect.ValueOf(int64(params[i].Uint()))
			} else if tag == _decimal {
				params[i] = reflect.ValueOf(params[i].Interface().(decimal.Decimal).IntPart())
			} else {
				params[i] = reflect.ValueOf(int64(params[i].Float()))
			}
			break
		case reflect.Uint:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(uint(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(uint(params[i].Uint()))
			} else if tag == _decimal {
				params[i] = reflect.ValueOf(uint(params[i].Interface().(decimal.Decimal).IntPart()))
			} else {
				params[i] = reflect.ValueOf(uint(params[i].Float()))
			}
			break
		case reflect.Uint8:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(uint8(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(uint8(params[i].Uint()))
			} else if tag == _decimal {
				params[i] = reflect.ValueOf(uint8(params[i].Interface().(decimal.Decimal).IntPart()))
			} else {
				params[i] = reflect.ValueOf(uint8(params[i].Float()))
			}
			break
		case reflect.Uint16:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(uint16(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(uint16(params[i].Uint()))
			} else if tag == _decimal {
				params[i] = reflect.ValueOf(uint16(params[i].Interface().(decimal.Decimal).IntPart()))
			} else {
				params[i] = reflect.ValueOf(uint16(params[i].Float()))
			}
			break
		case reflect.Uint32:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(uint32(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(uint32(params[i].Uint()))
			} else if tag == _decimal {
				params[i] = reflect.ValueOf(uint32(params[i].Interface().(decimal.Decimal).IntPart()))
			} else {
				params[i] = reflect.ValueOf(uint32(params[i].Float()))
			}
			break
		case reflect.Uint64:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(uint64(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(params[i].Uint())
			} else if tag == _decimal {
				params[i] = reflect.ValueOf(uint64(params[i].Interface().(decimal.Decimal).IntPart()))
			} else {
				params[i] = reflect.ValueOf(uint64(params[i].Float()))
			}
			break
		case reflect.Float32:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(float32(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(float32(params[i].Uint()))
			} else if tag == _decimal {
				params[i] = reflect.ValueOf(float32(params[i].Interface().(decimal.Decimal).InexactFloat64()))
			} else {
				params[i] = reflect.ValueOf(float32(params[i].Float()))
			}
			break
		case reflect.Float64:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(float64(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(float64(params[i].Uint()))
			} else if tag == _decimal {
				params[i] = reflect.ValueOf(params[i].Interface().(decimal.Decimal).InexactFloat64())
			} else {
				params[i] = reflect.ValueOf(params[i].Float())
			}
			break
		case reflect.Ptr:
			break
		case reflect.Interface:
			fmt.Println(tf.In(i).Kind().String())
			fmt.Println("func params is interface")
			dd := params[i].Type().String()
			fmt.Println(dd)
			if !reflect.ValueOf(params[i]).IsValid() {
				params[i] = reflect.New(tf.In(i))
			}
		default:
			continue
		}
	}
	return params
}

// 获取入参的参数类型
func getNumType(param reflect.Value) int {
	ts := param.Type().String()
	if strings.HasPrefix(ts, "int") {
		return _int
	}
	if strings.HasPrefix(ts, "uint") {
		return _uint
	}
	if strings.HasPrefix(ts, "float") {
		return _float
	}
	if strings.HasPrefix(ts, "decimal") {
		return _decimal
	}

	panic(fmt.Sprintf("it is not number type, type is %s !", ts))
}

// newValue: 真实入参类型
// toKind: 期望的入参类型
func GetWantedValue(newValue reflect.Value, toKind reflect.Type) (reflect.Value, error) {
	fmt.Println(fmt.Sprintf("入参类型：%s", newValue.Kind().String()))
	fmt.Println(fmt.Sprintf("期望类型：%s", toKind.Kind().String()))
	if newValue.Kind() == toKind.Kind() {
		return newValue, nil
	}
	// decimal的处理逻辑
	if newValue.Type().String() == "decimal.Decimal" {
		return ConvertDecimalToType(newValue, toKind)
	}

	switch toKind.Kind() {
	case reflect.Int:
		return reflect.ValueOf(int(newValue.Int())), nil
	case reflect.Int8:
		return reflect.ValueOf(int8(newValue.Int())), nil
	case reflect.Int16:
		return reflect.ValueOf(int16(newValue.Int())), nil
	case reflect.Int32:
		return reflect.ValueOf(int32(newValue.Int())), nil
	case reflect.Int64:
		return newValue, nil

	case reflect.Uint:
		return reflect.ValueOf(uint(newValue.Uint())), nil
	case reflect.Uint8:
		return reflect.ValueOf(uint8(newValue.Uint())), nil
	case reflect.Uint16:
		return reflect.ValueOf(uint16(newValue.Uint())), nil
	case reflect.Uint32:
		return reflect.ValueOf(uint32(newValue.Uint())), nil
	case reflect.Uint64:
		return newValue, nil

	case reflect.Float32:
		return reflect.ValueOf(float32(newValue.Float())), nil
	case reflect.Float64:
		return newValue, nil
	}

	return newValue, nil
}

// newValue 入参类型
// toKind 期望类型
func ConvertDecimalToType(newValue reflect.Value, toKind reflect.Type) (reflect.Value, error) {
	tk := toKind.String()

	if tk == "string" {
		midValue := newValue.Interface().(decimal.Decimal).String()
		return reflect.ValueOf(midValue), nil
	}

	if strings.HasPrefix(tk, "int") || strings.HasPrefix(tk, "uint") {
		midValue := newValue.Interface().(decimal.Decimal).IntPart()
		switch toKind.Kind() {
		case reflect.Int:
			return reflect.ValueOf(int(midValue)), nil
		case reflect.Int8:
			return reflect.ValueOf(int8(midValue)), nil
		case reflect.Int16:
			return reflect.ValueOf(int16(midValue)), nil
		case reflect.Int32:
			return reflect.ValueOf(int32(midValue)), nil
		case reflect.Int64:
			return reflect.ValueOf(midValue), nil

		case reflect.Uint:
			return reflect.ValueOf(uint(midValue)), nil
		case reflect.Uint8:
			return reflect.ValueOf(uint8(midValue)), nil
		case reflect.Uint16:
			return reflect.ValueOf(uint16(midValue)), nil
		case reflect.Uint32:
			return reflect.ValueOf(uint32(midValue)), nil
		case reflect.Uint64:
			return reflect.ValueOf(uint64(midValue)), nil
		}
	}

	if strings.HasPrefix(tk, "float") {
		midValue := newValue.Interface().(decimal.Decimal).InexactFloat64()
		switch toKind.Kind() {
		case reflect.Float32:
			return reflect.ValueOf(float32(midValue)), nil
		case reflect.Float64:
			return reflect.ValueOf(midValue), nil
		}
	}

	return newValue, nil
}

func GetFunctionName(i reflect.Value) string {
	return runtime.FuncForPC(i.Pointer()).Name()
}
