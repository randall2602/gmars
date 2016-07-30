//Generated TestNew
//Generated TestScanner_Next
package lex

import (
	//"bufio"
	"io"
	//"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	//file, _ := os.Open("../red/dwarf.red")
	//reader := bufio.NewReader(file)
	type args struct {
		name string
		r    io.ByteReader
	}
	tests := []struct {
		name string
		args args
		want *Scanner
	}{
	/*
		{
			name: "Test New()",
			args: args{
				name: "My Scanner",
				r:    reader,
			},
			want: &Scanner{
				r:      reader,
				name:   "My Scanner",
				line:   1,
				tokens: make(chan Token, 2),
				state:  lexAny,
			},
		},
	*/
	}
	for _, tt := range tests {
		if got := New(tt.args.name, tt.args.r); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. New() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestScanner_Next(t *testing.T) {
	type fields struct {
		tokens     chan Token
		r          io.ByteReader
		done       bool
		name       string
		buf        []byte
		input      string
		leftDelim  string
		rightDelim string
		state      stateFn
		line       int
		pos        int
		start      int
		width      int
	}
	tests := []struct {
		name   string
		fields fields
		want   Token
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		l := &Scanner{
			tokens:     tt.fields.tokens,
			r:          tt.fields.r,
			done:       tt.fields.done,
			name:       tt.fields.name,
			buf:        tt.fields.buf,
			input:      tt.fields.input,
			leftDelim:  tt.fields.leftDelim,
			rightDelim: tt.fields.rightDelim,
			state:      tt.fields.state,
			line:       tt.fields.line,
			pos:        tt.fields.pos,
			start:      tt.fields.start,
			width:      tt.fields.width,
		}
		if got := l.Next(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Scanner.Next() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
