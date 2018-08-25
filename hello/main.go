package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const thai = "Thai"
const english = "English"
const englishPrefix = "Hello, "
const thaiHelloPrefix = "Sawaddee, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case thai:
		prefix = thaiHelloPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Nan", "Thai"))
}
