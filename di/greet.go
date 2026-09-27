// Package di will teach use about dependency injection in Go
package di

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
