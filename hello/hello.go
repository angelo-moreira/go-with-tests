package main

import "fmt"

const french = "French"
const spanish = "Spanish"

const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	switch language {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func main() {
	fmt.Println(Hello("World", ""))
}
