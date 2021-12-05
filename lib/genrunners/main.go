package main

import (
	"log"
	"os"

	"github.com/vmorsell/aoc2021go/lib"
)

func main() {
	if len(os.Args) > 2 {
		log.Fatal("missing day")
	}
	day := os.Args[1]
	if err := lib.GenerateRunners(day, "."); err != nil {
		log.Fatalf("generate: %v", err)
	}
}
