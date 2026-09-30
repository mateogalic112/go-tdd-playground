// Package walker will teach us about reflection
package walker

import "reflect"

func walk(x any, fn func(input string)) {
	val := reflect.ValueOf(x)

	for _, field := range val.Fields() {
		fn(field.String())
	}
}
