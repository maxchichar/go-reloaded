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
			cleaned := strings.ToLower(strings.Trim(current, "()"))
			
			switch {
			case strings.HasPrefix(cleaned, "up"):
				if len(result) > 0{
					lastWord := len(result) - 1
					result[lastWord] = to_upper(result[lastWord])
				}
				continue

			case strings.HasPrefix(cleaned, "low"):
				if len(result) > 0 {
					lastWord := len(result) - 1
					result[lastWord] = to_lower(result[lastWord])
				}
				continue

			case strings.HasPrefix(cleaned, "cap"):
				if len(result) > 0{
					lastIndex := len(result) - 1
					result[lastIndex] = to_capitalise(result[lastIndex])
				}
				continue
			}
		} else {
			result = append(result, current)
		}
	}
	return result
}

// Detecting markers like (up) (low) and (cap)
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