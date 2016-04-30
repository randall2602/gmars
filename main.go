// Package main is the controller of the gMARS system. It uses the Core, Compiler, and
package main

import (
	"fmt"

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
}
