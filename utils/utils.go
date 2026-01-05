// Package utils
package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func PrintMatrix[T any](input []T, name ...string) {
	var matrixName string
	if len(name) != 0 {
		matrixName = name[0]
	}
	fmt.Println("---------------------------------")
	fmt.Println("matrix: ", matrixName)
	for _, row := range input {
		fmt.Println(row)
	}
	fmt.Println("---------------------------------")
}

type Pos struct {
	Row int
	Col int
}

// GetDis deprecated use GetDirs
func GetDis() []Pos {
	return []Pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
}

func GetDirs() []Pos {
	return []Pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
}

func GetDirsAlsoDiagonal() []Pos {
	return []Pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
}

func SplitTasks[T any](tasks []T, n int) [][]T {
	var chunks [][]T
	length := len(tasks)
	if n <= 0 {
		n = 1
	}
	for i := 0; i < n; i++ {
		start := i * length / n
		end := (i + 1) * length / n
		chunks = append(chunks, tasks[start:end])
	}
	return chunks
}
