package godataset

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	go_ora "github.com/sijms/go-ora/v2"
)

type Variant struct {
	IDataType *DataType
	Value     any
	Silent    bool
}

func (v Variant) SetSilent(value bool) Variant {
	v.Silent = value
	return v
}

func (v Variant) AsValue() any {
	return v.Value
}

func (v Variant) AsString() string {
	if v.Value == nil {
		return ""
	}

	// Usar type assertion diretamente em vez de switch type para melhor performance
	switch val := v.Value.(type) {
	case string:
		return val
	case time.Time:
		return val.String()
	case int:
		return strconv.Itoa(val)
	case int8:
		return strconv.Itoa(int(val))
	case int16:
		return strconv.Itoa(int(val))
	case int32:
		return strconv.Itoa(int(val))
	case int64:
		return strconv.FormatInt(val, 10)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case []uint8:
		return string(val)
	case go_ora.Clob:
		return val.String
	case bool:
		return strconv.FormatBool(val)
	default:
		// Fallback para reflection apenas quando necessário
		t := reflect.TypeOf(v.Value)
		msg := fmt.Sprintf("unable to convert data type to string. Type: %v", t)
		if v.Silent {
			fmt.Println(msg)
		} else {
			panic(msg)
		}
		return ""
	}
}

func (v Variant) AsStringNil() *string {
	if v.Value == nil {
		return nil
	}

	str := v.AsString()
	return &str
}

func (v Variant) AsInt() int {
	if v.Value == nil {
		return 0
	}

	// Usar type assertion diretamente para melhor performance
	switch val := v.Value.(type) {
	case int:
		return val
	case int8:
		return int(val)
	case int16:
		return int(val)
	case int32:
		return int(val)
	case int64:
		return int(val)
	case uint:
		return int(val)
	case uint8:
		return int(val)
	case uint16:
		return int(val)
	case uint32:
		return int(val)
	case uint64:
		return int(val)
	case float32:
		return int(val)
	case float64:
		return int(val)
	case string:
		if intValue, err := strconv.Atoi(val); err == nil {
			return intValue
		}
		v.logError("unable to convert string to int", val)
		return 0
	case bool:
		if val {
			return 1
		}
		return 0
	default:
		v.logError("unable to convert data type to int", val)
		return 0
	}
}

func (v Variant) AsIntNil() *int {
	if v.Value == nil {
		return nil
	}

	intVal := v.AsInt()
	return &intVal
}

func (v Variant) AsInt8() int8 {
	if v.Value == nil {
		return 0
	}

	switch val := v.Value.(type) {
	case int8:
		return val
	case int:
		return int8(val)
	case int16:
		return int8(val)
	case int32:
		return int8(val)
	case int64:
		return int8(val)
	case uint:
		return int8(val)
	case uint8:
		return int8(val)
	case uint16:
		return int8(val)
	case uint32:
		return int8(val)
	case uint64:
		return int8(val)
	case float32:
		return int8(val)
	case float64:
		return int8(val)
	case string:
		if int8Value, err := strconv.ParseInt(val, 10, 8); err == nil {
			return int8(int8Value)
		}
		v.logError("unable to convert string to int8", val)
		return 0
	case bool:
		if val {
			return 1
		}
		return 0
	default:
		v.logError("unable to convert data type to int8", val)
		return 0
	}
}

func (v Variant) AsInt8Nil() *int8 {
	if v.Value == nil {
		return nil
	}

	int8Val := v.AsInt8()
	return &int8Val
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
	case uint:
		return int16(val)
	case uint8:
		return int16(val)
	case uint16:
		return int16(val)
	case uint32:
		return int16(val)
	case uint64:
		return int16(val)
	case string:
		int16Value, err := strconv.ParseInt(val, 10, 16)
		if err != nil {
			t := reflect.TypeOf(val)
			msg := fmt.Sprintf("unable to convert data type to int16. Type: %v", t)
			if v.Silent {
				fmt.Println(msg)
			} else {
				panic(msg)
			}
			return int16(0)
		}
		return int16(int16Value)
	default:
		t := reflect.TypeOf(val)
		msg := fmt.Sprintf("unable to convert data type to int16. Type: %v", t)
		if v.Silent {
			fmt.Println(msg)
		} else {
			panic(msg)
		}
		return int16(0)
	}
}
func (v Variant) AsInt16Nil() *int16 {
	valor := v.AsValue()
	var tvalor any

	if valor == nil {
		return nil
	} else {
		tvalor = v.AsInt16()
		t, ok := tvalor.(int16)
		if ok {
			return &t
		}
		return nil
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
	case uint:
		return int32(val)
	case uint8:
		return int32(val)
	case uint16:
		return int32(val)
	case uint32:
		return int32(val)
	case uint64:
		return int32(val)
	case float32:
		return int32(val)
	case float64:
		return int32(val)
	case string:
		int32Value, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			t := reflect.TypeOf(val)
			msg := fmt.Sprintf("unable to convert data type to int32. Type: %v", t)
			if v.Silent {
				fmt.Println(msg)
			} else {
				panic(msg)
			}
			return int32(0)
		}
		return int32(int32Value)
	default:
		t := reflect.TypeOf(val)
		msg := fmt.Sprintf("unable to convert data type to int32. Type: %v", t)
		if v.Silent {
			fmt.Println(msg)
		} else {
			panic(msg)
		}
		return int32(0)
	}
}
func (v Variant) AsInt32Nil() *int32 {
	valor := v.AsValue()
	var tvalor any

	if valor == nil {
		return nil
	} else {
		tvalor = v.AsInt32()
		t, ok := tvalor.(int32)
		if ok {
			return &t
		}
		return nil
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
	case uint:
		return int64(val)
	case uint8:
		return int64(val)
	case uint16:
		return int64(val)
	case uint32:
		return int64(val)
	case uint64:
		return int64(val)
	case float32:
		return int64(val)
	case float64:
		return int64(val)
	case string:
		int64Value, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			t := reflect.TypeOf(val)
			msg := fmt.Sprintf("unable to convert data type to int64. Type: %v", t)
			if v.Silent {
				fmt.Println(msg)
			} else {
				panic(msg)
			}
			return int64(0)
		}
		return int64Value
	default:
		t := reflect.TypeOf(val)
		msg := fmt.Sprintf("unable to convert data type to int64. Type: %v", t)
		if v.Silent {
			fmt.Println(msg)
		} else {
			panic(msg)
		}
		return int64(0)
	}
}
func (v Variant) AsInt64Nil() *int64 {
	valor := v.AsValue()
	var tvalor any

	if valor == nil {
		return nil
	} else {
		tvalor = v.AsInt64()
		t, ok := tvalor.(int64)
		if ok {
			return &t
		}
		return nil
	}
}
func (v Variant) AsFloat() float32 {
	switch val := v.Value.(type) {
	case nil:
		return float32(0)
	case int:
		return float32(val)
	case int8:
		return float32(val)
	case int16:
		return float32(val)
	case int32:
		return float32(val)
	case int64:
		return float32(val)
	case uint:
		return float32(val)
	case uint8:
		return float32(val)
	case uint16:
		return float32(val)
	case uint32:
		return float32(val)
	case uint64:
		return float32(val)
	case float32:
		return val
	case float64:
		return float32(val)
	case string:
		floatValue, err := strconv.ParseFloat(val, 32)
		if err != nil {
			t := reflect.TypeOf(val)
			msg := fmt.Sprintf("unable to convert data type to float32. Type: %v", t)
			if v.Silent {
				fmt.Println(msg)
			} else {
				panic(msg)
			}
			return float32(0)
		}
		return float32(floatValue)
	default:
		t := reflect.TypeOf(val)
		msg := fmt.Sprintf("unable to convert data type to float32. Type: %v", t)
		if v.Silent {
			fmt.Println(msg)
		} else {
			panic(msg)
		}
		return float32(0)
	}
}
func (v Variant) AsFloatNil() *float32 {
	valor := v.AsValue()
	var tvalor any

	if valor == nil {
		return nil
	} else {
		tvalor = v.AsFloat()
		t, ok := tvalor.(float32)
		if ok {
			return &t
		}
		return nil
	}
}
func (v Variant) AsFloat64() float64 {
	switch val := v.Value.(type) {
	case nil:
		return float64(0)
	case int:
		return float64(val)
	case int8:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case uint:
		return float64(val)
	case uint8:
		return float64(val)
	case uint16:
		return float64(val)
	case uint32:
		return float64(val)
	case uint64:
		return float64(val)
	case float32:
		return float64(val)
	case float64:
		return val
	case string:
		floatValue, err := strconv.ParseFloat(val, 64)
		if err != nil {
			t := reflect.TypeOf(val)
			msg := fmt.Sprintf("unable to convert data type to float64. Type: %v", t)
			if v.Silent {
				fmt.Println(msg)
			} else {
				panic(msg)
			}
			return float64(0)
		}
		return floatValue
	default:
		t := reflect.TypeOf(val)
		msg := fmt.Sprintf("unable to convert data type to float64. Type: %v", t)
		if v.Silent {
			fmt.Println(msg)
		} else {
			panic(msg)
		}
		return float64(0)
	}
}
func (v Variant) AsFloat64Nil() *float64 {
	valor := v.AsValue()
	var tvalor any

	if valor == nil {
		return nil
	} else {
		tvalor = v.AsFloat64()
		t, ok := tvalor.(float64)
		if ok {
			return &t
		}
		return nil
	}
}
func (v Variant) AsBool() bool {
	switch val := v.Value.(type) {
	case nil:
		return false
	case bool:
		return v.Value.(bool)
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
		msg := fmt.Sprintf("unable to convert data type to bool. Type: %v", t)
		if v.Silent {
			fmt.Println(msg)
		} else {
			panic(msg)
		}
		return false
	}
}
func (v Variant) AsBoolNil() *bool {
	valor := v.AsValue()
	var tvalor any

	if valor == nil {
		return nil
	} else {
		tvalor = v.AsBool()
		t, ok := tvalor.(bool)
		if ok {
			return &t
		}
		return nil
	}
}
func (v Variant) AsDateTime() time.Time {
	switch val := v.Value.(type) {
	case nil:
		data, _ := time.Parse(time.DateTime, time.DateTime)
		return data
	case time.Time:
		return v.Value.(time.Time)
	case string:
		dataLocal, err := dateparse.ParseAny(val)

		if err != nil {
			preferMonthFirstFalse := dateparse.PreferMonthFirst(false)
			dataLocal, err = dateparse.ParseAny(val, preferMonthFirstFalse)

			if err != nil {
				msg := fmt.Sprintf("unable to convert data type to time.")
				if v.Silent {
					fmt.Println(msg)
				} else {
					panic(msg)
				}
				data, _ := time.Parse(time.DateTime, time.DateTime)
				return data
			}
		}

		return dataLocal
	default:
		msg := fmt.Sprintf("unable to convert data type to time. ")
		if v.Silent {
			fmt.Println(msg)
		} else {
			panic(msg)
		}
		data, _ := time.Parse(time.DateTime, time.DateTime)
		return data
	}
}

func (v Variant) AsDateTimeNil() *time.Time {
	valor := v.AsValue()
	var tvalor any

	if valor == nil {
		return nil
	} else {
		tvalor = v.AsDateTime()
		t, ok := tvalor.(time.Time)
		if ok {
			return &t
		}
		return nil
	}
}

func (v Variant) AsByte() []byte {
	switch val := v.Value.(type) {
	case nil:
		return nil
	case []byte:
		return v.Value.([]byte)
	case string:
		return []byte(v.AsString())
	case *go_ora.Clob:
		return []byte(val.String)
	default:
		t := reflect.TypeOf(val)
		msg := fmt.Sprintf("unable to convert data type to byte. Type: %v", t)
		if v.Silent {
			fmt.Println(msg)
		} else {
			panic(msg)
		}
		return nil
	}
}

func (v Variant) AsByteNil() *[]byte {
	valor := v.AsValue()
	var tvalor any

	if valor == nil {
		return nil
	} else {
		tvalor = v.AsByte()
		t, ok := tvalor.([]byte)
		if ok {
			return &t
		}
		return nil
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

// Método helper para logging de erros
func (v Variant) logError(message string, value any) {
	t := reflect.TypeOf(value)
	msg := fmt.Sprintf("%s. Type: %v", message, t)
	if v.Silent {
		fmt.Println(msg)
	} else {
		panic(msg)
	}
}
