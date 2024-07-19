package sql

import (
	"slices"
	"unicode"
)

// Lexer is a lexeme emitter for SQL dialects.
type Lexer struct {
	// src stores the runes to iterate over.
	src []rune

	// pos stores the index the lexer is currently at.
	pos int
}

// NewLexer returns a new Lexer.
func NewLexer(src []rune) *Lexer {
	return &Lexer{
		src: src,
		pos: 0,
	}
}

// Next returns the next lexeme from the source provided to the Lexer.
func (l *Lexer) Next() []rune {
	depth := 0

	for !l.end() {
		char := l.src[l.pos]

		if char == '/' && l.nextIs('*') {
			depth++
			l.pos++
			continue
		}

		if depth > 0 {
			if char == '*' && l.nextIs('/') {
				depth--
				l.pos++
			}

			l.pos++
			continue
		}

		if char == '-' && l.nextIs('-') {
			for l.src[l.pos] != '\n' {
				l.pos++
			}

			continue
		}

		if char == '\'' {
			return l.eatUntil('\'')
		}

		if unicode.IsLetter(char) {
			return l.eatWord()
		}

		if unicode.IsDigit(char) {
			return l.eatNumber()
		}

		if !unicode.IsSpace(char) {
			return l.eatPunctuation()
		}

		l.pos++
	}

	return nil
}

// eatUntil will consume characters until the next instance of target.
func (l *Lexer) eatUntil(target rune) []rune {
	start := l.pos
	l.pos++

	for !l.end() {
		char := l.src[l.pos]

		if char == target {
			l.pos++
			break
		}

		l.pos++
	}

	return l.src[start:l.pos]
}

// eatWord will return the next word.
// A word is any continuous set of unicode letters or '_'.
func (l *Lexer) eatWord() []rune {
	start := l.pos

	for !l.end() && (unicode.IsLetter(l.src[l.pos]) || l.src[l.pos] == '_') {
		l.pos++
	}

	return l.src[start:l.pos]
}

// eatNumber will return the next number.
// A number is any continuous set of unicode digits, with at most one period.
func (l *Lexer) eatNumber() []rune {
	start := l.pos
	decimal := false

	for !l.end() {
		char := l.src[l.pos]

		if !unicode.IsDigit(char) {
			if char != '.' || decimal {
				break
			}

			decimal = true
		}

		l.pos++
	}

	return l.src[start:l.pos]
}

// complexOperators represents the multi-character comparison operators that
// the SQL language contains.
var complexOperators = [][2]rune{
	{'<', '='},
	{'>', '='},
	{'<', '>'},
	{'!', '='},
}

// eatPunctuation will return the next punctuation character.
func (l *Lexer) eatPunctuation() []rune {
	start := l.pos

	if !l.last() {
		set := (*[2]rune)(l.src[l.pos : l.pos+2])

		if slices.Contains(complexOperators, *set) {
			l.pos += 2
			return l.src[start:l.pos]
		}
	}

	char := l.src[start : l.pos+1]
	l.pos++
	return char
}

// end indicates if the lexer has exhausted src.
func (l *Lexer) end() bool {
	return l.pos >= len(l.src)
}

// last indicates if the lexer is currently on the last rune in src.
func (l *Lexer) last() bool {
	return l.pos+1 >= len(l.src)
}

// nextIs indicates if the next character in the lexer (if present) matches the
// provided character.
func (l *Lexer) nextIs(r rune) bool {
	return !l.last() && l.src[l.pos+1] == r
}
