package lex

import (
	//"fmt"
	"fmt"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

const eof = -1

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*Scanner) stateFn

// Scanner holds the state of the scanner.
type Scanner struct {
	tokens     chan Token // channel of scanned items
	r          io.ByteReader
	done       bool
	name       string // the name of the input; used only for error reports
	buf        []byte
	input      string  // the line of text being scanned
	leftDelim  string  // start of action
	rightDelim string  // end of action
	state      stateFn // the next lexing function to enter
	line       int     // line number in input
	pos        int     // current position of the input
	start      int     // start position of this item
	width      int     // width of last rune read from input
}

// loadLine reads the next line of input and stores it in (appends it to) the input.
// (l.input may have data left over when we are called.)
// It strips carriage returns to make subsequent processing simpler.
func (l *Scanner) loadLine() {
	l.buf = l.buf[:0]
	for {
		c, err := l.r.ReadByte()
		if err != nil {
			l.done = true
			break
		}
		if c != '\r' {
			l.buf = append(l.buf, c)
		}
		if c == '\n' {
			break
		}
	}
	l.input = l.input[l.start:l.pos] + string(l.buf)
	l.pos -= l.start
	l.start = 0
}

func (l *Scanner) next() rune {
	if !l.done && l.pos == len(l.input) {
		l.loadLine()
	}
	if len(l.input) == l.pos {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	l.pos += l.width
	return r
}

// peek returns but does not consume the next rune in the input.
func (l *Scanner) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *Scanner) backup() {
	l.pos -= l.width
}

// emit passes an item back to the client.
func (l *Scanner) emit(t Type) {
	if t == NEWLINE {
		l.line++
	}
	s := l.input[l.start:l.pos]
	l.tokens <- Token{Type: t, Line: l.line, Text: s}
	l.start = l.pos
	l.width = 0
}

// ignore skips over the pending input before this point.
func (l *Scanner) ignore() {
	l.start = l.pos
}

// accept consumes the next rune if it's from the valid set.
func (l *Scanner) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *Scanner) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {

	}
	l.backup()
}

// errorf returns an error token and continues to scan.
func (l *Scanner) errorf(format string, args ...interface{}) stateFn {
	l.tokens <- Token{Type: ILLEGAL, Line: l.start, Text: fmt.Sprintf(format, args...)}
	return lexAny
}

// New creates a new scanner for the input string.
func New(name string, r io.ByteReader) *Scanner {
	l := &Scanner{
		r:      r,
		name:   name,
		line:   1,
		tokens: make(chan Token, 2), // We need a little room to save tokens.
		state:  lexAny,
	}
	return l
}

// Next returns the next token.
func (l *Scanner) Next() Token {
	// The lexer is concurrent but we don't want it to run in parallel
	// with the rest of the interpreter, so we only run the state machine
	// when we need a token.
	for l.state != nil {
		select {
		case tok := <-l.tokens:
			return tok
		default:
			// Run the machine
			l.state = l.state(l) // Magic!!!
		}
	}
	if l.tokens != nil {
		close(l.tokens)
		l.tokens = nil
	}
	return Token{Type: EOF, Line: l.pos, Text: "EOF"}
}

// state functions

// lexAny scans non-space items.
func lexAny(l *Scanner) stateFn {
	switch r := l.next(); {
	case r == eof:
		return nil
	case r == '\n' || r == '\r':
		l.emit(NEWLINE)
		if l.peek() == '\n' || l.peek() == '\r' {
			l.next()
		}
		return lexAny
	case r == ';':
		l.emit(SEMICOLON)
		return lexComment
	case isSpace(r):
		return lexSpace
	case '0' <= r && r <= '9':
		l.backup()
		return lexNumber
	case r == '+':
		l.emit(PLUS)
		return lexAny
	case r == '-':
		l.emit(MINUS)
		return lexAny
	case r == '*':
		l.emit(ASTERISK)
		return lexAny
	case r == '/':
		l.emit(FSLASH)
		return lexAny
	case r == '%':
		l.emit(PERCENT)
		return lexAny
	case r == '(':
		l.emit(LPAREN)
		return lexAny
	case r == ')':
		l.emit(RPAREN)
		return lexAny
	case r == '.':
		l.emit(DOT)
		return lexAny
	case r == ',':
		l.emit(COMMA)
		return lexAny
	case r == '#':
		l.emit(HASH)
		return lexAny
	case r == '$':
		l.emit(DSIGN)
		return lexAny
	case r == '@':
		l.emit(ATSIGN)
		return lexAny
	case r == '<':
		l.emit(LTHAN)
		return lexAny
	case r == '>':
		l.emit(GTHAN)
		return lexAny
	case isAlphaNumeric(r): // Already got numbers
		l.backup()
		return lexIdentifier
	default:
		return l.errorf("unrecognized character: %#U", r)
	}
}

// lexComment scans a comment. The comment marker has been consumed.
func lexComment(l *Scanner) stateFn {
	for {
		r := l.next()
		if r == eof || r == '\n' {
			break
		}
	}
	if len(l.input) > 0 {
		l.pos = len(l.input)
		l.start = l.pos - 1
		// Emitting newline also advances l.line.
		l.emit(NEWLINE)
	}
	return lexSpace
}

// lexSpace scans a run of space characters.
// One space has already been seen.
func lexSpace(l *Scanner) stateFn {
	for isSpace(l.peek()) {
		l.next()
	}
	l.ignore()
	return lexAny
}

// lexNumber scans a number. "05" == "5"
func lexNumber(l *Scanner) stateFn {
	l.acceptRun("0123456789")
	l.emit(INT)
	return lexAny
}

func lexIdentifier(l *Scanner) stateFn {
Loop:
	for {
		switch r := l.next(); {
		case isAlphaNumeric(r):
			// absorb.
		default:
			l.backup()
			word := l.input[l.start:l.pos]
			t := Lookup(word)
			l.emit(t.Type)
			break Loop
		}
	}
	return lexAny
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}
