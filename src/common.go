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
