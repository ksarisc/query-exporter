package lib

import "unicode"

// StringIsWhitespace assess if string is
// null/empty or whitespace without allocations
func StringIsWhitespace(value string) bool {
	if value == "" {
		return true
	}
	// cannot loop for i := range value due
	// to value[i] being byte rather than rune
	for _, v := range value {
		if !unicode.IsSpace(v) {
			return false
		}
	}
	return true
} // END StringIsWhitespace
