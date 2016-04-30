//  Package main is the controller of the gMARS system. It uses the Core, Compiler, and
package main

import (
<<<<<<< HEAD
	"fmt"

	"github.com/randall2602/gmars/token"
=======
        "fmt"
>>>>>>> f4aa608d940419d15db7ee320c29f2cfc2d187e4
)

type Instruction struct {
        txt string
}

func (i Instruction) IsLiteral() bool {
        return false
}

func (i Instruction) IsOpcode() bool {
        return false
}

func (i Instruction) IsOperator() bool {
        return false
}

func (t Instruction) Precedence() int {
        return 0
}

func (t Instruction) String() string {
        return t.txt
}

func Lookup(i string) Instruction {
        return Instruction{i}
}

func main() {
<<<<<<< HEAD
	t := token.Lookup("JMP")

	fmt.Println(t)
	fmt.Println("Is literal: ", t.IsLiteral())
	fmt.Println("Is opcode: ", t.IsOpcode())
	fmt.Println("Is operator: ", t.IsOperator())
	fmt.Println("Precedence: ", t.Precedence())
	fmt.Println("Token string: ", t.String())
=======
        t := Lookup("JMP")

        fmt.Println(t)
        fmt.Println("Is literal: ", t.IsLiteral())
        fmt.Println("Is opcode: ", t.IsOpcode())
        fmt.Println("Is operator: ", t.IsOperator())
        fmt.Println("Precedence: ", t.Precedence())
        fmt.Println("Token string: ", t.String())
>>>>>>> f4aa608d940419d15db7ee320c29f2cfc2d187e4
}
