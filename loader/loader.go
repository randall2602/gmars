// Package loader implements the Loader of the Memory Array Redcode Simulator (MARS).
package loader

import (
	"strings"

	"github.com/randall2602/gmars/core"
)

// A Warrior is an array of core instructions.
type Warrior struct {
	Body []core.Instruction
}

// Compile converts a redcode file into an array of instructions.
func Compile(d []byte) Warrior {
	// convert to array of strings
	code := strings.Split(string(d[0:len(d)-1]), "\n") //code is now an array of strings

	// parse each string into opcode, modifier, operand a, operand b and make into instruction
	w := Warrior{}
	for i := range code {
		iSet := strings.Split(code[i], " ")
		var ins core.Instruction
		for j := range iSet {
			if j == 0 {
				//fmt.Print("Opcode: " + iSet[j] + "\n")
				ins.Opcode = iSet[j]
			} else if (j == 1) && (len(iSet) == 2) {
				//fmt.Print("OperandA: " + iSet[j] + "\n")
				//fmt.Print("OperandB: N/A")
				ins.OperandA = iSet[j]
			} else if (j == 1) && (len(iSet) == 3) {
				//fmt.Print("OperandA: " + strings.Replace(iSet[j], ",", "", 1) + "\n")
				ins.OperandA = strings.Replace(iSet[j], ",", "", 1)
			} else if j == 2 {
				//fmt.Print("OperandB: " + iSet[j] + "\n")
				ins.OperandB = iSet[j]
			}
		}
		w.Body = append(w.Body, ins)
	}
	// make instructions into warrior

	return w
}

// Load inserts warriors into core per configuration.
func Load(w []Warrior, core *core.Core, config core.Config) {
	// get 2 random addresses, if less than config.MinSeparation then fix or try again
	rand := []int{15, 38}

	for war := range w {
		for ins := range w[war].Body {
			core.Memory[rand[war]+ins] = w[war].Body[ins]
		}
	}
}
