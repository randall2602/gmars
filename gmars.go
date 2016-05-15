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
	name := "test3.red"
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
		fmt.Print("tok: ", tok, " ")
		if tok.Type == token.NEWLINE {
			fmt.Print("\n")
		}
	}
	fmt.Print("\n")
}
