package functions

import (
	"strconv"
)

func hex_to_dec(hexStr string) (int64, error) {
	n, err := strconv.ParseInt(hexStr, 16, 64) // conversion from hexadecimal to decimal
	// error handling
	if err != nil{
		return n, err
	}
	return n, nil
}

func bin_to_dec(binStr string) (int64, error) {
	n, err := strconv.ParseInt(binStr, 2, 64) // conversion of binary to decimal
	// error handling
	if err != nil {
		return n, err
	}
	return n, nil
}

func ProcessHexBin(token []string) []string {
	result := make([]string, 0,len(token))

	for i := 0; i < len(token); i++ {

		if i+1 < len(token) {

			nextToken := token[i+1]

			switch nextToken {
			case "(hex)":
				transHextoDec, err := hex_to_dec(token[i]) // calling hex_to_dec for conversion

				if err == nil {
					transHtoD := strconv.Itoa(int(transHextoDec)) // converts the integer to string
					result = append(result, transHtoD) // adds the converted to an empty string result
				} else{
					result = append(result, token[i]) // leave it the same way. if it not a hexadecimal
				}
				i++ // removes the marker by skipping it in the loop
				continue
			case "(bin)":
				transBintoDec, err := bin_to_dec(token[i])

				if err == nil {
					transBtoD := strconv.Itoa(int(transBintoDec)) // converts the integer to alphabet aka string
					result = append(result, transBtoD)
				} else{
					result = append(result, token[i]) // leaves the figure same way if not a binary
				}
				i++ // skipping marker
				continue
			}
		}
		result = append(result, token[i]) // leave as same if not (hex) or (bin)
	}
	return result // produce processed token
}
