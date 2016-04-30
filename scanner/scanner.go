// Package scanner implements a scanner for redcode source text. It takes a
// []byte as a source which can then be tokenized through repeated calls to the
// Scan method
//
package scanner

import "github.com/randall2602/gmars/token"

// An ErrorHandler may be provided to Scanner.Init. If a syntax error is
// encountered and a handler was installed, the handler is called with a
// position and an error message. The position points to the beginning of the
// offending token.
//
type ErrorHandler func(pos token.Position, msg string)

<<<<<<< HEAD
// A Scanner holds the scanner's internal state while processing
// a given text.  It can be allocated as part of another data
// structure but must be initialized via Init before use.
//
type Scanner struct{}
=======
type Mode int

type Scanner struct {
	file *token.File
	dir  string
	src  []byte
	err  ErrorHandler
	mode Mode
>>>>>>> f4aa608d940419d15db7ee320c29f2cfc2d187e4

const bom = 0xFEFF

// Read the next Unicode char into s.ch.
// s.ch < 0 means end-of-file.
//
func (s *Scanner) next() {}

// A mode value is a set of flags (or 0).
// They control scanner behavior.
//
type Mode uint

const (
	ScanComments Mode = 1 << iota
	dontInsertSemis
)

// Init prepares the scanner s to tokenize the text src by setting the
// scanner at the beginning of src. The scanner uses the file set file
// for position information and it adds line information for each line.
// It is ok to re-use the same file when re-scanning the same file as
// line information which is already present is ignored. Init causes a
// panic if the file size does not match the src size.
//
// Calls to Scan will invoke the error handler err if they encounter a
// syntax error and err is not nil. Also, for each error encountered,
// the Scanner field ErrorCount is incremented by one. The mode parameter
// determines how comments are handled.
//
// Note that Init may call err if there is an error in the first character
// of the file.
//
func (s *Scanner) Init(file *token.File, src []byte, err ErrorHandler, mode Mode) {}

func (s *Scanner) error(offs int, msg string) {}

var prefix []byte

func (s *Scanner) interpretLineComment(text []byte) {}

func (s *Scanner) scanComment() string {}

func (s *Scanner) findLineEnd() bool {}

func isLetter(ch rune) bool {}

func isDigit(ch rune) bool {}

func (s *Scanner) scanIdentifier() string {}

func digitVal(ch rune) int {}

func (s *Scanner) scanMantissa(base int) {}

func (s *Scanner) scanNumber(seenDecimalPoint bool) (token.Token, string) {}

// scanEscape parses an escape sequence where rune is the accepted
// escaped quote. In case of a syntax error, it stops at the offending
// character (without consuming it) and returns false. Otherwise
// it returns true.
func (s *Scanner) scanEscape(quote rune) bool {}

func (s *Scanner) scanRune() string {}

func (s *Scanner) scanString() string {}

func stripCR(b []byte) []byte {}

func (s *Scanner) scanRawString() string {}

func (s *Scanner) skipWhitespace() {}

// Helper functions for scanning multi-byte tokens such as >> += >>= .
// Different routines recognize different length tok_i based on matches
// of ch_i. If a token ends in '=', the result is tok1 or tok3
// respectively. Otherwise, the result is tok0 if there was no other
// matching character, or tok2 if the matching character was ch2.

func (s *Scanner) switch2(tok0, tok1 token.Token) token.Token {}

func (s *Scanner) switch3(tok0, tok1 token.Token, ch2 rune, tok2 token.Token) token.Token {}

func (s *Scanner) switch4(tok0, tok1 token.Token, ch2 rune, tok2, tok3 token.Token) token.Token {}

// Scan scans the next token and returns the token position, the token,
// and its literal string if applicable. The source end is indicated by
// token.EOF.
//
// If the returned token is a literal (token.IDENT, token.INT, token.FLOAT,
// token.IMAG, token.CHAR, token.STRING) or token.COMMENT, the literal string
// has the corresponding value.
//
// If the returned token is a keyword, the literal string is the keyword.
//
// If the returned token is token.SEMICOLON, the corresponding
// literal string is ";" if the semicolon was present in the source,
// and "\n" if the semicolon was inserted because of a newline or
// at EOF.
//
// If the returned token is token.ILLEGAL, the literal string is the
// offending character.
//
// In all other cases, Scan returns an empty literal string.
//
// For more tolerant parsing, Scan will return a valid token if
// possible even if a syntax error was encountered. Thus, even
// if the resulting token sequence contains no illegal tokens,
// a client may not assume that no error occurred. Instead it
// must check the scanner's ErrorCount or the number of calls
// of the error handler, if there was one installed.
//
// Scan adds line information to the file added to the file
// set with Init. Token positions are relative to that file
// and thus relative to the file set.
//
func (s *Scanner) Scan() (pos token.Pos, tok token.Token, lit string) {}
