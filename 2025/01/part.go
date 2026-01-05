package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/alex/flipflop/utils"
)

func main() {
	part3()
	part3Para()
	part3ParaMutex()
}

func part1() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "input")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	ans := 0
	for _, line := range data {
		a := strings.Count(line, "ba")
		b := strings.Count(line, "na")
		c := strings.Count(line, "ne")
		ans += a + b + c

		// fmt.Println(line)
		// fmt.Printf("a: %v\n", a)
		// fmt.Printf("b: %v\n", b)
		// fmt.Println()
	}
	fmt.Printf("ans: %v\n", ans)
}

func part2() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "input")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	ans := 0
	for _, line := range data {
		a := strings.Count(line, "ba")
		b := strings.Count(line, "na")
		c := strings.Count(line, "ne")
		sum := a + b + c
		if sum%2 == 0 {
			ans += sum
		}

		// fmt.Println(line)
		// fmt.Printf("a: %v\n", a)
		// fmt.Printf("b: %v\n", b)
		// fmt.Println()
	}
	fmt.Printf("ans: %v\n", ans)
}

func part3() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "input")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	ans := 0
	for _, line := range data {
		a := strings.Count(line, "ba")
		b := strings.Count(line, "na")
		c := strings.Count(line, "ne")
		sum := a + b + c
		if c == 0 {
			ans += sum
		}

		// fmt.Println(line)
		// fmt.Printf("a: %v\n", a)
		// fmt.Printf("b: %v\n", b)
		// fmt.Println()
	}
	// fmt.Printf("ans: %v\n", ans)
}

func part3Para() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "input")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	splits := runtime.NumCPU()
	tasks := utils.SplitTasks(data, splits)
	var wg sync.WaitGroup
	ch := make(chan int, len(tasks))
	for _, data := range tasks {
		wg.Add(1)
		go taskPara(data, ch, &wg)
	}

	wg.Wait()
	close(ch)

	ans := 0
	for v := range ch {
		ans += v
	}
	// fmt.Printf("ans: %v\n", ans)
}

func taskPara(data []string, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	ans := 0
	for _, line := range data {
		a := strings.Count(line, "ba")
		b := strings.Count(line, "na")
		c := strings.Count(line, "ne")
		sum := a + b + c
		if c == 0 {
			ans += sum
		}
	}
	ch <- ans
}

func part3ParaMutex() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "input")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	splits := runtime.NumCPU()
	tasks := utils.SplitTasks(data, splits)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var ans int
	for _, data := range tasks {
		wg.Add(1)
		go taskMutex(data, &mu, &ans, &wg)
	}

	wg.Wait()

	// fmt.Printf("ans: %v\n", ans)
}

func taskMutex(data []string, mu *sync.Mutex, sum *int, wg *sync.WaitGroup) {
	defer wg.Done()

	ans := 0
	for _, line := range data {
		a := strings.Count(line, "ba")
		b := strings.Count(line, "na")
		c := strings.Count(line, "ne")
		sum := a + b + c
		if c == 0 {
			ans += sum
		}
	}
	mu.Lock()
	defer mu.Unlock()
	*sum += ans
}
