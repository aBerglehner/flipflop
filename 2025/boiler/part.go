package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/alex/flipflop/utils"
)

func main() {
	part1()
}

func part1() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "test")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range data {
		fmt.Println(line)
	}
}
