package main

import (
	"fmt"
	"log"
	"math/big"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/alex/flipflop/utils"
)

func main() {
	part3()
	part3Para()
}

type pos struct {
	r int
	c int
}

// 2*r / c -> total paths
func part1() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "input")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	distances := []pos{}
	for _, line := range data {
		fmt.Println(line)
		strs := strings.Split(line, " ")
		x, _ := strconv.Atoi(strs[0])
		y, _ := strconv.Atoi(strs[1])
		distances = append(distances, pos{x, y})
	}
	fmt.Printf("distances: %+v\n", distances)
	fmt.Println()

	ans := 0
	for _, v := range distances {
		fmt.Printf("v: %v\n", v)
		cur := findPaths(v)
		fmt.Printf("cur: %v\n", cur)
		ans += cur
		fmt.Println()
	}
	fmt.Printf("ans: %v\n", ans)
}

func findPaths(p pos) int {
	// r+c-2! / (r-1)! * (c-1)!
	upper := fact(p.r + p.c - 2)
	lowerLeft := fact(p.r - 1)
	lowerRight := fact(p.c - 1)
	// fmt.Printf("upper: %v\n", upper)
	// fmt.Printf("lowerLeft: %v\n", lowerLeft)
	// fmt.Printf("lowerRight: %v\n", lowerRight)

	return upper / (lowerLeft * lowerRight)
}

func fact(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

type dim struct {
	r int
	c int
	y int
}

func part2() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "input")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	distances := []dim{}
	for _, line := range data {
		fmt.Println(line)
		strs := strings.Split(line, " ")
		x, _ := strconv.Atoi(strs[0])
		y, _ := strconv.Atoi(strs[1])
		distances = append(distances, dim{x, y, x})
	}
	fmt.Printf("distances: %+v\n", distances)
	fmt.Println()

	ans := 0
	for _, v := range distances {
		fmt.Printf("v: %v\n", v)
		cur := findPaths3d(v)
		fmt.Printf("cur: %v\n", cur)
		ans += cur
		fmt.Println()
	}
	fmt.Printf("ans: %v\n", ans)
}

func findPaths3d(p dim) int {
	// r+c-2! / (r-1)! * (c-1)!
	upper := factB(p.r + p.c + p.y - 3)
	lowerLeft := factB(p.r - 1)
	lowerRight := factB(p.c - 1)
	lower3d := factB(p.y - 1)
	fmt.Printf("upper: %v\n", upper)
	fmt.Printf("lowerLeft: %v\n", lowerLeft)
	fmt.Printf("lowerRight: %v\n", lowerRight)
	fmt.Printf("lower3d: %v\n", lower3d)
	temp := new(big.Int).Mul(lowerLeft, lowerRight)
	res := new(big.Int).Mul(temp, lower3d)
	r := new(big.Int).Div(upper, res)

	return int(r.Int64())
}

func factB(n int) *big.Int {
	result := big.NewInt(1)
	for i := int64(2); i <= int64(n); i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}

func part3() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "input")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	distances := [][]int64{}
	for _, line := range data {
		// fmt.Println(line)
		strs := strings.Split(line, " ")
		n, _ := strconv.Atoi(strs[0])
		l, _ := strconv.Atoi(strs[1])
		dim := []int64{}
		for range n {
			dim = append(dim, int64(l))
		}
		distances = append(distances, dim)
	}
	// fmt.Printf("distances: %+v\n", distances)
	// fmt.Println()

	ans := 0
	for _, v := range distances {
		cur := multinomialBig(v)
		ans += cur
		// fmt.Printf("v: %v\n", v)
		// fmt.Printf("cur: %v\n", cur)
		// fmt.Println()
	}
	// fmt.Printf("ans: %v\n", ans)
}

func multinomialBig(lengths []int64) int {
	moves := make([]int64, len(lengths))
	total := int64(0)
	for i, L := range lengths {
		moves[i] = L - 1
		total += moves[i]
	}

	result := big.NewInt(1)
	used := int64(0)
	for _, m := range moves {
		for i := int64(0); i < m; i++ {
			result.Mul(result, big.NewInt(total-used-i))
			result.Div(result, big.NewInt(i+1))
		}
		used += m
	}
	return int(result.Int64())
}

func part3Para() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "input")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	distances := [][]int64{}
	for _, line := range data {
		// fmt.Println(line)
		strs := strings.Split(line, " ")
		n, _ := strconv.Atoi(strs[0])
		l, _ := strconv.Atoi(strs[1])
		dim := []int64{}
		for range n {
			dim = append(dim, int64(l))
		}
		distances = append(distances, dim)
	}
	// fmt.Printf("distances: %+v\n", distances)
	// fmt.Println()

	var mu sync.Mutex
	var wg sync.WaitGroup

	ans := 0
	for _, v := range distances {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			cur := multinomialBig(v)
			ans += cur
		}()
		// fmt.Printf("v: %v\n", v)
		// fmt.Printf("cur: %v\n", cur)
		// fmt.Println()
	}
	wg.Wait()
	// fmt.Printf("ans: %v\n", ans)
}
