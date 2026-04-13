package functions

import (
	"strings"
)

func to_upper(word string) string {
	if len(word) == 0 {
		return word
	}

	return strings.ToUpper(word)
}

func to_lower(word string) string {
	if len(word) == 0 {
		return word
	}

	return strings.ToLower(word)
}

func to_capitalise(word string) string {
	if len(word) == 0 {
		return word
	}

	return strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
}

func ProcessCase(tokens []string) []string {
	result := make([]string, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		
		current := tokens[i]

		if isCaseMarker(current) {
			
			var transform func(string) string
			switch current {
			case "up":
				transform = to_upper
				result = append(result, transform(tokens[i-1]))
				i++
				continue
			// case "low"
			}
		}
	}
	return result
}

// Checking for markers
func isCaseMarker(token string) bool {
	if len(token) < 3{
		return false
	}

	if token[0] != '(' || token[len(token)-1] != ')' {
		return false
	}

	cleaned := strings.ToLower(strings.Trim(token, "()"))
	if strings.Contains(cleaned, "up") ||
	strings.Contains(cleaned, "low") ||
	strings.Contains(cleaned, "cap") {
		return true
	}

	return false
}

// func parseMarker(token string) bool {

// }