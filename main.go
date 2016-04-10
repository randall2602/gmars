package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/randall2602/gmars/core"
	"github.com/randall2602/gmars/loader"
)

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

	imp := loader.Compile(impFile)
	dwarf := loader.Compile(dwarfFile)
	warriors := []loader.Warrior{imp, dwarf}

	loader.Load(warriors, &MyCore, config)

	MyCore.View(8)

	fmt.Println(imp)
	fmt.Println(dwarf)
}
