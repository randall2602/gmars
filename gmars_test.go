package main

import (
	"bufio"
	"fmt"
	"testing"
	"os"

	"github.com/randall2602/gmars/lex"
)

func TestGmars(t *testing.T) {
	name := "test1.red"
	file, err := os.Open(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gmars: %s\n", err)
		os.Exit(1)
	}
	scanner := lex.New(name, bufio.NewReader(file))
	var tokens []lex.Token
	for tok := scanner.Next(); tok.Type != lex.EOF; tok = scanner.Next() {
		tokens = append(tokens, tok)
	}
	if len(tokens) != 6 {
		t.Errorf("Incorrect number of tokens, expected 6, got %v", len(tokens))
	}
	
	for _, tok := range tokens {
		t.Logf("tok: %q, type: %v", tok.Text, tok)
	}
	fmt.Print("\n")
	return
}
