package functions

import (
	"regexp"
)

var tokenRegex = regexp.MustCompile(`\([^)]+\)|\S+`)

func SplitElement(text string) []string {
	if len(text) == 0 {
		return []string{}
	}

	return tokenRegex.FindAllString(text, -1)
}

// minimal method, if i'm not handling (up,n)
/*
func splitText(text string) []string {
	return strings.FieldsFunc(text, unicode.IsSpace)
}
*/
