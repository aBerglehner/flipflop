package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"

	"github.com/alex/flipflop/utils"
)

func main() {
	part3()
}

type Pos struct {
	x int
	y int
}

// 101x101
// range 0-100
// first x, second y
func part1() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "input")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	trashes := []Pos{{0, 0}}
	fmt.Printf("trashes: %v\n", trashes)
	for _, line := range data {
		fmt.Println(line)
		strs := strings.Split(line, ",")
		x, _ := strconv.Atoi(strs[0])
		y, _ := strconv.Atoi(strs[1])
		trashes = append(trashes, Pos{x, y})
	}
	fmt.Printf("trashes: %v\n", trashes)
	ans := 0
	for i := 1; i < len(trashes); i++ {
		prev := trashes[i-1]
		cur := trashes[i]
		ans += diff(prev.x, cur.x) + diff(prev.y, cur.y)
	}
	fmt.Printf("ans: %v\n", ans)
}

func diff(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func part2() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "input")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	trashes := []Pos{{0, 0}}
	fmt.Printf("trashes: %v\n", trashes)
	for _, line := range data {
		fmt.Println(line)
		strs := strings.Split(line, ",")
		x, _ := strconv.Atoi(strs[0])
		y, _ := strconv.Atoi(strs[1])
		trashes = append(trashes, Pos{x, y})
	}
	fmt.Printf("trashes: %v\n", trashes)
	ans := 0
	for i := 1; i < len(trashes); i++ {
		prev := trashes[i-1]
		cur := trashes[i]
		ans += max(diff(prev.x, cur.x), diff(prev.y, cur.y))
	}
	fmt.Printf("ans: %v\n", ans)
}

type PosX struct {
	x   int
	y   int
	dis int
}

// first sort in manhatten distance
// than get part2 distance
func part3() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "input")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	trashes := []PosX{{0, 0, 0}}
	fmt.Printf("trashes: %v\n", trashes)
	for _, line := range data {
		fmt.Println(line)
		strs := strings.Split(line, ",")
		x, _ := strconv.Atoi(strs[0])
		y, _ := strconv.Atoi(strs[1])
		trashes = append(trashes, PosX{x, y, 0})
	}

	for i := 1; i < len(trashes); i++ {
		prev := PosX{0, 0, 0}
		cur := trashes[i]
		dis := diff(prev.x, cur.x) + diff(prev.y, cur.y)
		trashes[i].dis = dis
	}
	fmt.Printf("trashes before: %v\n", trashes)
	slices.SortFunc(trashes, func(a, b PosX) int {
		return a.dis - b.dis
	})
	fmt.Printf("trashes after: %v\n", trashes)

	fmt.Printf("trashes: %v\n", trashes)
	ans := 0
	for i := 1; i < len(trashes); i++ {
		prev := trashes[i-1]
		cur := trashes[i]
		ans += max(diff(prev.x, cur.x), diff(prev.y, cur.y))
	}
	fmt.Printf("ans: %v\n", ans)
}
