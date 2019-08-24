package main

import (
	"fmt"
)

var englishHelloPrefix = "Hello "

// Hello returns "Hello World"
func Hello(name string) string {

	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Daniel"))
}
