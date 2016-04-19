//  Package main is the controller of the gMARS system. It uses the Core, Compiler, and 
package main

import (
	"fmt"
)

type Warrior struct {}

func (w *Warrior) unload() (byte, error) {
  var c byte
  var err error
  return c, err
}

func main() {
  fmt.Println("Hello world")
  w := Warrior{}
  myScanner := New("my scanner", w)
}