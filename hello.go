package main

const (
	spanish = "spanish"
	french  = "french"

	englishHelloPrefix = "hello "
	spanishHelloPrefix = "hola "
	frenchHelloPrefix  = "bonjour "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
