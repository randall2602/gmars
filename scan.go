package main

import "strconv"

type Token int

const (
	ILLEGAL Token = iota
	EOF
	NEWLINE
	COMMENT

	literal_beg
	LABEL
	INT
	literal_end

	operator_beg
	PLUS
	MINUS
	ASTERISK
	FSLASH
	PERCENT
	LPAREN
	RPAREN
	DOT
	COMMA
	SEMICOLON
	operator_end

	opcode_beg
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
	A
	B
	AB
	BA
	F
	X
	I
	mod_end

	mode_beg
	HASH
	DSIGN
	ATSIGN
	LTHAN
	GTHAN
	mode_end
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

	HASH:   "HASH",
	DSIGN:  "DSIGN",
	ATSIGN: "ATSIGN",
	LTHAN:  "LTHAN",
	GTHAN:  "GTHAN",
}

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
