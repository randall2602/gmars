//  Package main is the controller of the gMARS system. It uses the Core, Compiler, and
package main

import (
        "fmt"
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
        t := Lookup("JMP")

        fmt.Println(t)
        fmt.Println("Is literal: ", t.IsLiteral())
        fmt.Println("Is opcode: ", t.IsOpcode())
        fmt.Println("Is operator: ", t.IsOperator())
        fmt.Println("Precedence: ", t.Precedence())
        fmt.Println("Token string: ", t.String())
}
