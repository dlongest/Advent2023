package main

import (
	"fmt"
	"slices"
	"strings"
)

func Day9RunA() {
	fmt.Println("Running Day 9 Part A")

	seqs, err := LoadSequences("../Data/day9.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0

	for _, seq := range seqs {
		next := PredictNext(seq)

		sum += next
	}

	fmt.Printf("Predicted Value Sum = %v\n", sum)
}

func Day9RunB() {
	fmt.Println("Running Day 9 Part B")

	seqs, err := LoadSequences("../Data/day9.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0

	for _, seq := range seqs {
		next := PredictBefore(seq)

		sum += next
	}

	fmt.Printf("Predicted Value Sum = %v\n", sum)
}
func LoadSequences(filePath string) ([][]int, error) {

	lines := ReadLines(filePath)

	seq := [][]int{}

	for _, line := range lines {
		xs, err := ToIntSlice(strings.Split(line, " "))

		if err != nil {
			return nil, err
		}

		seq = append(seq, xs)
	}

	return seq, nil
}

func PredictBefore(xs []int) int {

	xs2 := make([]int, len(xs))
	copy(xs2, xs)

	slices.Reverse(xs2)

	return PredictNext(xs2)
}

func PredictNext(xs []int) int {

	s2 := [][]int{xs}

	s2 = GenerateAllSequences(s2)

	s2Count := len(s2)

	s2[s2Count-1] = append(s2[s2Count-1], 0)

	for i := s2Count - 2; i >= 0; i-- {
		currentRowSize := len(s2[i])

		lastElementCurrentRow := s2[i][currentRowSize-1]

		// We need the last row of the row below this one. We actually know it's at the index equal to size of the current row - 1
		lastElementRowAfter := s2[i+1][currentRowSize-1]

		s2[i] = append(s2[i], lastElementCurrentRow+lastElementRowAfter)
	}

	// Return last element of first row, which is the new predicted value for the xs sequence
	return s2[0][len(s2[0])-1]
}

func GenerateAllSequences(xs [][]int) [][]int {

	sequence := xs[0]

	for !IsSequenceComplete(sequence) {

		sequence = NextSequence(sequence)

		xs = append(xs, sequence)
	}

	return xs
}

func NextSequence(xs []int) []int {

	sequence := []int{}

	for i := 1; i < len(xs); i++ {

		sequence = append(sequence, xs[i]-xs[i-1])

	}

	return sequence
}

func IsSequenceComplete(xs []int) bool {

	for _, x := range xs {
		if x != 0 {
			return false
		}
	}

	return true
}
