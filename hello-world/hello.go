package main

import (
	"fmt"
)

const (
	portuguese = "Portuguese"
	french     = "French"

	englishHelloPrefix    = "Hello, "
	portugueseHelloPrefix = "Olá, "
	frenchHelloPrefix     = "Bonjour, "

	englishDefaultName    = "world"
	portugueseDefaultName = "mundo"
	frenchDefaultName     = "monde"
)

func Hello(name, language string) string {
	if name == "" {
		name = defaultName(language)
	}

	return greetingPrefix(language) + name + "!"
}

func defaultName(language string) (name string) {
	switch language {
	case portuguese:
		name = portugueseDefaultName
	case french:
		name = frenchDefaultName
	default:
		name = englishDefaultName
	}

	return
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case portuguese:
		prefix = portugueseHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
