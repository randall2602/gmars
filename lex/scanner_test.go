package lex

import (
	"bufio"
	"fmt"
	"os"
	"testing"

)

func TestGmars(t *testing.T) {
	name := "../test1.red"
	file, err := os.Open(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gmars: %s\n", err)
		os.Exit(1)
	}
	scanner := lex.New(name, bufio.NewReader(file))
	var tokens []lex.Token
	expected := []string{"1", "\n", "a", "\n", "+", "\n"}
	i := 0
	for tok := scanner.Next(); tok.Type != lex.EOF; tok = scanner.Next() {
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
