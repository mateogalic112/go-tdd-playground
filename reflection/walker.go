// Package walker will teach us about reflection
package walker

import "reflect"

func walk(x any, fn func(input string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Struct:
		for _, field := range val.Fields() {
			walk(field.Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		for i := range val.Len() {
			walk(val.Index(i).Interface(), fn)
		}
	}
}

func getValue(x any) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
