package main

import (
	"fmt"
	"strings"
)

const englishPrefix = "Hello, %s."

func grettingFromLanguage(lang string) (prefix string) {
	prefix = englishPrefix

	switch lang {
	case "spanish":
		prefix = "Holla, %s."
	case "portuguese":
		prefix = "Oi, %s."
	}
	return
}
func Hello(name, language string) string {
	if strings.TrimSpace(name) == "" {
		return "Hello World"
	}

	gretting := grettingFromLanguage(strings.ToLower(language))
	return fmt.Sprintf(gretting, name)
}
