package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/alex/flipflop/utils"
)

func main() {
	part2()
	part2ParaCh()
	part2ParaMutex()
}

// 10 to 99
// red,green,blue
func part1() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "input")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)
	for _, line := range data {
		// fmt.Println(line)
		v := counts[line]
		counts[line] = v + 1
	}
	// fmt.Printf("counts: %v\n", counts)
	var highest string
	hnum := 0
	for k := range counts {
		if counts[k] > hnum {
			hnum = counts[k]
			highest = k
		}
	}
	fmt.Printf("highest: %v\n", highest)
	fmt.Printf("hnum: %v\n", hnum)
}

// red , green , blue
// Red if the amount of red (the first number) is greater than both green and blue.
// Green if the amount of green (the second number) is greater than red and greater than blue.
// Blue if the amount of blue (the third number) is greater than red and greater than green.
// Special if 2 or more colors are equal.
func part2() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "inputBig")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	ans := 0
	for _, line := range data {
		// fmt.Println(line)
		strs := strings.Split(line, ",")
		red, _ := strconv.Atoi(strs[0])
		green, _ := strconv.Atoi(strs[1])
		blue, _ := strconv.Atoi(strs[2])
		if red == green || red == blue || green == blue {
			continue
		}

		if green > red && green > blue {
			ans++
		}
	}
	// fmt.Printf("ans: %v\n", ans)
}

func part2ParaCh() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "inputBig")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	dataTasks := utils.SplitTasks(data, runtime.NumCPU())
	ch := make(chan int, len(dataTasks))
	var wg sync.WaitGroup
	for _, v := range dataTasks {
		wg.Add(1)
		go paraCompCh(v, ch, &wg)
	}
	wg.Wait()
	close(ch)

	ans := 0
	for v := range ch {
		ans += v
	}
	// fmt.Printf("ans: %v\n", ans)
}

func paraCompCh(data []string, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	ans := 0
	for i := range data {
		// fmt.Println(line)
		strs := strings.Split(data[i], ",")
		red, _ := strconv.Atoi(strs[0])
		green, _ := strconv.Atoi(strs[1])
		blue, _ := strconv.Atoi(strs[2])
		if red == green || red == blue || green == blue {
			continue
		}

		if green > red && green > blue {
			ans++
		}
	}
	ch <- ans
}

func part2ParaMutex() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "inputBig")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	dataTasks := utils.SplitTasks(data, runtime.NumCPU())
	var wg sync.WaitGroup
	var mu sync.Mutex
	ans := 0
	for _, v := range dataTasks {
		wg.Add(1)
		go paraCompMutex(v, &ans, &wg, &mu)
	}
	wg.Wait()

	// fmt.Printf("ans: %v\n", ans)
}

func paraCompMutex(data []string, globalAns *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	ans := 0
	for _, line := range data {
		// fmt.Println(line)
		strs := strings.Split(line, ",")
		red, _ := strconv.Atoi(strs[0])
		green, _ := strconv.Atoi(strs[1])
		blue, _ := strconv.Atoi(strs[2])
		if red == green || red == blue || green == blue {
			continue
		}

		if green > red && green > blue {
			ans++
		}
	}
	mu.Lock()
	*globalAns += ans
	mu.Unlock()
}

// red , green , blue
// Red if the amount of red (the first number) is greater than both green and blue.
// Green if the amount of green (the second number) is greater than red and greater than blue.
// Blue if the amount of blue (the third number) is greater than red and greater than green.
// Special if 2 or more colors are equal.
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
		// fmt.Println(line)
		strs := strings.Split(line, ",")
		red, _ := strconv.Atoi(strs[0])
		green, _ := strconv.Atoi(strs[1])
		blue, _ := strconv.Atoi(strs[2])
		// case special
		if red == green || red == blue || green == blue {
			ans += 10
			continue
		}
		// case red
		if red > green && red > blue {
			ans += 5
			continue
		}
		// case green
		if green > red && green > blue {
			ans += 2
			continue
		}
		// case blue
		if blue > red && blue > green {
			ans += 4
			continue
		}
		fmt.Println(line)

	}
	fmt.Printf("ans: %v\n", ans)
}
