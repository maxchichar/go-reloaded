package functions

import (
    "regexp"
    "strings"
)

func ProcessQuotes(tokens []string) []string {
	if len(tokens) == 0 {
		return tokens
	}
	
    text := strings.Join(tokens, " ")
	text = regexp.MustCompile(`'\s*(.*?)\s*'`).ReplaceAllString(text, "'$1'")

    // re := regexp.MustCompile(`'\s*(.*?)\s*'`) // \s* --> zero or more spaces. . --> the character in the group. ? --> get as much as possible element in the ' '
    // result := re.ReplaceAllString(text, "'$1'") // just like strings.ReplaceAll
	
    return strings.Fields(text)
}