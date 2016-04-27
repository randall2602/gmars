//  Package main is the controller of the gMARS system. It uses the Core, Compiler, and
package main

import (
	"fmt"
)

type Warrior struct {
	body []byte
}

func (w Warrior) ReadByte() (byte, error) {
	var c byte = w.body
	var err error
	return c, err
}

func main() {
	fmt.Println("Hello world")
	w := Warrior{[]byte("abc def")}
	myScanner := New("my scanner", w)
	fmt.Println(myScanner)
	myScanner.loadLine()

}
