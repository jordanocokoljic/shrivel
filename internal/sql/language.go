package sql

import "unicode"

// ShouldPad will indicate if the content of the provided rune slice is a SQL
// term that requires a ' ' before it.
func ShouldPad(word []rune) bool {
	for _, padded := range pad {
		if compare(word, padded) {
			return true
		}
	}

	return false
}

// GetShort will return the shortened version of a SQL keyword if applicable,
// or just the provided word if not.
func GetShort(word []rune) []rune {
	for _, shrunk := range shrink {
		if compare(word, shrunk.word) {
			return word[shrunk.frame[0]:shrunk.frame[1]]
		}
	}

	return word
}

// compare will case insensitively compare a rune slice and a string character
// by character.
func compare(r []rune, s string) bool {
	if len(r) != len(s) {
		return false
	}

	for i, char := range s {
		if char != unicode.ToLower(r[i]) {
			return false
		}
	}

	return true
}

var pad = []string{
	"create",
	"table",
	"integer",
	"primary",
	"key",
}

var shrink = []struct {
	word  string
	frame [2]int
}{
	{
		word:  "integer",
		frame: [2]int{0, 3},
	},
}
