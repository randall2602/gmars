// Package token defines constants representing the lexical tokens of Redcode
// and basic operations on tokens (printing, predicates).
//
package token

import "strconv"

// Token is the set of lexical tokens of Redcode
type Token int

// The list of tokens.
const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
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

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
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
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
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

func (op Token) Precedence() int {
	switch op {
	case PLUS, MINUS:
		return 1
	case ASTERISK, FSLASH, PERCENT: // 1*2, 3/4, 5%6
		return 2
	}
	return LowestPrec
}

var opcodes map[string]Token

func init() {
	opcodes = make(map[string]Token)
	for i := opcode_beg + 1; i < opcode_end; i++ {
		opcodes[tokens[i]] = i
	}
}

func Lookup(ident string) Token {
	if tok, is_opcode := opcodes[ident]; is_opcode {
		return tok
	}
	return LABEL
}

func (tok Token) IsLiteral() bool { return literal_beg < tok && tok < literal_end }

func (tok Token) IsOperator() bool { return operator_beg < tok && tok < operator_end }

func (tok Token) IsOpcode() bool { return opcode_beg < tok && tok < opcode_end }
