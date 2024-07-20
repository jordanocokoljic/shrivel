package sql_test

import (
	"github.com/jordanocokoljic/shrivel/v2/internal/sql"
	"math/rand/v2"
	"testing"
	"unicode"
)

func TestShouldPad_Valid(t *testing.T) {
	tests := []string{
		"create",
		"table",
		"integer",
		"primary",
		"key",
	}

	for _, word := range tests {
		t.Run(word, func(t *testing.T) {
			fuzzed := fuzz(word)
			pad := sql.ShouldPad(fuzz(word))
			if !pad {
				t.Errorf("returned false, fuzzed: %s", string(fuzzed))
			}
		})
	}
}

func TestShouldPad_Invalid(t *testing.T) {
	pad := sql.ShouldPad([]rune("not_a_word"))
	if pad {
		t.Errorf("returned true")
	}
}

func TestGetShort(t *testing.T) {
	tests := []struct {
		word   string
		result string
		frame  [2]int
	}{
		{
			word:   "integer",
			result: "int",
			frame:  [2]int{0, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			short := sql.GetShort([]rune(tt.word))
			if string(short) != tt.result {
				t.Errorf("got %s, expected %s", string(short), tt.result)
			}

			fuzzed := fuzz(tt.word)
			short = sql.GetShort(fuzzed)
			expected := fuzzed[tt.frame[0]:tt.frame[1]]

			if string(short) != string(expected) {
				t.Errorf(
					"got %s, expected %s: fuzzed %s",
					string(short), string(expected), string(fuzzed),
				)
			}
		})
	}
}

func fuzz(s string) []rune {
	r := []rune(s)

	for i := range r {
		if rand.Float32() > 0.5 {
			r[i] = unicode.ToUpper(r[i])
		}
	}

	return r
}
