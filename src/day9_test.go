package main

import (
	"reflect"
	"testing"
)

func TestNextSequence(t *testing.T) {
	input := [][]int{{0, 3, 6, 9, 12, 15}, {3, 3, 3, 3, 3}, {1, 3, 6, 10, 15, 21}, {10, 13, 16, 21, 30, 45}}
	expected := [][]int{{3, 3, 3, 3, 3}, {0, 0, 0, 0}, {2, 3, 4, 5, 6}, {3, 3, 5, 9, 15}}

	for i := range input {

		got := NextSequence(input[i])

		if !reflect.DeepEqual(expected[i], got) {
			t.Errorf("Expected %v :: Got %v", expected[i], got)
			return
		}
	}
}

type PredictTestCase struct {
	xs       []int
	expected int
}

func TestPredictNext(t *testing.T) {

	tests := []PredictTestCase{{xs: []int{0, 3, 6, 9, 12, 15}, expected: 18},
		{xs: []int{1, 3, 6, 10, 15, 21}, expected: 28}, {xs: []int{10, 13, 16, 21, 30, 45}, expected: 68}}
	// 1 3 6 10 15 21
	// 10 13 16 21 30 45

	for _, tt := range tests {

		got := PredictNext(tt.xs)

		if tt.expected != got {
			t.Errorf("Expected %v :: Got %v", tt.expected, got)
		}
	}
}

func TestPredictBefore(t *testing.T) {
	tests := []PredictTestCase{{xs: []int{0, 3, 6, 9, 12, 15}, expected: -3},
		{xs: []int{1, 3, 6, 10, 15, 21}, expected: 0}, {xs: []int{10, 13, 16, 21, 30, 45}, expected: 5}}

	for _, tt := range tests {

		got := PredictBefore(tt.xs)

		if tt.expected != got {
			t.Errorf("Expected %v :: Got %v", tt.expected, got)
		}
	}
}
