package main

import (
	"reflect"
	"testing"
)

type digitParserTests struct {
	input  string
	digits []int
}

func TestDigitParser(t *testing.T) {

	cases := []digitParserTests{
		{input: "two1nine", digits: []int{2, 1, 9}},
		{input: "eightwothree", digits: []int{8, 2, 3}},
		{input: "abcone2threexyz", digits: []int{1, 2, 3}},
		{input: "xtwone3four", digits: []int{2, 1, 3, 4}},
		{input: "4nineeightseven2", digits: []int{4, 9, 8, 7, 2}},
		{input: "zoneight234", digits: []int{1, 8, 2, 3, 4}}, // this is an interesting case because of the overlap in words
		{input: "7pqrstsixteen", digits: []int{7, 6}},
	}

	parser := NewDigitParser()

	for _, tc := range cases {
		got := parser.Parse(tc.input)

		if !reflect.DeepEqual(tc.digits, got) {
			t.Error("Failed on input string ", tc.input)
		}
	}
}
