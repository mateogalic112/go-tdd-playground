package main

import "fmt"

const (
	spanish            = "Spanish"
	spanishHelloPrefix = "Hola, "

	french            = "French"
	frenchHelloPrefix = "Bonjour, "

	croatian            = "Croatian"
	croatianHelloPrexis = "Bok, "

	englishHelloPrefix = "Hello, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case croatian:
		prefix = croatianHelloPrexis
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
