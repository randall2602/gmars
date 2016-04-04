package main

import (
	"fmt"

	"github.com/randall2602/gmars/core"
)

func main() {

	config := core.ConfigureKOTH()
	MyCore := core.Core{
		Memory: make([]core.Instruction, config.CoreSize),
	}
	for i := range MyCore.Memory {
		MyCore.Memory[i] = config.InitialInstructions
	}
	fmt.Println(MyCore)
}
