package lex

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestGmars(t *testing.T) {
	name := "../red/dwarf.red"
	file, err := os.Open(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gmars: %s\n", err)
		os.Exit(1)
	}
	scanner := New(name, bufio.NewReader(file))
	var tokens []Token
	expected := []string{
		"ADD", "#", "4", ",", "3", "\n",
		"MOV", "2", ",", "@", "2", "\n",
		"JMP", "-", "2", "\n",
		"DAT", "#", "0", ",", "#", "0",
	}
	i := 0
	for tok := scanner.Next(); tok.Type != EOF; tok = scanner.Next() {
		if len(expected) <= i {
			t.Errorf("Expected EOF, got %q", tok.Text)
			break
		}
		if tok.Text != expected[i] {
			t.Errorf("Expected %q, got %q", expected[i], tok.Text)
		}
		tokens = append(tokens, tok)
		i++
	}
	return
}
