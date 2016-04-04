package main

import (
	"fmt"

	"github.com/randall2602/gmars/core"
)

func main() {

	KOTHConfig := core.Config{
		CoreSize:        8000,
		CyclesBeforeTie: 80000,
		InitialInstructions: core.Instruction{
			Opcode:   "DAT",
			Modifier: "F",
			OperandA: "$0",
			OperandB: "$0",
		},
		InstructionLimit: 100,
		MaxTasks:         8000,
		MinSeparation:    100,
		ReadDistance:     8000,
		Separation:       0,
		Warriors:         2,
		WriteDistance:    8000,
	}

	MyCore := core.Core{
		Memory: make([]core.Instruction, 8000),
	}
	fmt.Println(KOTHConfig.CoreSize)
	fmt.Println(MyCore)
}
