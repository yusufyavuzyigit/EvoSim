package main

import (
	"log"

	core "github.com/yusufyavuzyigit/EvoSim/core"
)

func main() {
	log.SetPrefix("EvoSim> ")
	log.SetFlags(0)

	err := core.Configure()

	if err != nil {
		log.Fatal(err)
	}

	err = core.Run()

	if err != nil {
		log.Fatal(err)
	}
}