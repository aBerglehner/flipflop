package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"slices"
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

	distances := []string{}
	for _, line := range data {
		fmt.Println(line)
		for v := range strings.SplitSeq(line, "") {
			distances = append(distances, v)
		}
	}
	fmt.Printf("distances: %v\n", distances)

	lookup := make(map[string][]int)
	for i, v := range distances {
		_, ok := lookup[v]
		if !ok {
			lookup[v] = []int{i}
		} else {
			lookup[v] = append(lookup[v], i)
		}
	}
	fmt.Printf("lookup: %v\n", lookup)
	idx := 0
	ans := 0
	for idx < len(distances) {
		cur := distances[idx]
		strs := lookup[cur]
		if strs[0] == idx {
			ans += dis(idx, strs[1])
			idx = strs[1] + 1
		} else {
			ans += dis(idx, strs[0])
			idx = strs[0] + 1
		}
	}
	fmt.Printf("ans: %v\n", ans)
}

func dis(a int, b int) int {
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

	distances := []string{}
	for _, line := range data {
		fmt.Println(line)
		for v := range strings.SplitSeq(line, "") {
			distances = append(distances, v)
		}
	}
	fmt.Printf("distances: %v\n", distances)

	lookup := make(map[string][]int)
	for i, v := range distances {
		_, ok := lookup[v]
		if !ok {
			lookup[v] = []int{i}
		} else {
			lookup[v] = append(lookup[v], i)
		}
	}
	fmt.Printf("lookup: %v\n", lookup)

	idx := 0
	ans := 0
	visited := make(map[string]bool)
	for idx < len(distances) {
		cur := distances[idx]
		visited[cur] = true
		strs := lookup[cur]
		if strs[0] == idx {
			ans += dis(idx, strs[1])
			idx = strs[1] + 1
		} else {
			ans += dis(idx, strs[0])
			idx = strs[0] + 1
		}
	}
	fmt.Printf("ans: %v\n", ans)

	res := []string{}
	for _, v := range distances {
		if !visited[v] {
			if !slices.Contains(res, v) {
				res = append(res, v)
			}
		}
	}
	for _, v := range res {
		fmt.Printf("%v", v)
	}
	fmt.Println()
}

func part3() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "input")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	distances := []string{}
	for _, line := range data {
		fmt.Println(line)
		for v := range strings.SplitSeq(line, "") {
			distances = append(distances, v)
		}
	}
	fmt.Printf("distances: %v\n", distances)

	lookup := make(map[string][]int)
	for i, v := range distances {
		_, ok := lookup[v]
		if !ok {
			lookup[v] = []int{i}
		} else {
			lookup[v] = append(lookup[v], i)
		}
	}
	fmt.Printf("lookup: %v\n", lookup)

	idx := 0
	ans := 0
	for idx < len(distances) {
		cur := distances[idx]
		strs := lookup[cur]
		if strs[0] == idx {
			if cur >= "A" && cur <= "Z" {
				ans -= dis(idx, strs[1])
			} else {
				ans += dis(idx, strs[1])
			}
			idx = strs[1] + 1
		} else {
			if cur >= "A" && cur <= "Z" {
				ans -= dis(idx, strs[0])
			} else {
				ans += dis(idx, strs[0])
			}
			idx = strs[0] + 1
		}
	}
	fmt.Printf("ans: %v\n", ans)
}
