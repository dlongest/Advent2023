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

type PartNumberEvaluator struct {
	parser *SchematicParser
}

type SchematicDetails struct {
	firstRow, firstColumn int
	rowCount, columnCount int
	rowToNumbersToIndices map[int]map[int][]int // Key: row index to a map of number (key) to column indices (value)
	rowToSymbolIndices    map[int]int           // We only need the row it's found on (key) and single column index (value)
}

func NewSchematicDetails(rowCount int, columnCount int,
	rowToNumbersToIndices map[int]map[int][]int,
	rowToSymbolIndices map[int]int) SchematicDetails {
	return SchematicDetails{firstRow: 0, firstColumn: 0, rowCount: rowCount, columnCount: columnCount,
		rowToNumbersToIndices: rowToNumbersToIndices, rowToSymbolIndices: rowToSymbolIndices}
}

type RowColumnIndex struct {
	rowIndex, columnIndex int
}

func NewRowColumnIndex(rowIndex int, columnIndex int) RowColumnIndex {
	return RowColumnIndex{rowIndex: rowIndex, columnIndex: columnIndex}
}

type RowColumnIndexPair struct {
	start, end RowColumnIndex
}

func NewRowColumnIndexPair(start RowColumnIndex, end RowColumnIndex) RowColumnIndexPair {
	return RowColumnIndexPair{start: start, end: end}
}

func NewPartNumberEvaluator() PartNumberEvaluator {
	return PartNumberEvaluator{parser: NewSchematicParser()}
}

type Grid struct {
	rowCount, columnCount int
	grid                  [][]rune
}

func NewGrid(lines []string) *Grid {

	grid := &Grid{rowCount: len(lines), columnCount: len(lines[0])}

	for _, row := range lines {
		grid.grid = append(grid.grid, []rune(row))
	}

	return grid
}

func (e PartNumberEvaluator) Evaluate(schematic []string) (int, error) {

	//details := NewSchematicDetails(len(schematic[0]), len(schematic))

	// rowsToNumber, rowsToSymbols, err := e.parser.FullParse(schematic)

	// if err != nil {
	// 	return -1, err
	// }

	// for rowIndex := 0; rowIndex < rowCount; rowIndex++ {

	// 	row := rowsToNumber[rowIndex]
	// 	for number, indices := range row {

	// 	}
	// }
	return 0, nil
}

func generateSymbolIndicesToCheck(currentRowIndex int, details SchematicDetails, indices []int) []RowColumnIndex {
	// For a given currentRowIndex and indices within that row, we need to generate all the indices that form
	// a box around the indices. For internal indices (0 < currentRowIndex < rowCount and 0 < anything in indices < columnCount),
	// we can generate the indices to check by:
	// / * Generate [currentRowIndex-1, indices[0]-1] through [currentRowIndex-1, indices[1]+1]
	// / * Generate [currentRowIndex, indices[0]-1] and [currentRowIndex, indices[1]+1]
	// / * Generate [currentRowIndex+1, indices[0]-1] through [currentRowIndex+1, indices[1]+1]
	// /
	// / For indices on the edge (so currentRowIndex = 0 or rowCount-1 and similar for columns), we have to simply
	// / prune them appropriately: skip the first or last row, or don't step up/down/left/right depending

	return nil
}

func NewSchematicParser() *SchematicParser {
	p := &SchematicParser{}
	p.numberRegex = regexp.MustCompile("[0-9]+")

	symbolPattern := fmt.Sprintf("[%v]", regexp.QuoteMeta("!@#$%^&*"))

	p.symbolRegex = regexp.MustCompile(symbolPattern)

	return p
}

func (p *SchematicParser) FullParse(schematic []string) (SchematicDetails, error) {

	// numbers holds a mapping of row positions in the schematic to a map of the number values to corresponding indices
	// Given line 0 with value 467 starting at index 1, then numbers[0] = map[int][]int with entry [467] = [1, 3]
	rowToNumbers := map[int]map[int][]int{}

	rowToSymbols := map[int]map[string]int{}

	for i, line := range schematic {
		numbers, symbols, err := p.Parse(line)

		if err != nil {
			return SchematicDetails{}, err
		}

		rowToNumbers[i] = numbers
		rowToSymbols[i] = symbols
	}

	//return NewSchematicDetails(len(schematic[0], len(schematic), rowToNumbers,

	return SchematicDetails{}, nil
}

func (p *SchematicParser) Parse(line string) (map[int][]int, map[string]int, error) {

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

// convertNumberMatchesToMap takes a collection of number matches (where each number is of varable length) and
// corresponding collection of indexes identifying the first position and 1 past the last position of the number match
// and returns a map of each number to the index pair for the first position and last position of the match.
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

// convertSymbolMatchesToMap takes a collection of symbol matches and a corresponding collection of index pairs
// for each match and returns a map of each symbol to the single index where it exists.
// CRAP THIS WON'T WORK BECAUSE A GIVEN SYMBOL CAN SHOW UP MULTIPLE TIMES IN THE ROW.
// THE EASIEST SOLUTION IS WE DON'T ACTUALLY NEED THE SYMBOLS AT ALL AT THIS STAGE. WE JUST NEED
// THE INDICES. SO MAYBE THAT'S THE BEST BET.
func convertSymbolMatchesToMap(symbolMatches []string, symbolIndices [][]int) (map[string]int, error) {
	symbols := map[string]int{}

	for i := range symbolMatches {

		if symbolIndices[i][1]-symbolIndices[i][0] > 1 {
			return nil, errors.New("Matched symbol contains more than 1 character, which is unexpected.")
		}

		symbols[symbolMatches[i]] = symbolIndices[i][0]
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
