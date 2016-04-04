package main

import (
	"github.com/randall2602/gmars/core"
	"github.com/randall2602/gmars/display"
)

func main() {

	config := core.ConfigureKOTH()
	config.CoreSize = 80 // for easier debugging
	MyCore := core.Initialize(config)

	// print core n addresses at a time
	display.Print(&MyCore, 8)
}
