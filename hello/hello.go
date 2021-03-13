package main

import "fmt"

var prefixes = map[string]string{
	"es": "Hola, ",
	"fr": "Bonjour, ",
	"en": "Hello, ",
}

func greetingPrefix(language string) (prefix string) {
	prefix = prefixes[language]
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "" {
		return greetingPrefix("en") + name
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("Alex", ""))
}
