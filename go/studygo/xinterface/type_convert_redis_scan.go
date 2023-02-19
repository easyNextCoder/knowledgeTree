package xinterface

import (
	"fmt"
	"reflect"
	"strconv"
)

type Error string

func ensureLen(d reflect.Value, n int) {
	if n > d.Cap() {
		d.Set(reflect.MakeSlice(d.Type(), n, n))
	} else {
		d.SetLen(n)
	}
}

func cannotConvert(d reflect.Value, s interface{}) error {
	var sname string
	switch s.(type) {
	case string:
		sname = "Redis simple string"
	case Error:
		sname = "Redis error"
	case int64:
		sname = "Redis integer"
	case []byte:
		sname = "Redis bulk string"
	case []interface{}:
		sname = "Redis array"
	default:
		sname = reflect.TypeOf(s).String()
	}
	return fmt.Errorf("cannot convert from %s to %s", sname, d.Type())
}

func convertAssignBulkString(d reflect.Value, s []byte) (err error) {
	switch d.Type().Kind() {
	case reflect.Float32, reflect.Float64:
		var x float64
		x, err = strconv.ParseFloat(string(s), d.Type().Bits())
		d.SetFloat(x)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var x int64
		x, err = strconv.ParseInt(string(s), 10, d.Type().Bits())
		d.SetInt(x)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var x uint64
		x, err = strconv.ParseUint(string(s), 10, d.Type().Bits())
		d.SetUint(x)
	case reflect.Bool:
		var x bool
		x, err = strconv.ParseBool(string(s))
		d.SetBool(x)
	case reflect.String:
		d.SetString(string(s))
	case reflect.Slice:
		if d.Type().Elem().Kind() != reflect.Uint8 {
			err = cannotConvert(d, s)
		} else {
			d.SetBytes(s)
		}
	default:
		err = cannotConvert(d, s)
	}
	return
}

func convertAssignInt(d reflect.Value, s int64) (err error) {
	switch d.Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		d.SetInt(s)
		if d.Int() != s {
			err = strconv.ErrRange
			d.SetInt(0)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if s < 0 {
			err = strconv.ErrRange
		} else {
			x := uint64(s)
			d.SetUint(x)
			if d.Uint() != x {
				err = strconv.ErrRange
				d.SetUint(0)
			}
		}
	case reflect.Bool:
		d.SetBool(s != 0)
	default:
		err = cannotConvert(d, s)
	}
	return
}

func convertAssignValue(d reflect.Value, s interface{}) (err error) {
	if d.Kind() != reflect.Ptr {
		if d.CanAddr() {
			d2 := d.Addr()
			if d2.CanInterface() {
				fmt.Println("convertAssignValue0 ")
				//if scanner, ok := d2.Interface().(Scanner); ok {
				//	return scanner.RedisScan(s)
				//}
			}
		}
	} else if d.CanInterface() {
		// Already a reflect.Ptr
		if d.IsNil() {
			d.Set(reflect.New(d.Type().Elem()))
		}
		fmt.Println("convertAssignValue ")
		//if scanner, ok := d.Interface().(Scanner); ok {
		//	return scanner.RedisScan(s)
		//}
	}

	switch s := s.(type) {
	case []byte:
		err = convertAssignBulkString(d, s)
	case int64:
		err = convertAssignInt(d, s)
	default:
		err = cannotConvert(d, s)
	}
	return err
}

func convertAssignArray(d reflect.Value, s []interface{}) error {
	if d.Type().Kind() != reflect.Slice {
		return cannotConvert(d, s)
	}
	ensureLen(d, len(s))
	for i := 0; i < len(s); i++ {
		if err := convertAssignValue(d.Index(i), s[i]); err != nil {
			return err
		}
	}
	return nil
}

//此代码片段为redisgo中将reply转换成对应数据结构的函数
//源代码：在game-hand中
//_, err = redis.Scan(reply, &pos, &keyVals)
//仿造代码：yorm/yservices/read.go中 使用反射的方式来获取数据
func ConvertAssign(d interface{}, s interface{}) (err error) {
	// Handle the most common destination types using type switches and
	// fall back to reflection for all other types.
	switch s := s.(type) {
	case nil:
		// ignore
	case []byte:
		switch d := d.(type) {
		case *string:
			*d = string(s)
		case *int:
			*d, err = strconv.Atoi(string(s))
		case *bool:
			*d, err = strconv.ParseBool(string(s))
		case *[]byte:
			*d = s
		case *interface{}:
			*d = s
		case nil:
			// skip value
		default:
			if d := reflect.ValueOf(d); d.Type().Kind() == reflect.Ptr {
				err = convertAssignBulkString(d.Elem(), s)
			}
		}
	case int64:
		switch d := d.(type) {
		case *int:
			x := int(s)
			if int64(x) != s {
				err = strconv.ErrRange
				x = 0
			}
			*d = x
		case *bool:
			*d = s != 0
		case *interface{}:
			*d = s
		case nil:
			// skip value
		default:
			if d := reflect.ValueOf(d); d.Type().Kind() == reflect.Ptr {
				err = convertAssignInt(d.Elem(), s)
			}
		}
	case string:
		switch d := d.(type) {
		case *string:
			*d = s
		case *interface{}:
			*d = s
		case nil:
			// skip value
		default:
			fmt.Println("can't convert")
		}
	case []interface{}:
		switch d := d.(type) {
		case *[]interface{}:
			*d = s
		case *interface{}:
			*d = s
		case nil:
			// skip value
		default:
			if d := reflect.ValueOf(d); d.Type().Kind() == reflect.Ptr {
				err = convertAssignArray(d.Elem(), s)
			}
		}
	default:
		err = cannotConvert(reflect.ValueOf(d), s)
	}
	return
}
