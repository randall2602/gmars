package main

import (
	"fmt"

	"github.com/randall2602/gmars/core"
)

func main() {

	config := core.ConfigureKOTH()
	MyCore := core.Initialize(config)
	fmt.Println(MyCore)
}
