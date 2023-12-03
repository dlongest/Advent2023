package main

import (
	"testing"
)

func TestSchematicParser(t *testing.T) {
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

	if len(symbols) != 1 {
		t.Error("Incorrect count of number matches")
	}

	if symbols["*"][0] != 9 && symbols["*"][1] != 9 {
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
