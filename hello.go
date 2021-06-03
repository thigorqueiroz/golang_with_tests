package main

import (
	"fmt"
	"strings"
)

const englishPrefix = "Hello, %s."

func Hello(name ...string) string {
	if name == nil || strings.TrimSpace(name[0]) == "" {
		return "Hello World"
	}
	return fmt.Sprintf(englishPrefix, name[0])
}
