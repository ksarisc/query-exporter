package lib

import "unicode"

// StringIsWhitespace assess if string is
// null/empty or whitespace without allocations
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
