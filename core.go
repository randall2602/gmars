// package implements the Core of the Memory Array Redcode Simulator (MARS)
package core

import "fmt"

var coreDescription string = 'The core (the memory of the simulated computer) is a continuous
 array of instructions, empty except for the competing programs. The core wraps around, so
 that after the last instruction comes the first one again.'

// Returns a description of what the MARS Core is
func WhatIsCore() string {
  return coreDescription
}
