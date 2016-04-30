package scanner

import "github.com/randall2602/gmars/token"

type ErrorHandler func(pos token.Position, msg string)

type Mode int

type Scanner struct {
	file *token.File
	dir  string
	src  []byte
	err  ErrorHandler
	mode Mode

	ch         rune
	offset     int
	rdOffset   int
	lineOffset int

	ErrorCount int
}
