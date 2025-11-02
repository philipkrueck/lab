package reflection

import (
	"reflect"
)

func Walk(x interface{}, fn func(input string)) {
	value := extractValue(x)

	walkValue := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}

	switch value.Kind() {
	case reflect.String:
		fn(value.String())
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			walkValue(value.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			walkValue(value.Index(i))
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			walkValue(value.MapIndex(key))
		}
	case reflect.Chan:
		val, ok := value.Recv()
		for ok {
			walkValue(val)
			val, ok = value.Recv()
		}
	case reflect.Func:
		valFnResult := value.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
}

func extractValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}
	return value
}
