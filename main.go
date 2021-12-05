package main

import (
	"log"
	"os"
	"strconv"

	"github.com/vmorsell/aoc2021go/lib"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("missing date")
	}
	date, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("atoi: %v", err)
	}

	err = lib.GenerateDay(date)
	if err != nil {
		log.Fatal(err)
	}
}
