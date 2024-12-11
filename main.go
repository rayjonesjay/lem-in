package main

import (
	"encoding/json"
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

	//fmt.Printf("%q\n", contents)
	// contains the struct in json format
	logFile := "logger.json"
	l := log.Logger{}
	fd, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	defer fd.Close()
	l.SetOutput(fd)

	Colony, err := types.ParseFileContentsToColony(contents)
	if err != nil {
		log.Fatalf("%v\n", fmt.Errorf(err.Error(), file))
	}

	c, _ := json.MarshalIndent(Colony, "", "\t")
	l.Println(string(c))
	fmt.Printf("look for a file named %s\n", logFile)
}
