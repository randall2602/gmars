package display

import (
	"fmt"

	"github.com/randall2602/gmars/core"
)

func Print(c *core.Core, width int) {
	for i := 0; i < len(c.Memory)/width; i++ {
		fmt.Println(c.Memory[i*width : i*width+width])
	}
}
