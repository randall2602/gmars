package main

import (
	"fmt"

	"github.com/randall2602/gmars/core"
	"github.com/randall2602/gmars/loader"
	"github.com/randall2602/gmars/queues"
)

var EXECUTIVE_DESCRIPTION = `The Executive Function defines how redcode programs are executed in the Core.`

func WhatIsExecutive() string {
	return EXECUTIVE_DESCRIPTION
}

func main() {
	fmt.Println("This is gMARS; a go implementation of the CoreWars game intended to be hosted as a web service.")
	fmt.Println("A minimally-complete MARS consists of a core, a loader, task queues, the MARS executive function, and a way to present the final results of a battle.")
	fmt.Println(WhatIsExecutive())
	fmt.Println(core.WhatIsCore())
	fmt.Println(loader.WhatIsLoader())
	fmt.Println(queues.WhatIsQueues())
}
