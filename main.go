package main

import (
	"fmt"
	"log"
	"os"

	"lemin/controllers"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Expecting: go run . [name of file]")
		return
	}
	p := controllers.NewParser()
	fmt.Println("Parsing file...")
	c, err := p.ParseFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(c)
	fmt.Println("Success")

	fmt.Println("Initializing ants to paths..")
	controllers.InitializeAnts(c)
	fmt.Println("Initialization successful")

	fmt.Println("Starting the movement...let them in ")

	controllers.Mover(c)

	fmt.Println("Movements successful")
	fmt.Println("let them in ")
	// read the file name from the command line
	arguments := os.Args[1:]
	if len(arguments) != 1 {
		log.Println("ERROR: lem-in expects one argument")
	}

	// call the parser function to read the file data
	controllers.NewParser().ParseFile(arguments[0])
}
