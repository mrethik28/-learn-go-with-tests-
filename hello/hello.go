package hello

import "strings"

func Hello(name string, language string) string {
	greetings := map[string]string{
		"English": "Hello",
		"Spanish": "Hola",
		"French":  "Bonjour",
	}
	if name == "" {
		name = "World"
	}
	if strings.TrimSpace(language) == "" {
		language = "English"
	}
	return greetings[strings.TrimSpace(language)] + " " + name
}
