package main

import (
	"fmt"
)

// Hello returns "Hello World"
func Hello(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println(Hello("Daniel"))
}
