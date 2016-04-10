// Package core implements the Core of the Memory Array Redcode Simulator (MARS)
//
// Draft design:
//     New(config) -> Core
//     Load(warrior...) -> position
//     Reset([InitialInstruction]) -> err
package core

import "fmt"

type Instruction struct {
	// An opcode is any of the following: DAT, MOV, ADD, SUB, MUL, DIV, MOD, JMP, JMZ, JMN, DJN, CMP, SLT, or SPL.
	Opcode string

	// An opcode may be followed by a modifier.  A modifier always begins with a dot.  A modifier is any of the
	// following: .A, .B, .AB, .BA, .F, .X, or .I.
	Modifier string

	// Each operand is blank, contains an address, or contains an addressing mode indicator and an address.
	OperandA string
	OperandB string
}

type Config struct {
	CoreSize            int
	CyclesBeforeTie     int
	InitialInstructions Instruction
	InstructionLimit    int
	MaxTasks            int
	MinSeparation       int
	ReadDistance        int
	Separation          int // if zero, Separation is random
	Warriors            int
	WriteDistance       int
}

type Core struct {
	Memory []Instruction
}

// View prints out the Core w columns at a time
func (c *Core) View(w int) {
	for i := 0; i < len(c.Memory)/w; i++ {
		fmt.Println(c.Memory[i*w : i*w+w])
	}
}

func ConfigureKOTH() Config {
	return Config{
		CoreSize:        8000,
		CyclesBeforeTie: 80000,
		InitialInstructions: Instruction{
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
}

func Initialize(config Config) Core {
	core := Core{Memory: make([]Instruction, config.CoreSize)}
	for i := range core.Memory {
		core.Memory[i] = config.InitialInstructions
	}
	return core
}
