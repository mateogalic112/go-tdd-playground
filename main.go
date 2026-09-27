package main

import (
	"os"

	"example.com/hello/di"
)

func main() {
	di.Greet(os.Stdout, "Lockie")
}
