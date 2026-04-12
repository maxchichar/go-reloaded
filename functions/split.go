package functions

import (
	"strings"
	"unicode"
)

func is_space(r rune) bool {
	return unicode.IsSpace(r)
}

func SplitElement(word string) []string {
	if word == ""{
		return []string{}
	}

	return strings.FieldsFunc(word, is_space)
}

// minimal method
/*
func splitText(text string) []string {
	return strings.FieldsFunc(text, unicode.IsSpace)
}
*/