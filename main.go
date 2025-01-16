package main

import (
	"fmt"
	"log"
	"os"

	"lemin/controllers"
)

func main() {
	fmt.Println("let them in ")
	// read the file name from the command line
	arguments := os.Args[1:]
	if len(arguments) != 1 {
		log.Println("ERROR: lem-in expects one argument")
	}

	// call the parser function to read the file data
	controllers.NewParser().ParseFile(arguments[0])
	
}
