// package implements the Core of the Memory Array Redcode Simulator (MARS)
package core

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

// Initialize with KOTH settings. http://www.koth.org/
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
