package main

import "fmt"

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

func main() {
	fmt.Println(Hello())
	fmt.Println(Hello2("Chris"))
	fmt.Println(Hello3(""))
	fmt.Println(Hello4("Elodie", "Spanish"))
	fmt.Println(Hello4("Divyue", "French"))
}
