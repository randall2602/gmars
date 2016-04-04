package main

import (
	"io/ioutil"

	"github.com/randall2602/gmars/core"
	"github.com/randall2602/gmars/display"
	"github.com/randall2602/gmars/loader"
)

func main() {

	config := core.ConfigureKOTH()
	config.CoreSize = 80 // for easier debugging
	MyCore := core.Initialize(config)

	// print core n addresses at a time
	display.Print(&MyCore, 8)

	imp := loader.Compile(ioutil.ReadFile("imp.redcode"))
	dwarf := loader.Compile(ioutil.ReadFile("dwarf.redcode"))
	warriors := []loader.Warrior{imp, dwarf}

	loader.Load(warriors, &MyCore, config)

	display.Print(&MyCore, 8)
}
