// Package mocking will teach us about mocking in Go
package mocking

import (
	"fmt"
	"io"
)

func Countdown(out io.Writer) {
	fmt.Fprintf(out, "3")
}
