package lib

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type dayParams struct {
	BaseName string
}

func GenerateDay(date int) error {
	params := dayParams{
		BaseName: baseName(date),
	}

	_, err := os.Stat(params.BaseName)
	if !os.IsNotExist(err) {
		return fmt.Errorf("day already exists")
	}
	if err = os.Mkdir(params.BaseName, 0755); err != nil {
		return fmt.Errorf("mkdir: %v", err)
	}

	f, err := os.Create(fmt.Sprintf("%s/%s.go", params.BaseName, params.BaseName))
	if err != nil {
		log.Fatalf("create: %v", err)
	}
	defer f.Close()

	tpl := template.Must(template.New("tpl").Parse(dayTpl))
	if err := tpl.Execute(f, params); err != nil {
		log.Fatalf("execute: %v", err)
	}

	// Touch some files
	filenames := []string{
		fmt.Sprintf("%s.in", params.BaseName),
		fmt.Sprintf("%s.test", params.BaseName),
	}
	for _, name := range filenames {
		f, err := os.Create(fmt.Sprintf("%s/%s", params.BaseName, name))
		if err != nil {
			log.Fatalf("touch %s: %v", name, err)
		}
		f.Close()
	}
	return nil
}

func baseName(date int) string {
	return fmt.Sprintf("%02d", date)
}

var dayTpl = `package main

//go:generate go run ../lib/genrunners {{ .BaseName }}

import (
	"io"
)

var (
	wantPart1Test interface{}
	wantPart2Test interface{}
)

func parse(r io.Reader) (interface{}, error) {
	return nil, nil
}

func part1(input interface{}) (interface{}, error) {
	return nil, nil
}

func part2(input interface{}) (interface{}, error) {
	return nil, nil
}
`
