package main

import (
	"fmt"
	"regexp"
	"slices"
	"unicode"
)

type Day1 struct {
	Day
}

func (d Day1) RunA() {
	fmt.Println("Running Day 1 A")

	data := ReadLines("../Data/day1.txt")

	sum := 0
	for _, s := range data {
		sum += FindCalibrationValue(s)
	}

	fmt.Println("Sum = ", sum)
}

func FindCalibrationValue(calibrationString string) (calibrationValue int) {
	first := FindFirstNumberIn(calibrationString)
	last := FindLastNumberIn(calibrationString)

	return ComputeCalibrationValue(first, last)
}

func ComputeCalibrationValue(first, last int) int {
	return first*10 + last
}

func FindFirstNumberIn(s string) (value int) {
	for _, c := range s {
		if unicode.IsDigit(c) {
			return int(c) - '0'
		}
	}

	panic("Couldn't find any digits in the calibration string")
}

func FindLastNumberIn(s string) (value int) {
	for i := range s {
		adjustedIndex := len(s) - i - 1
		c := rune(s[adjustedIndex])
		if unicode.IsDigit(c) {
			return int(c) - '0'
		}
	}

	panic("Couldn't find any digits in the calibration string")
}

func (d Day1) RunB() {
	fmt.Println("Running Day 1 B")

	data := ReadLines("../Data/day1.txt")

	parser := NewDigitParser()
	sum := 0
	for _, s := range data {
		sum += ParseCalibrationValue(parser, s)
	}

	fmt.Println("Sum = ", sum)
}

func ParseCalibrationValue(parser *DigitParser, s string) int {
	allDigits := parser.Parse(s)
	first, last := selectCalibrationDigits(allDigits)
	return ComputeCalibrationValue(first, last)
}

func selectCalibrationDigits(digits []int) (int, int) {
	first := digits[0]

	if len(digits) == 1 {
		return first, first
	}

	return first, digits[len(digits)-1]
}

type DigitParser struct {
	expressions map[int]*regexp.Regexp
}

func (p *DigitParser) Expressions() map[int]*regexp.Regexp {
	return p.expressions
}

func (p *DigitParser) SetExpression(number int, regex *regexp.Regexp) {
	p.expressions[number] = regex
}

func NewDigitParser() *DigitParser {

	p := &DigitParser{}

	p.expressions = make(map[int]*regexp.Regexp)

	p.expressions[0] = regexp.MustCompile("0|zero")
	p.expressions[1] = regexp.MustCompile("1|one")
	p.expressions[2] = regexp.MustCompile("2|two")
	p.expressions[3] = regexp.MustCompile("3|three")
	p.expressions[4] = regexp.MustCompile("4|four")
	p.expressions[5] = regexp.MustCompile("5|five")
	p.expressions[6] = regexp.MustCompile("6|six")
	p.expressions[7] = regexp.MustCompile("7|seven")
	p.expressions[8] = regexp.MustCompile("8|eight")
	p.expressions[9] = regexp.MustCompile("9|nine")

	return p
}

func (p *DigitParser) Parse(s string) []int {

	matchedValues := make(map[int]int) // Key = Index, Value = parsed number

	// Loop through each regular expression looking for all matches in the parameter s.
	// Each match is stored in matchedValues with key of the start index of the match and
	// value equal to the digit matching the regex being used
	for digit, regex := range p.Expressions() {

		matches := regex.FindAllStringIndex(s, len(s))

		// After getting all the matches for the current regex and value, add just the
		// start index for each match to the values map.
		for _, matchedIndexTuple := range matches {
			matchedValues[matchedIndexTuple[0]] = digit
		}
	}

	// matchedValues now contains a map in the form { index: digit }
	// We know each index can only appear one time since they're found in a slice.
	// Each digit can show up several times, mapped to separate indexes
	indexes := make([]int, 0, len(matchedValues))

	for i := range matchedValues {
		indexes = append(indexes, i)
	}

	slices.Sort(indexes)

	digitsInOrder := make([]int, 0, len(matchedValues))

	for _, i := range indexes {
		digitsInOrder = append(digitsInOrder, matchedValues[i])
	}

	return digitsInOrder
}
