package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/alex/flipflop/utils"
)

func main() {
	part3()
}

func part1() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "input")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	max := 0
	for _, line := range data {
		// fmt.Println(line)
		strs := strings.SplitSeq(line, "")
		height := 0
		for v := range strs {
			// fmt.Printf("v: %v\n", v)
			if v == "v" {
				height--
			} else {
				height++
			}
			if height > max {
				max = height
			}
		}
		fmt.Printf("height: %v\n", height)
	}
	fmt.Printf("max: %v\n", max)
}

func part2() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "input")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	max := 0
	for _, line := range data {
		// fmt.Println(line)
		strs := strings.SplitSeq(line, "")
		height := 0
		last := ""
		count := 1
		for v := range strs {
			if v == "v" {
				if last == "down" {
					count++
				} else {
					count = 1
					last = "down"
				}
				height -= count
			} else {
				if last == "up" {
					count++
				} else {
					count = 1
					last = "up"
				}
				height += count
			}
			if height > max {
				max = height
			}
			// fmt.Printf("v: %v\n", v)
			// fmt.Printf("height: %v\n", height)
			// fmt.Println()
		}
	}
	fmt.Printf("max: %v\n", max)
}

func part3() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "input")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	max := 0
	for _, line := range data {
		// fmt.Println(line)
		strs := strings.SplitSeq(line, "")
		height := 0
		last := ""
		count := 0
		for v := range strs {
			if v == "v" {
				if last == "down" {
					count++
				} else {
					fmt.Printf("adding ups count: %v\n", count)
					height += fib(count)
					count = 1
					last = "down"
				}
			} else {
				if last == "up" {
					count++
				} else {
					fmt.Printf("adding downs count: %v\n", count)
					height -= fib(count)
					count = 1
					last = "up"
				}
			}
			if height > max {
				max = height
			}
			fmt.Printf("v: %v\n", v)
			fmt.Printf("height: %v\n", height)
			fmt.Println()
		}
	}
	fmt.Printf("max: %v\n", max)
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	a := 0
	b := 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
