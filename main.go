package main

import (
	"fmt"
	"log"
	"os"

	"lem-in/parse"
	"lem-in/types"
	xerr "lem-in/xerrors"
)

func main() {
	// read the arguments from the command line and ignore the name of the program at index 0
	arguments := os.Args[1:]

	len := len(arguments)

	if len != 1 {
		fmt.Println(xerr.ErrNoArgsPassed)
		os.Exit(1)
	}
	file := arguments[0]

	contents, err := parse.ReadValidateFileContents(file)
	if err != nil {
		log.Fatalf("%v\n", err)
		return
	}

	fmt.Printf("%q\n", contents)

	Colony, err := types.ParseFileContentsToColony(contents)
	if err != nil {
		log.Fatalf("%v\n", fmt.Errorf(err.Error(), file))
	}

	fmt.Printf("%+#v\n", Colony)
}
