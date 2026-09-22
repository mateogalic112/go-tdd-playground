// Package iteration applies repetitive logic to some data
package iteration

import "strings"

func Repeat(character string, iterations int) string {
	var repeated strings.Builder
	for range iterations {
		repeated.WriteString(character)
	}
	return repeated.String()
}
