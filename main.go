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
	// Read the arguments from the command line
	arguments := os.Args[1:]
	if len(arguments) != 1 {
		fmt.Println(xerr.ErrNoArgsPassed)
		os.Exit(1)
	}
	file := arguments[0]

	// Read and validate file contents
	log.Println("Reading and validating file contents...")
	contents, err := parse.ReadValidateFileContents(file)
	if err != nil {
		log.Fatalf("Error reading file: %v\n", err)
		return
	}

	// Parse file contents into a Colony structure
	log.Println("Parsing file contents into colony...")
	Colony, err := types.ParseFileContentsToColony(contents)
	if err != nil {
		log.Fatalf("Error parsing file contents: %v\n", err)
		return
	}

	// Validate the Colony structure
	log.Println("Validating colony...")
	err = types.ValidateColony(Colony)
	if err != nil {
		log.Printf("Colony validation failed: %v\n", err)
		xerr.Logger(Colony, "err.txt")
		os.Exit(1)
	}

	// Log the Colony structure for debugging
	log.Println("Logging colony structure...")
	c, _ := json.MarshalIndent(Colony, "", "\t")
	xerr.Logger(c, "logger.json")
	fmt.Println("Colony structure logged to logger.json")

	// Start moving ants
	log.Println("Starting to move ants...")
	err = Colony.MoveAnts()
	if err != nil {
		log.Fatalf("Error during ant movement: %v\n", err)
		os.Exit(1)
	}

	log.Println("Ant movement completed successfully.")
}
