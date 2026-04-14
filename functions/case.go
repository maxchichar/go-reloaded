package functions

import (
	"strconv"
	// "fmt"
	"strings"
)

// To upper case function
func to_upper(word string) string {
	if len(word) == 0 {
		return word
	}

	return strings.ToUpper(word)
}

// To lower case function
func to_lower(word string) string {
	if len(word) == 0 {
		return word
	}

	return strings.ToLower(word)
}

// To capitalise
func to_cap(word string) string {
	if len(word) == 0 {
		return word
	}

	return strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
}

// For processing the upper, lower and capitalizing words with number
func ProcessCase(tokens []string) []string {
	result := make([]string, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		
		current := tokens[i]

		if isCaseMarker(current) {
			action, count := parseMarker(current)
			// fmt.Printf("DEBUG Marker: %s → Action: %s, Count: %d\n", current, action, count)

			var transform func(string) string
			switch action {
			case "up":
				transform = to_upper // transform becomes the container for the uppercase function
			case "low":
				transform = to_lower // transform becomes the container for the lowerrcase function
			case "cap":
				transform = to_cap // transform becomes the container for the capitalize function
			default:
				result = append(result, current)
				continue
			}
			for n := 0; n < count; n++ {
				idx := len(result) - 1 - n
				if idx < 0 {
					break
				}
				result[idx] = transform(result[idx])
			}
			continue
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
	// if cleaned == "up" || cleaned == "low" || cleaned == "cap" {
	// 	return true
	// }
	if strings.Contains(cleaned, "up") ||
	strings.Contains(cleaned, "low") ||
	strings.Contains(cleaned, "cap") {
		return true
	}

	return false
}

// Detecting markers like (up, 3) (low, 2)
func parseMarker(token string) (string, int) {
	cleaned := strings.ToLower(strings.Trim(token,"()"))

	if strings.Contains(cleaned, ",") {
		parts := strings.Split(cleaned, ",")

		if len(parts) >= 2{
			action := strings.TrimSpace(parts[0]) // represent the transformer e.g up low and cap
			countStr := strings.TrimSpace(parts[1]) // represents the number
			if count, err := strconv.Atoi(countStr); err == nil && count > 0 { // convert the string to int
				return action, count
			}
			return action, 1
		}
		
	}
	return cleaned, 1
}