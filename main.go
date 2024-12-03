package main

import (
	"fmt"
	xerr "lem-in/xerrors"
	"os"
)

func main() {
	fmt.Println("let's lem-n")

	// read the arguments from the command line and ignore the name of the program at index 0
	arguments := os.Args[1:]

	len := len(arguments)

	if len != 1 {
		fmt.Println(xerr.ErrNoArgsPassed)
		os.Exit(1)
	}
	
}
