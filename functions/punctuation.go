package functions

import (
	"regexp"
	"strings"
) 

func ProcessPunctuation(tokens []string) []string {

	good := strings.Join(tokens, " ")

	re := regexp.MustCompile(`\s+([.,;!:?])`)
	
	good = re.ReplaceAllString(good, "$1")
	
	re2 := regexp.MustCompile(`([.,;!:?])(\s*)(\w)`)

	result := re2.ReplaceAllString(good, "$1 $3")

	return strings.Fields(result)

	// text := strings.Join(tokens, " ")

	// // Handle ellipsis ... . . .
	// text = regexp.MustCompile(`\.\s*\.\s*\.`).ReplaceAllString(text, "__ELLIPSIS__")

	// // Handle colon
	// text = regexp.MustCompile(`\s*:\s*`).ReplaceAllString(text, ": ")
	
	// // Handle Punctuations
	// text = regexp.MustCompile(`\s+([,.:;!?])`).ReplaceAllString(text, "$1")
	// text = regexp.MustCompile(`([,.:;!?])([A-Za-z0-9])`).ReplaceAllString(text, "$1 $2")

	// // Handle single quote
	// text = regexp.MustCompile(`'\s+`).ReplaceAllString(text, "'")
	// text = regexp.MustCompile(`\s+'`).ReplaceAllString(text, "'")

	// // Handle multiple punctuation
	// text = regexp.MustCompile(`([!?])\s+([!?])`).ReplaceAllString(text, "$1$2")

	// // Restore ellipsis
	// text = strings.ReplaceAll(text, "__ELLIPSIS__", "...")

	// text = strings.Join(strings.Fields(text), " ") // clean multiple spaces
	// // text = strings.TrimSpace(text)

	// return strings.Fields(text)
}



