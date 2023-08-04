package godata

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Variant struct {
	Value any
}

func (v Variant) AsValue() any {
	return v.Value
}

func (v Variant) AsString() string {
	value := ""
	switch val := v.Value.(type) {
	case nil:
		value = ""
	case time.Time:
		value = val.String()
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		value = fmt.Sprintf("%v", val)
	case float32, float64:
		value = fmt.Sprintf("%f", val)
	case string:
		value = val
	default:
		t := reflect.TypeOf(v.Value)
		fmt.Printf("unable to convert data type to string. Type: %v", t)
		value = ""
	}
	value = strings.Replace(value, "\r", "\n", -1)
	return value
}

func (v Variant) AsInt() int {
	switch val := v.Value.(type) {
	case nil:
		return int(0)
	case int:
		return v.Value.(int)
	case int8:
		return int(v.Value.(int8))
	case int16:
		return int(v.Value.(int16))
	case int32:
		return int(v.Value.(int32))
	case int64:
		return int(v.Value.(int64))
	case string:
		intValue, err := strconv.Atoi(val)
		if err != nil {
			t := reflect.TypeOf(val)
			fmt.Printf("unable to convert data type to int, type: %v", t)
			return int(0)
		}
		return intValue
	default:
		t := reflect.TypeOf(val)
		fmt.Printf("unable to convert data type to int, type: %v", t)
		return int(0)
	}
}

func (v Variant) AsInt8() int8 {
	switch val := v.Value.(type) {
	case nil:
		return int8(0)
	case int:
		return int8(val)
	case int8:
		return val
	case int16:
		return int8(val)
	case int32:
		return int8(val)
	case int64:
		return int8(val)
	case string:
		int8Value, err := strconv.ParseInt(val, 10, 8)
		if err != nil {
			t := reflect.TypeOf(val)
			fmt.Printf("unable to convert data type to int8, type: %v", t)
			return int8(0)
		}
		return int8(int8Value)
	default:
		t := reflect.TypeOf(val)
		fmt.Printf("unable to convert data type to int8, type: %v", t)
		return int8(0)
	}
}

func (v Variant) AsInt16() int16 {
	switch val := v.Value.(type) {
	case nil:
		return int16(0)
	case int:
		return int16(val)
	case int8:
		return int16(val)
	case int16:
		return val
	case int32:
		return int16(val)
	case int64:
		return int16(val)
	case string:
		int16Value, err := strconv.ParseInt(val, 10, 16)
		if err != nil {
			t := reflect.TypeOf(val)
			fmt.Printf("unable to convert data type to int16, type: %v", t)
			return int16(0)
		}
		return int16(int16Value)
	default:
		t := reflect.TypeOf(val)
		fmt.Printf("unable to convert data type to int16, type: %v", t)
		return int16(0)
	}
}

func (v Variant) AsInt32() int32 {
	switch val := v.Value.(type) {
	case nil:
		return int32(0)
	case int:
		return int32(val)
	case int8:
		return int32(val)
	case int16:
		return int32(val)
	case int32:
		return val
	case int64:
		return int32(val)
	case float32:
		return int32(val)
	case float64:
		return int32(val)
	case string:
		int32Value, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			t := reflect.TypeOf(val)
			fmt.Printf("unable to convert data type to int32, type: %v", t)
			return int32(0)
		}
		return int32(int32Value)
	default:
		t := reflect.TypeOf(val)
		fmt.Printf("unable to convert data type to int32, type: %v", t)
		return int32(0)
	}
}

func (v Variant) AsInt64() int64 {
	switch val := v.Value.(type) {
	case nil:
		return int64(0)
	case int:
		return int64(val)
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	case int64:
		return val
	case float32:
		return int64(val)
	case float64:
		return int64(val)
	case string:
		int64Value, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			t := reflect.TypeOf(val)
			fmt.Printf("unable to convert data type to int64, type: %v", t)
			return int64(0)
		}
		return int64Value
	default:
		t := reflect.TypeOf(val)
		fmt.Printf("unable to convert data type to int64, type: %v", t)
		return int64(0)
	}
}

func (v Variant) AsFloat() float32 {
	switch val := v.Value.(type) {
	case nil:
		return float32(0)
	case float32:
		return val
	case float64:
		return float32(val)
	case string:
		floatValue, err := strconv.ParseFloat(val, 32)
		if err != nil {
			t := reflect.TypeOf(val)
			fmt.Printf("unable to convert data type to float32, type: %v", t)
			return float32(0)
		}
		return float32(floatValue)
	default:
		t := reflect.TypeOf(val)
		fmt.Printf("unable to convert data type to float32, type: %v", t)
		return float32(0)
	}
}

func (v Variant) AsFloat64() float64 {
	switch val := v.Value.(type) {
	case nil:
		return float64(0)
	case float32:
		return float64(val)
	case float64:
		return val
	case string:
		floatValue, err := strconv.ParseFloat(val, 64)
		if err != nil {
			t := reflect.TypeOf(val)
			fmt.Printf("unable to convert data type to float64, type: %v", t)
			return float64(0)
		}
		return floatValue
	default:
		t := reflect.TypeOf(val)
		fmt.Printf("unable to convert data type to float64, type: %v", t)
		return float64(0)
	}
}

func (v Variant) AsBool() bool {
	switch val := v.Value.(type) {
	case nil:
		return false
	case int:
		return v.Value.(int) == 1
	case int8:
		return v.Value.(int8) == 1
	case int16:
		return v.Value.(int16) == 1
	case int32:
		return v.Value.(int32) == 1
	case int64:
		return v.Value.(int64) == 1
	case string:
		value := strings.ToUpper(strings.Trim(v.Value.(string), " "))
		if value == "1" || value == "S" || value == "Y" {
			return true
		} else {
			return false
		}
	default:
		t := reflect.TypeOf(val)
		fmt.Printf("unable to convert data type to bool, type: %v", t)
		return false
	}
}

func (v Variant) AsDateTime() time.Time {
	switch v.Value.(type) {
	case nil:
		data, _ := time.Parse(time.DateTime, time.DateTime)
		return data
	case time.Time:
		return v.Value.(time.Time)
	default:
		fmt.Printf("unable to convert data type to time.")
		data, _ := time.Parse(time.DateTime, time.DateTime)
		return data
	}
}

func (v Variant) IsNull() bool {
	switch val := v.Value.(type) {
	case nil:
		return true
	case string:
		return val == ""
	default:
		return false
	}
}

func (v Variant) IsNotNull() bool {
	return !v.IsNull()
}

func IsPointer(value interface{}) bool {
	t := reflect.TypeOf(value)
	return t.Kind() == reflect.Ptr
}

func getScale(number float64) int {
	str := fmt.Sprintf("%f", number)

	// Dividir a string em partes inteira e decimal
	parts := strings.Split(str, ".")

	// Se houver uma parte decimal, retornar o número de dígitos na parte decimal
	if len(parts) == 2 {
		return len(parts[1])
	}

	// Caso contrário, o número não possui parte decimal (escala = 0)
	return 0
}
