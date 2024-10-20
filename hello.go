package main

import "fmt"

const englishHelloPrefix = "Hola, "

func Hello(name string) string {
	if name == "" {
		name = "Mundo"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("mundo"))
}
