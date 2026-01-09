package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/alex/flipflop/utils"
)

func main() {
	// part2()
	fmt.Println("")
	part3()
}

type bird struct {
	speedX int
	speedY int
	x      int
	y      int
}

func (b *bird) mov(upperLimit int) {
	xx := b.x + b.speedX
	if xx >= upperLimit {
		b.x = xx - upperLimit
	} else if xx < 0 {
		b.x = upperLimit + xx
	} else {
		b.x = xx
	}
	yy := b.y + b.speedY
	if yy >= upperLimit {
		b.y = yy - upperLimit
	} else if yy < 0 {
		b.y = upperLimit + yy
	} else {
		b.y = yy
	}
}

func part1() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "input")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	birds := []bird{}
	for _, line := range data {
		// fmt.Println(line)
		strs := strings.Split(line, ",")
		x, _ := strconv.Atoi(strs[0])
		y, _ := strconv.Atoi(strs[1])
		birds = append(birds, bird{x, y, 0, 0})
	}
	// fmt.Printf("birds: %v\n", birds)
	count := 0
	// TODO: 100
	for count < 100 {
		count++
		for i := range birds {
			// TODO: 1000
			birds[i].mov(1000)
		}
		// fmt.Printf("count: %v birds: %v\n", count, birds)
	}
	// fmt.Printf("count: %v birds: %v\n", count, birds)

	view := 500
	// included
	lower := 249
	// not included
	upper := lower + view

	ans := 0
	for _, bird := range birds {
		if bird.x >= lower && bird.x < upper {
			if bird.y >= lower && bird.y < upper {
				ans++
			}
		}
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

	birds := []bird{}
	for _, line := range data {
		// fmt.Println(line)
		strs := strings.Split(line, ",")
		x, _ := strconv.Atoi(strs[0])
		y, _ := strconv.Atoi(strs[1])
		birds = append(birds, bird{x, y, 0, 0})
	}
	// fmt.Printf("birds: %v\n", birds)
	ans := 0
	for range 1000 {
		ans += t(&birds, 3600)
	}
	fmt.Printf("ans: %v\n", ans)
}

func t(birds *[]bird, upperBound int) int {
	secCount := 0
	// TODO:
	for secCount < upperBound {
		secCount++
		for i := range *birds {
			// TODO: 1000
			(*birds)[i].mov(1000)
		}
		// fmt.Printf("count: %v birds: %v\n", count, birds)
	}
	// fmt.Printf("count: %v birds: %v\n", count, birds)

	view := 500
	// included
	lower := 249
	// not included
	upper := lower + view

	ans := 0
	for _, bird := range *birds {
		if bird.x >= lower && bird.x < upper {
			if bird.y >= lower && bird.y < upper {
				ans++
			}
		}
	}
	return ans
}

// check when and if they reset to 0
// 31_556_926
func part3() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	inputPath := filepath.Join(baseDir, "input")
	data, err := utils.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	birds := []bird{}
	for _, line := range data {
		// fmt.Println(line)
		strs := strings.Split(line, ",")
		x, _ := strconv.Atoi(strs[0])
		y, _ := strconv.Atoi(strs[1])
		birds = append(birds, bird{x, y, 0, 0})
	}
	// fmt.Printf("birds: %v\n", birds)
	ans := 0
	for range 1000 {
		cur := t2(&birds, 926)
		ans += cur
	}
	fmt.Printf("ans: %v\n", ans)
}

func t2(birds *[]bird, upperBound int) int {
	secCount := 0
	// TODO:
	for secCount < upperBound {
		// for secCount < upperBound {
		secCount++
		for i := range *birds {
			// TODO: 1000
			(*birds)[i].mov(1000)
		}
		// fmt.Printf("count: %v birds: %v\n", count, birds)
	}
	// fmt.Printf("count: %v birds: %v\n", count, birds)

	view := 500
	// included
	lower := 249
	// not included
	upper := lower + view

	ans := 0
	for _, bird := range *birds {
		if bird.x >= lower && bird.x < upper {
			if bird.y >= lower && bird.y < upper {
				ans++
			}
		}
	}
	return ans
}
