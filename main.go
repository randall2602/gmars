package main

import (
	"fmt"
	"io/ioutil"

	"github.com/randall2602/gmars/core"
	"github.com/randall2602/gmars/loader"
)

func main() {

	config := core.ConfigureKOTH()
	config.CoreSize = 80 // for easier debugging
	config.MinSeparation = 8
	MyCore := core.Initialize(config)

	MyCore.View(8)

	impFile, _ := ioutil.ReadFile("imp.redcode")
	dwarfFile, _ := ioutil.ReadFile("dwarf.redcode")

	imp := loader.Compile(impFile)
	dwarf := loader.Compile(dwarfFile)
	warriors := []loader.Warrior{imp, dwarf}

	loader.Load(warriors, &MyCore, config)

	MyCore.View(8)

	fmt.Println(imp)
	fmt.Println(dwarf)
}
