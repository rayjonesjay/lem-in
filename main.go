package main

import (
	"fmt"
	"lemin/controllers"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Expecting: go run . [name of file]")
		return
	}
	p := controllers.NewParser()

	c, err := p.ParseFile(os.Args[1])
	fmt.Println("Parsing file...")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c)
	fmt.Println("Success")

	fmt.Println("Initializing ants to paths..")
	controllers.InitializeAnts(c)
	fmt.Println("Initialization successful")

	fmt.Println("Starting the movement...let them in ")

	mover := controllers.NewMover(c)
	movements := mover.ExecuteMovements()

	fmt.Println(movements)
}
