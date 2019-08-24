package main

import (
	"fmt"
)

const spanish = "Spanish"
const german = "Deutsch"
const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const GermanHelloPrefix = "Hallo "

// Hello returns "Hello World"
func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrexif(language) + name
}

func greetingPrexif(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case german:
		prefix = GermanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Daniel", ""))
}
