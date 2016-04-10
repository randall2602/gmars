package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/randall2602/gmars/core"
)

func Compile(f []byte) []string { // []Instruction {
	var i []string                          // var i []Instruction
	lines := strings.Split(string(f), "\n") // Returns [][]byte after splitting on /n
	for _, l := range lines {
		i = append(i, l) // i = append(i, parse(l)) // Returns Instruction
	}
	return l
}

func main() {

	config := core.ConfigureKOTH()
	config.CoreSize = 80 // for easier debugging
	config.MinSeparation = 8
	MyCore := core.Initialize(config)

	MyCore.View(8)

	impFileName := "imp.redcode"
	dwarfFileName := "dwarf.redcode"
	impFile, err := ioutil.ReadFile(impFileName)
	if err != nil {
		log.Fatalf("Couldn't read file %s", impFileName)
	}
	dwarfFile, err := ioutil.ReadFile(dwarfFileName)
	if err != nil {
		log.Fatalf("Couldn't read file %s", dwarfFileName)
	}

	imp := Compile(impFile)
	dwarf := Compile(dwarfFile)

	fmt.Println(imp)
	fmt.Println(dwarf)
}
