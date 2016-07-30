// Package main is the controller of the gMARS system. It uses the Core, Compiler, and
package main

import (
	"bufio"
	"fmt"
	"os"
	//"github.com/randall2602/gmars/core"
	//"github.com/randall2602/gmars/exec"
	"github.com/randall2602/gmars/lex"
	//"github.com/randall2602/gmars/load"
	//"github.com/randall2602/gmars/parse"
)

func check(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "gmars: %s\n", e)
		os.Exit(1)
	}
}

func main() {
	names := [...]string{"./red/imp.red", "./red/dwarf.red"}
	files := [len(names)]*os.File{}
	for i, v := range names {
		var err error
		files[i], err = os.Open(v)
		check(err)
	}

	// double check that files loaded correctly
	for i, _ := range files {
		buf := make([]byte, 4096)
		n, err := files[i].Read(buf)
		check(err)
		fmt.Printf("%d bytes: \n%s\n\n", n, string(buf))
		files[i].Seek(0, 0)
	}
	fmt.Println()

	lexed_files := [len(names)][]lex.Token{}
	for i, _ := range files {
		scanner := lex.New(names[i], bufio.NewReader(files[i]))
		for tok := scanner.Next(); tok.Type != lex.EOF; tok = scanner.Next() {
			lexed_files[i] = append(lexed_files[i], tok)
		}
	}

}
