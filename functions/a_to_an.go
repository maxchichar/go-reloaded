package functions

import (
	"strings"
)

func ProcessAnRule(tokens []string) []string {
	result := make([]string, 0, len(tokens))

	for i := 0; i < len(tokens); i++{
		current := tokens[i]

		if isA(current) {
			if i+1 < len(tokens) { // to check if there's a next word
				next := tokens[i+1]
				if startWithVowel(next) {
					current = "an"
				} else {
					current = "a"
				}
				result = append(result, current)
			}
		}else {
			result = append(result, current)
		}
	}
	return result

}

func isA(token string) bool {
	token = strings.ToLower(strings.TrimSpace(token))
	return token == "a"
}

func startWithVowel(token string) bool {
	if len(token) == 0{
		return false
	}

	first := strings.ToLower(string(token[0]))
	return strings.Contains("aeiouh", first)
}