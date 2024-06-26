package shrivel

import "unicode"

// Sql will minify the contents of src writing the result to dst. It wil return
// the number of bytes written to dst. dst should be large enough to contain
// the minified SQL.
//
// It can be used with two separate slices:
//
//	src := []rune("... sql chunk ...")
//	dst = make([]rune, len(src))
//	n := shrivel.Sql(dst, src)
//	dst = dst[:n]
//
// Or operate on src in place:
//
//	sql := []rune("... sql chunk ...")
//	sql = sql[:shrivel.Sql(sql, sql)]
//
// Sql will not verify the validity of the SQL provided in src.
func Sql(dst, src []rune) int {
	var (
		write = 0

		last   rune = 0
		target rune = 0

		skipLine   = false
		blockDepth = 0
	)

	for read := 0; read < len(src); read++ {
		char := src[read]

		if char == '\r' {
			continue
		}

		if blockDepth > 0 {
			if char == '/' && len(src) > read+1 && src[read+1] == '*' {
				blockDepth++
			}

			if char == '*' && len(src) > read+1 && src[read+1] == '/' {
				blockDepth--
				read++
			}

			continue
		}

		if char == '/' && len(src) > read+1 && src[read+1] == '*' {
			blockDepth++
			continue
		}

		if skipLine {
			if char == '\n' {
				skipLine = false
			}

			continue
		}

		if char == '-' && len(src) > read+1 && src[read+1] == '-' {
			skipLine = true
			continue
		}

		if target != 0 {
			if char == target {
				target = 0
			}

			dst[write] = char
			write++
			continue
		}

		if isQuote(char) {
			target = char
		}

		if (char == '(' || isQuote(char)) && unicode.IsSpace(dst[write-1]) {
			dst[write-1] = char
			last = char
			continue
		}

		if !unicode.IsSpace(char) {
			dst[write] = char
			write++
		}

		if unicode.IsSpace(char) && !unicode.IsSpace(last) && char != '\n' {
			dst[write] = char
			write++
		}

		last = char
	}

	return write
}

func isQuote(char rune) bool {
	return char == '"' || char == '\''
}
