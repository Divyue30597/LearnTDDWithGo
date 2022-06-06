package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello() string {
	return englishHelloPrefix + "world"
}

func Hello2(name string) string {
	return englishHelloPrefix + name
}

func Hello3(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func Hello4(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return spanishHelloPrefix + name
	}

	if language == "French" {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func Hello5(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	// replacing multiple if's with switch
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}

	return prefix + name
}

// above code further be refactored to two different functions

func Hello6(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

// named return value as `prefix string`
// This will create a variable called prefix in the function and will be assigned to "" for string in case of int it will be initialised to 0
// we can return whatever it's set to by just calling `return` in the function rather than `return prefix`
// 
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello())
	fmt.Println(Hello2("Chris"))
	fmt.Println(Hello3(""))
	fmt.Println(Hello4("Elodie", "Spanish"))
	fmt.Println(Hello4("Divyue", "French"))
	fmt.Println(Hello5("Tomchi", spanish))
	fmt.Println(Hello5("Tom", french))
	fmt.Println(Hello5("Tomchi", ""))
	fmt.Println(Hello6("Tomchi", spanish))
	fmt.Println(Hello6("Tom", french))
	fmt.Println(Hello6("Tomchi", ""))
}
