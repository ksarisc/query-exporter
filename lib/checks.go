package lib

import "unicode"

// StringIsWhitespace without allocations assess
// if string is null/empty or whitespace
func StringIsWhitespace(value string) bool {
	if value == "" {
		return true
	}
	for _, v := range value {
		if !unicode.IsSpace(v) {
			return false
		}
	}
	return true
} // END StringIsWhitespace
