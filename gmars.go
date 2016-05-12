// Package main is the controller of the gMARS system. It uses the Core, Compiler, and
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/randall2602/gmars/scanner"
	"github.com/randall2602/gmars/token"
)

func main() {
	t := token.Lookup("JMP")

	fmt.Println(t)
	fmt.Println("Is literal: ", t.IsLiteral())
	fmt.Println("Is opcode: ", t.IsOpcode())
	fmt.Println("Is operator: ", t.IsOperator())
	fmt.Println("Precedence: ", t.Precedence())
	fmt.Println("Token string: ", t.String())
	
	name := "test.redcode"
	fd, err := os.Open(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gmars: %s\n", err)
		os.Exit(1)
	}
	scanner := scanner.New(name, bufio.NewReader(fd))
	fmt.Println("scanner:", scanner)
	fmt.Println()
	fmt.Println("Tokens:")
	for tok := scanner.Next(); tok.Type != token.EOF; tok = scanner.Next() {
		fmt.Println("tok: ", tok)
	}
}
