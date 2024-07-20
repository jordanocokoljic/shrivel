package shrivel

import (
	"github.com/jordanocokoljic/shrivel/v2/internal/sql"
	"unicode"
)

// Sql will minify the contents of src writing the result to dst. It wil return
// the number of bytes written to dst. dst should be large enough to contain
// the minified SQL.
func Sql(dst, src []rune) (int, error) {
	var (
		offset = 0
		lexer  = sql.NewLexer(src)

		start, end int
	)

	if word := lexer.Next(); word != nil {
		start = offset
		offset += copy(dst[offset:], word)
		end = offset
	}

	for word := lexer.Next(); word != nil; word = lexer.Next() {
		if shouldPad(dst[start:end], word) {
			dst[offset] = ' '
			offset++
		}

		start = offset
		offset += copy(dst[offset:], sql.GetShort(word))
		end = offset
	}

	if dst[offset-1] == ';' {
		offset--
	}

	return offset, nil
}

// shouldPad will indicate if the current word needs padding.
func shouldPad(last, current []rune) bool {
	for _, r := range current {
		if !(unicode.IsLetter(r) || unicode.IsNumber(r)) {
			return false
		}
	}

	return sql.ShouldPad(last) || sql.ShouldPad(current)
}
