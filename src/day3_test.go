package main

import (
	"reflect"
	"testing"
)

func TestFullParse(t *testing.T) {

	// input := []string{"467..114.*", "...*@....."}

	// numbers := map[int]map[int][]int{0: map[int][]int{467: []int{0, 2}, 114: []int{5, 7}}}
	// row0Symbols := map[int][]int{0: []int{9}}
	// row1Symbols := map[int][]int{1: []int{3, 4}}

	//want := NewSchematicDetails(2, 10, numbers,

	// sut := NewSchematicParser()

	// got, err := sut.FullParse(input)

	// if err != nil {
	// 	t.Error("Full Parse failed: ", err)
	// }

}

func TestNewGrid(t *testing.T) {
	input := []string{"467..114.*", "...*@....."}

	got := NewGrid(input)

	if got.rowCount != 2 {
		t.Error("Row Count wrong")
	}

	if got.columnCount != 10 {
		t.Error("Column Count wrong")
	}

	if !reflect.DeepEqual(got.grid[0], []rune(input[0])) {
		t.Error("First row of the grid is wrong")
	}

}

func TestSingleLineParse(t *testing.T) {
	sut := NewSchematicParser()

	numbers, symbols, err := sut.Parse("467..114.*")

	if err != nil {
		t.Error("Failed parsing line")
	}

	if len(numbers) != 2 {
		t.Error("Incorrect count of number matches")
	}

	if numbers[467][0] != 0 && numbers[467][1] != 2 {
		t.Error("Target value 467 does not have correct indices")
	}

	if numbers[114][0] != 5 && numbers[114][1] != 7 {
		t.Error("Target value 467 does not have correct indices")
	}

	if len(symbols) != 1 {
		t.Error("Incorrect count of number matches")
	}

	if symbols["*"] != 9 {
		t.Error("Target symbol '*' does not have correct indices")
	}
}

type ToIntTestCase struct {
	digits   []int
	expected int
}

func TestToInt(t *testing.T) {

	tt := []ToIntTestCase{{digits: []int{}, expected: 0}, {digits: []int{2}, expected: 2},
		{digits: []int{4, 1}, expected: 41}, {digits: []int{2, 0, 1}, expected: 201}}

	for _, tc := range tt {

		got := ToInt(tc.digits)

		if got != tc.expected {
			t.Errorf("Expected %v :: Got %v", tc.expected, got)
		}
	}
}
