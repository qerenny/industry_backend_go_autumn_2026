package main

import (
	"strings"
)

func greet(name string) string {
	cleaned := strings.TrimSpace(name)
	if cleaned == "" {
		cleaned = "World"
	}

	return "Hello, " + cleaned + "!"
}
