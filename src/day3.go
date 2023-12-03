package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func Day3RunA() {
	fmt.Println("Starting Day 3 Part A")

}

func Day3RunB() {
	fmt.Println("Starting Day 3 Part B")
}

type SchematicParser struct {
	numberRegex *regexp.Regexp
	symbolRegex *regexp.Regexp
}

func NewSchematicParser() *SchematicParser {
	p := &SchematicParser{}
	p.numberRegex = regexp.MustCompile("[0-9]+")

	symbolPattern := fmt.Sprintf("[%v]", regexp.QuoteMeta("!@#$%^&*"))

	p.symbolRegex = regexp.MustCompile(symbolPattern)

	return p
}

func (p *SchematicParser) FullParse(schematic []string) error {

	// numbers holds a mapping of row positions in the schematic to a map of the number values to corresponding indices
	// Given line 0 with value 467 starting at index 1, then numbers[0] = map[int][]int with entry [467] = [0, 2]
	rowToNumbers := map[int]map[int][]int{}

	rowToSymbols := map[int]map[string][]int{}

	for i, line := range schematic {
		numbers, symbols, err := p.Parse(line)

		if err != nil {
			return err
		}

		rowToNumbers[i] = numbers
		rowToSymbols[i] = symbols
	}

	return nil
}

func (p *SchematicParser) Parse(line string) (map[int][]int, map[string][]int, error) {

	numberMatches := p.numberRegex.FindAllString(line, len(line))
	numberIndices := p.numberRegex.FindAllStringIndex(line, len(line))
	symbolMatches := p.symbolRegex.FindAllString(line, len(line))
	symbolIndices := p.symbolRegex.FindAllStringIndex(line, len(line))

	if len(numberMatches) != len(numberIndices) || len(symbolMatches) != len(symbolIndices) {
		return nil, nil, errors.New("The number of regexp matches doesn't match the number of index matches and it has to by definition.")
	}

	numbers, err := convertNumberMatchesToMap(numberMatches, numberIndices)

	if err != nil {
		return nil, nil, err
	}

	symbols, err := convertSymbolMatchesToMap(symbolMatches, symbolIndices)

	if err != nil {
		return nil, nil, err
	}

	return numbers, symbols, nil
}

func convertNumberMatchesToMap(numberMatches []string, numberIndices [][]int) (map[int][]int, error) {
	numbers := map[int][]int{}

	for i := range numberMatches {

		number, err := strconv.Atoi(numberMatches[i])

		if err != nil {
			return nil, err
		}

		adjustedIndices := []int{numberIndices[i][0], numberIndices[i][1] - 1}

		numbers[number] = adjustedIndices
	}

	return numbers, nil
}

func convertSymbolMatchesToMap(symbolMatches []string, symbolIndices [][]int) (map[string][]int, error) {
	symbols := map[string][]int{}

	for i := range symbolMatches {

		adjustedIndices := []int{symbolIndices[i][0], symbolIndices[i][1] - 1}
		symbols[symbolMatches[i]] = adjustedIndices
	}

	return symbols, nil
}

// ToInt takes a BigEndian array of digits and returns the corresponding number.
// Given the array [2, 0, 1] where 2 is at index 0, ToInt returns 201.
func ToInt(digits []int) int {

	tensFactor := int(math.Pow(10, float64(len(digits)-1)))
	number := 0

	for _, digit := range digits {
		number += digit * tensFactor
		tensFactor /= 10
	}

	return number
}
