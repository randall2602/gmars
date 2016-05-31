// Package main is the controller of the gMARS system. It uses the Core, Compiler, and
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/randall2602/gmars/lex"
)

func main() {
	name := "test3.red"
	file, err := os.Open(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gmars: %s\n", err)
		os.Exit(1)
	}
	scanner := lex.New(name, bufio.NewReader(file))

	for tok := scanner.Next(); tok.Type != lex.EOF; tok = scanner.Next() {
		switch tok.Type {
		case lex.NEWLINE:
			fmt.Print(tok, " ")
			fmt.Print("\n")
		case lex.LABEL, lex.INT:
			fmt.Print(tok.Text, " ")
		default:
			fmt.Print(tok, " ")
		}
	}
	fmt.Print(lex.Token{lex.EOF,0,""})
	fmt.Print("\n")
}
