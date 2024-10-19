package main

import "fmt"

const englishHelloPrefix = "Hola "

func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("mundo"))
}
