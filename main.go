package main

import (
	"os"

	"example.com/hello/mocking"
)

func main() {
	mocking.Countdown(os.Stdout)
}
