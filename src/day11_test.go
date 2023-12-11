package main

import (
	"reflect"
	"testing"
)

func TestFindEmptyRows(t *testing.T) {

	lines := []string{"...#......", ".......#..", "#.........", "..........", "......#...", ".#........", ".........#",
		"..........", ".......#..", "#...#....."}

	expected := []int{3, 7}

	grid := NewGrid(lines)

	actual := grid.FindEmptyRows()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v :: Got %v", expected, actual)
	}
}

func TestFindEmptyColumns(t *testing.T) {

	lines := []string{"...#......", ".......#..", "#.........", "..........", "......#...", ".#........", ".........#",
		"..........", ".......#..", "#...#....."}

	expected := []int{2, 5, 8}

	grid := NewGrid(lines)

	actual := grid.FindEmptyColumns()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v :: Got %v", expected, actual)
	}
}

func TestFindGalaxies(t *testing.T) {

	lines := []string{"...#......", ".......#..", "#.........", "..........", "......#...", ".#........", ".........#",
		"..........", ".......#..", "#...#....."}

	expected := map[int]RowColumnIndex{1: NewRowColumnIndex(0, 3), 2: NewRowColumnIndex(1, 7), 3: NewRowColumnIndex(2, 0),
		4: NewRowColumnIndex(4, 6), 5: NewRowColumnIndex(5, 1), 6: NewRowColumnIndex(6, 9), 7: NewRowColumnIndex(8, 7),
		8: NewRowColumnIndex(9, 0), 9: NewRowColumnIndex(9, 4)}

	grid := NewGrid(lines)

	actual := grid.FindGalaxies()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v :: Got %v", expected, actual)
	}
}

type ShortestPathTestCase struct {
	pairName string
	expected int
}

func TestFindShortPathLength(t *testing.T) {
	lines := []string{"...#......", ".......#..", "#.........", "..........", "......#...", ".#........", ".........#",
		"..........", ".......#..", "#...#....."}

	tests := []ShortestPathTestCase{{pairName: "1-5", expected: 9}, {pairName: "1-7", expected: 15}, {pairName: "3-6", expected: 17},
		{pairName: "8-9", expected: 5}}

	grid := NewGrid(lines)

	galaxies := grid.FindGalaxies()

	pairs := grid.GenerateGalaxyPairs(galaxies)

	emptyRows := grid.FindEmptyRows()
	emptyColumns := grid.FindEmptyColumns()

	for _, tt := range tests {

		actual := FindShortestPathLength(pairs[tt.pairName], emptyRows, emptyColumns, 2, 2)

		if tt.expected != actual {
			t.Errorf("Expected %v :: Got %v", tt.expected, actual)
			return
		}
	}
}

func TestFindCombinedShortPathLengths(t *testing.T) {
	lines := []string{"...#......", ".......#..", "#.........", "..........", "......#...", ".#........", ".........#",
		"..........", ".......#..", "#...#....."}
	expected := 374
	grid := NewGrid(lines)

	actual := FindCombinedShortPathLengths(grid, 2, 2)

	if expected != actual {
		t.Errorf("Expected %v :: Got %v", expected, actual)
	}
}

func TestFindShortestPathLengthWeighted(t *testing.T) {
	lines := []string{"...#......", ".......#..", "#.........", "..........", "......#...", ".#........", ".........#",
		"..........", ".......#..", "#...#....."}

	tests := []ShortestPathTestCase{{pairName: "8-9", expected: 13}, {pairName: "1-5", expected: 25},
		{pairName: "3-6", expected: 49}}

	grid := NewGrid(lines)

	galaxies := grid.FindGalaxies()

	pairs := grid.GenerateGalaxyPairs(galaxies)

	emptyRows := grid.FindEmptyRows()
	emptyColumns := grid.FindEmptyColumns()

	for _, tt := range tests {

		actual := FindShortestPathLength(pairs[tt.pairName], emptyRows, emptyColumns, 10, 10)

		if tt.expected != actual {
			t.Errorf("Expected %v :: Got %v", tt.expected, actual)
			return
		}
	}
}
