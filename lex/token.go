// Package token defines constants representing the lexical tokens of Redcode
// and basic operations on tokens (printing, predicates).
//
package lex

import (
	"strconv"
	"strings"
)

// Token is the set of lexical tokens of Redcode
type Type int

type Token struct {
	Type Type
	Line int
	Text string
}

// The list of tokens.
const (
	// Special tokens
	EOF Type = iota
	ILLEGAL
	NEWLINE // LF | CR | LF CR | CR LF
	COMMENT // ; v* EOL | EOL

	literal_beg
	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	LABEL // alpha alphanumeral*
	INT   // int | -int | +int
	literal_end

	operator_beg
	// Operators and delimiters
	// Many have non-standard names to prevent conflicts with opcodes (PLUS/ADD, MINUS/SUB)
	PLUS      // 1+1
	MINUS     // 1-1
	ASTERISK  // 1*1
	FSLASH    // 1/1
	PERCENT   // 1%1
	LPAREN    // "("
	RPAREN    // ")"
	DOT       // "."
	COMMA     //","
	SEMICOLON //";"
	operator_end

	mode_beg
	// Addressing modes
	HASH   // "#"
	DSIGN  // "$"
	ATSIGN // "@"
	LTHAN  // "<"
	GTHAN  // ">"
	mode_end

	opcode_beg
	// Opcodes
	DAT
	MOV
	ADD
	SUB
	MUL
	DIV
	MOD
	JMP
	JMZ
	JMN
	DJN
	CMP
	SLT
	SPL
	ORG
	EQU
	END
	opcode_end

	mod_beg
	// Opcode modifiers
	A
	B
	AB
	BA
	F
	X
	I
	mod_end
)

var types = [...]string{
	EOF:     "EOF",
	ILLEGAL: "ILLEGAL",
	NEWLINE: "NEWLINE",
	COMMENT: "COMMENT",

	LABEL: "LABEL",
	INT:   "INT",

	PLUS:      "+",
	MINUS:     "-",
	ASTERISK:  "*",
	FSLASH:    "/",
	PERCENT:   "%",
	LPAREN:    "(",
	RPAREN:    ")",
	DOT:       ".",
	COMMA:     ",",
	SEMICOLON: ";",

	HASH:   "#",
	DSIGN:  "$",
	ATSIGN: "@",
	LTHAN:  "<",
	GTHAN:  ">",

	DAT: "DAT",
	MOV: "MOV",
	ADD: "ADD",
	SUB: "SUB",
	MUL: "MUL",
	DIV: "DIV",
	MOD: "MOD",
	JMP: "JMP",
	JMZ: "JMZ",
	JMN: "JMN",
	DJN: "DJN",
	CMP: "CMP",
	SLT: "SLT",
	SPL: "SPL",
	ORG: "ORG",
	EQU: "EQU",
	END: "END",

	A:  "A",
	B:  "B",
	AB: "AB",
	BA: "BA",
	F:  "F",
	X:  "X",
	I:  "I",
}

// String returns the string corresponding to the token tok.
// For operators, delimiters, and keywords the string is the actual
// token character sequence (e.g., for the token PLUS, the string is
// "+"). For all other tokens the string corresponds to the token
// constant name (e.g. for the token IDENT, the string is "IDENT").
//
func (tok Token) String() string {
	s := ""
	if 0 <= int(tok.Type) && int(tok.Type) < len(types) {
		s = types[tok.Type]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok.Type)) + ")"
	}
	return s
}

// A set of constants for precedence-based expression parsing.
// Non-operators have lowest precedence, followed by operators
// st:qarting with precedence 1 up to unary operators. The highest
// precedence serves as "catch-all" precedence for selector,
// indexing, and other operator and delimiter tokens.
//
const (
	LowestPrec  = 0
	UnaryPrec   = 2
	HighestPrec = 3
)

// Precedence returns the operator precedence of the binary
// operator op. If op is not a binary operator, the result
// is LowestPrecedence.
//
func (tok Token) Precedence() int {
	switch tok.Type {
	case PLUS, MINUS:
		return 1
	case ASTERISK, FSLASH, PERCENT: // 1*2, 3/4, 5%6
		return 2
	}
	return LowestPrec
}

var opcodes map[string]Type
var mods map[string]Type

func init() {
	opcodes = make(map[string]Type)
	for i := opcode_beg + 1; i < opcode_end; i++ {
		opcodes[types[i]] = i
	}
	mods = make(map[string]Type)
	for i := mod_beg + 1; i < mod_end; i++ {
		mods[types[i]] = i
	}
}

// Lookup maps an identifier to its keyword token or LABEL (if not a keyword).
//
func Lookup(ident string) Token {
	if t, is_opcode := opcodes[strings.ToUpper(ident)]; is_opcode {
		return Token{t, 0, ident}
	}
	if t, is_mod := mods[strings.ToUpper(ident)]; is_mod {
		return Token{t, 0, ident}
	}
	return Token{LABEL, 0, ident}
}

// Predicates

// IsLiteral returns true for tokens corresponding to identifiers
// and basic type literals; it returns false otherwise.
//
func (tok Token) IsLiteral() bool { return literal_beg < tok.Type && tok.Type < literal_end }

// IsOperator returns true for tokens corresponding to operators and
// delimiters; it returns false otherwise.
//
func (tok Token) IsOperator() bool { return operator_beg < tok.Type && tok.Type < operator_end }

// IsOpcode returns true for tokens corresponding to keywords;
// it returns false otherwise.
//
func (tok Token) IsOpcode() bool { return opcode_beg < tok.Type && tok.Type < opcode_end }
