package main

import (
	"log"
	"os"

	"lemin/controllers"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("ERROR: lem-in expects one argument\nExpecting: go run . [name of file]")
		return
	}
	p := controllers.NewParser()

	c, err := p.ParseFile(os.Args[1])
	if err != nil {
		log.Println(err)
		return
	}

	controllers.InitializeAnts(c)

	controllers.Mover(c)
}
