package main

import (
	"bufio"
	"fmt"
	"os"
)

type Day interface {
	RunA()
	RunB()
}

func ReadLines(filePath string) []string {
	lines := []string{}

	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	readFile.Close()

	return lines
}

func Minimum(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func Maximum(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func IsIntegral(val float64) bool {
	return val == float64(int(val))
}

func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func iabs(a int) int {

	if a < 0 {
		return -a
	}

	return a
}

func imax(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func imin(a, b int) int {
	if a > b {
		return b
	}

	return a
}
