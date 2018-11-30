package main

import "fmt"

const englishPrefix = "Hello, "
const frenchPrefix = "Bonjour "
const spanishPrefix = "Hola "

func languagePrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchPrefix
	case "Spanish":
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return (languagePrefix(language) + name)
}

func main() {
	fmt.Printf(Hello("Hitesh", "French"))
}
