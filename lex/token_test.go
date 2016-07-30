//Generated TestToken_String
//Generated TestToken_Precedence
//Generated TestLookup
//Generated TestToken_IsLiteral
//Generated TestToken_IsOperator
//Generated TestToken_IsOpcode
package lex

import (
	"reflect"
	"testing"
)

func TestToken_String(t *testing.T) {
	type fields struct {
		Type Type
		Line int
		Text string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ILLEGAL Token",
			fields: fields{
				Type: ILLEGAL,
				Line: 1,
				Text: "!",
			},
			want: "ILLEGAL",
		}, {
			name: "NEWLINE Token",
			fields: fields{
				Type: NEWLINE,
				Line: 1,
				Text: "\n",
			},
			want: "NEWLINE",
		}, {
			name: "COMMENT Token",
			fields: fields{
				Type: COMMENT,
				Line: 1,
				Text: "This is a comment.",
			},
			want: "COMMENT",
		}, {
			name: "LABEL Token",
			fields: fields{
				Type: LABEL,
				Line: 1,
				Text: "MYLABEL1",
			},
			want: "LABEL",
		}, {
			name: "INT Token",
			fields: fields{
				Type: INT,
				Line: 1,
				Text: "7",
			},
			want: "INT",
		}, {
			name: "PLUS Token",
			fields: fields{
				Type: PLUS,
				Line: 1,
				Text: "+",
			},
			want: "+",
		}, {
			name: "MINUS Token",
			fields: fields{
				Type: MINUS,
				Line: 1,
				Text: "-",
			},
			want: "-",
		}, {
			name: "ASTERISK Token",
			fields: fields{
				Type: ASTERISK,
				Line: 1,
				Text: "*",
			},
			want: "*",
		}, {
			name: "FSLASH Token",
			fields: fields{
				Type: FSLASH,
				Line: 1,
				Text: "/",
			},
			want: "/",
		}, {
			name: "PERCENT Token",
			fields: fields{
				Type: PERCENT,
				Line: 1,
				Text: "%",
			},
			want: "%",
		}, {
			name: "LPAREN Token",
			fields: fields{
				Type: LPAREN,
				Line: 1,
				Text: "(",
			},
			want: "(",
		}, {
			name: "RPAREN Token",
			fields: fields{
				Type: RPAREN,
				Line: 1,
				Text: ")",
			},
			want: ")",
		}, {
			name: "DOT Token",
			fields: fields{
				Type: DOT,
				Line: 1,
				Text: ".",
			},
			want: ".",
		}, {
			name: "COMMA Token",
			fields: fields{
				Type: COMMA,
				Line: 1,
				Text: ",",
			},
			want: ",",
		}, {
			name: "SEMICOLON Token",
			fields: fields{
				Type: SEMICOLON,
				Line: 1,
				Text: ";",
			},
			want: ";",
		}, {
			name: "HASH Token",
			fields: fields{
				Type: HASH,
				Line: 1,
				Text: "#",
			},
			want: "#",
		}, {
			name: "DSIGN Token",
			fields: fields{
				Type: DSIGN,
				Line: 1,
				Text: "$",
			},
			want: "$",
		}, {
			name: "ATSIGN Token",
			fields: fields{
				Type: ATSIGN,
				Line: 1,
				Text: "@",
			},
			want: "@",
		}, {
			name: "LTHAN Token",
			fields: fields{
				Type: LTHAN,
				Line: 1,
				Text: "<",
			},
			want: "<",
		}, {
			name: "GTHAN Token",
			fields: fields{
				Type: GTHAN,
				Line: 1,
				Text: ">",
			},
			want: ">",
		}, {
			name: "DAT Token",
			fields: fields{
				Type: DAT,
				Line: 1,
				Text: "DAT",
			},
			want: "DAT",
		}, {
			name: "MOV Token",
			fields: fields{
				Type: MOV,
				Line: 1,
				Text: "MOV",
			},
			want: "MOV",
		}, {
			name: "ADD Token",
			fields: fields{
				Type: ADD,
				Line: 1,
				Text: "ADD",
			},
			want: "ADD",
		}, {
			name: "SUB Token",
			fields: fields{
				Type: SUB,
				Line: 1,
				Text: "SUB",
			},
			want: "SUB",
		}, {
			name: "MUL Token",
			fields: fields{
				Type: MUL,
				Line: 1,
				Text: "MUL",
			},
			want: "MUL",
		}, {
			name: "DIV Token",
			fields: fields{
				Type: DIV,
				Line: 1,
				Text: "DIV",
			},
			want: "DIV",
		}, {
			name: "MOD Token",
			fields: fields{
				Type: MOD,
				Line: 1,
				Text: "MOD",
			},
			want: "MOD",
		}, {
			name: "JMP Token",
			fields: fields{
				Type: JMP,
				Line: 1,
				Text: "JMP",
			},
			want: "JMP",
		}, {
			name: "JMZ Token",
			fields: fields{
				Type: JMZ,
				Line: 1,
				Text: "JMZ",
			},
			want: "JMZ",
		}, {
			name: "JMN Token",
			fields: fields{
				Type: JMN,
				Line: 1,
				Text: "JMN",
			},
			want: "JMN",
		}, {
			name: "DJN Token",
			fields: fields{
				Type: DJN,
				Line: 1,
				Text: "DJN",
			},
			want: "DJN",
		}, {
			name: "CMP Token",
			fields: fields{
				Type: CMP,
				Line: 1,
				Text: "CMP",
			},
			want: "CMP",
		}, {
			name: "SLT Token",
			fields: fields{
				Type: SLT,
				Line: 1,
				Text: "SLT",
			},
			want: "SLT",
		}, {
			name: "SPL Token",
			fields: fields{
				Type: SPL,
				Line: 1,
				Text: "SPL",
			},
			want: "SPL",
		}, {
			name: "ORG Token",
			fields: fields{
				Type: ORG,
				Line: 1,
				Text: "ORG",
			},
			want: "ORG",
		}, {
			name: "EQU Token",
			fields: fields{
				Type: EQU,
				Line: 1,
				Text: "EQU",
			},
			want: "EQU",
		}, {
			name: "END Token",
			fields: fields{
				Type: END,
				Line: 1,
				Text: "END",
			},
			want: "END",
		}, {
			name: "A Token",
			fields: fields{
				Type: A,
				Line: 1,
				Text: "A",
			},
			want: "A",
		}, {
			name: "B Token",
			fields: fields{
				Type: B,
				Line: 1,
				Text: "B",
			},
			want: "B",
		}, {
			name: "AB Token",
			fields: fields{
				Type: AB,
				Line: 1,
				Text: "AB",
			},
			want: "AB",
		}, {
			name: "BA Token",
			fields: fields{
				Type: BA,
				Line: 1,
				Text: "BA",
			},
			want: "BA",
		}, {
			name: "F Token",
			fields: fields{
				Type: F,
				Line: 1,
				Text: "F",
			},
			want: "F",
		}, {
			name: "X Token",
			fields: fields{
				Type: X,
				Line: 1,
				Text: "X",
			},
			want: "X",
		}, {
			name: "I Token",
			fields: fields{
				Type: I,
				Line: 1,
				Text: "I",
			},
			want: "I",
		},
	}
	for _, tt := range tests {
		tok := Token{
			Type: tt.fields.Type,
			Line: tt.fields.Line,
			Text: tt.fields.Text,
		}
		if got := tok.String(); got != tt.want {
			t.Errorf("%q. Token.String() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestToken_Precedence(t *testing.T) {
	type fields struct {
		Type Type
		Line int
		Text string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "PLUS Operator Precedence",
			fields: fields{
				Type: PLUS,
				Line: 1,
				Text: "+",
			},
			want: 1,
		}, {
			name: "MINUS Operator Precedence",
			fields: fields{
				Type: MINUS,
				Line: 1,
				Text: "-",
			},
			want: 1,
		}, {
			name: "ASTERISK Operator Precedence",
			fields: fields{
				Type: ASTERISK,
				Line: 1,
				Text: "*",
			},
			want: 2,
		}, {
			name: "FSLASH Operator Precedence",
			fields: fields{
				Type: FSLASH,
				Line: 1,
				Text: "/",
			},
			want: 2,
		}, {
			name: "PERCENT Operator Precedence",
			fields: fields{
				Type: PERCENT,
				Line: 1,
				Text: "%",
			},
			want: 2,
		}, {
			name: "LPAREN Operator Precedence",
			fields: fields{
				Type: LPAREN,
				Line: 1,
				Text: "(",
			},
			want: 0,
		}, {
			name: "RPAREN Operator Precedence",
			fields: fields{
				Type: RPAREN,
				Line: 1,
				Text: ")",
			},
			want: 0,
		}, {
			name: "DOT Operator Precedence",
			fields: fields{
				Type: DOT,
				Line: 1,
				Text: ".",
			},
			want: 0,
		}, {
			name: "COMMA Operator Precedence",
			fields: fields{
				Type: COMMA,
				Line: 1,
				Text: ",",
			},
			want: 0,
		}, {
			name: "SEMICOLON Operator Precedence",
			fields: fields{
				Type: SEMICOLON,
				Line: 1,
				Text: ";",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		tok := Token{
			Type: tt.fields.Type,
			Line: tt.fields.Line,
			Text: tt.fields.Text,
		}
		if got := tok.Precedence(); got != tt.want {
			t.Errorf("%q. Token.Precedence() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestLookup(t *testing.T) {
	type args struct {
		ident string
	}
	tests := []struct {
		name string
		args args
		want Token
	}{
		{
			name: "Lookup LABEL",
			args: args{
				ident: "LABEL1",
			},
			want: Token{LABEL, 0, "LABEL1"},
		}, {
			name: "Lookup PLUS",
			args: args{
				ident: "+",
			},
			want: Token{PLUS, 0, "+"},
		}, {
			name: "Lookup MINUS",
			args: args{
				ident: "-",
			},
			want: Token{MINUS, 0, "-"},
		}, {
			name: "Lookup ASTERISK",
			args: args{
				ident: "*",
			},
			want: Token{ASTERISK, 0, "*"},
		}, {
			name: "Lookup FSLASH",
			args: args{
				ident: "/",
			},
			want: Token{FSLASH, 0, "/"},
		}, {
			name: "Lookup PERCENT",
			args: args{
				ident: "%",
			},
			want: Token{PERCENT, 0, "%"},
		}, {
			name: "Lookup LPAREN",
			args: args{
				ident: "(",
			},
			want: Token{LPAREN, 0, "("},
		}, {
			name: "Lookup RPAREN",
			args: args{
				ident: ")",
			},
			want: Token{RPAREN, 0, ")"},
		}, {
			name: "Lookup DOT",
			args: args{
				ident: ".",
			},
			want: Token{DOT, 0, "."},
		}, {
			name: "Lookup COMMA",
			args: args{
				ident: ",",
			},
			want: Token{COMMA, 0, ","},
		}, {
			name: "Lookup SEMICOLON",
			args: args{
				ident: ";",
			},
			want: Token{SEMICOLON, 0, ";"},
		},
	}
	for _, tt := range tests {
		if got := Lookup(tt.args.ident); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Lookup() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestToken_IsLiteral(t *testing.T) {
	type fields struct {
		Type Type
		Line int
		Text string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tok := Token{
			Type: tt.fields.Type,
			Line: tt.fields.Line,
			Text: tt.fields.Text,
		}
		if got := tok.IsLiteral(); got != tt.want {
			t.Errorf("%q. Token.IsLiteral() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestToken_IsOperator(t *testing.T) {
	type fields struct {
		Type Type
		Line int
		Text string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tok := Token{
			Type: tt.fields.Type,
			Line: tt.fields.Line,
			Text: tt.fields.Text,
		}
		if got := tok.IsOperator(); got != tt.want {
			t.Errorf("%q. Token.IsOperator() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestToken_IsOpcode(t *testing.T) {
	type fields struct {
		Type Type
		Line int
		Text string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tok := Token{
			Type: tt.fields.Type,
			Line: tt.fields.Line,
			Text: tt.fields.Text,
		}
		if got := tok.IsOpcode(); got != tt.want {
			t.Errorf("%q. Token.IsOpcode() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
