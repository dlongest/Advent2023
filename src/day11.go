package main

import (
	"fmt"
	"strconv"
)

func Day11RunA() {
	fmt.Println("Running Day 11 Part A")

	lines := ReadLines("../Data/day11-example.txt")

	grid := NewGrid(lines)

	sum := FindCombinedShortPathLengths(grid, 2, 2)

	fmt.Printf("Sum = %v\n", sum)
}

func Day11RunB() {
	fmt.Println("Running Day 11 Part B")

	lines := ReadLines("../Data/day11.txt")

	grid := NewGrid(lines)

	sum := FindCombinedShortPathLengths(grid, 1000000, 1000000)

	fmt.Printf("Sum = %v\n", sum)
}

func (grid *Grid) FindGalaxies() map[int]RowColumnIndex {

	galaxies := map[int]RowColumnIndex{}

	galaxyID := 1

	for rowIndex := range grid.grid {
		for colIndex, val := range grid.grid[rowIndex] {
			if val == '#' {
				galaxies[galaxyID] = NewRowColumnIndex(rowIndex, colIndex)
				galaxyID++
			}
		}
	}

	return galaxies
}

func (grid *Grid) FindEmptyColumns() []int {
	emptyCols := []int{}

	for colIndex := 0; colIndex < grid.columnCount; colIndex++ {

		rowIndex := 0

		for rowIndex < grid.rowCount && grid.grid[rowIndex][colIndex] == '.' {
			rowIndex++
		}

		if rowIndex == grid.columnCount {
			emptyCols = append(emptyCols, colIndex)
		}
	}

	return emptyCols
}

func (grid *Grid) FindEmptyRows() []int {

	emptyRows := []int{}

	for rowIndex := range grid.grid {

		colIndex := 0

		for colIndex < grid.columnCount && grid.grid[rowIndex][colIndex] == '.' {
			colIndex++
		}

		if colIndex == grid.columnCount {
			emptyRows = append(emptyRows, rowIndex)
		}
	}

	return emptyRows
}

func (grid *Grid) GenerateGalaxyPairs(galaxies map[int]RowColumnIndex) map[string]RowColumnIndexPair {

	pairs := map[string]RowColumnIndexPair{}

	for i := 1; i < len(galaxies); i++ {

		for j := i + 1; j <= len(galaxies); j++ {
			name := strconv.Itoa(i) + "-" + strconv.Itoa(j)

			pairs[name] = NewRowColumnIndexPair(galaxies[i], galaxies[j])
		}
	}

	return pairs
}

func FindCombinedShortPathLengths(grid *Grid, rowWeight int, columnWeight int) int {
	galaxies := grid.FindGalaxies()

	pairs := grid.GenerateGalaxyPairs(galaxies)

	emptyRows := grid.FindEmptyRows()
	emptyColumns := grid.FindEmptyColumns()

	sum := 0

	for _, pair := range pairs {

		length := FindShortestPathLength(pair, emptyRows, emptyColumns, rowWeight, columnWeight)
		sum += length
	}

	return sum
}

func FindShortestPathLength(pair RowColumnIndexPair, emptyRows []int, emptyColumns []int, rowWeight int, columnWeight int) int {

	emptyRowsCrossed := HowManyEmptyRowsCrossed(pair, emptyRows)
	emptyColumnsCrossed := HowManyEmptyColumnsCrossed(pair, emptyColumns)

	x_distance := iabs(pair.start.columnIndex-pair.end.columnIndex) - emptyRowsCrossed
	y_distance := iabs(pair.start.rowIndex-pair.end.rowIndex) - emptyColumnsCrossed

	return x_distance + y_distance + (emptyRowsCrossed * rowWeight) + (emptyColumnsCrossed * columnWeight)
}

func HowManyEmptyRowsCrossed(pair RowColumnIndexPair, emptyRows []int) int {

	minRow := imin(pair.start.rowIndex, pair.end.rowIndex)
	maxRow := imax(pair.start.rowIndex, pair.end.rowIndex)

	count := 0

	for _, emptyRow := range emptyRows {
		if emptyRow > minRow && emptyRow < maxRow {
			count++
		}
	}

	return count
}

func HowManyEmptyColumnsCrossed(pair RowColumnIndexPair, emptyColumns []int) int {

	minColumn := imin(pair.start.columnIndex, pair.end.columnIndex)
	maxColumn := imax(pair.start.columnIndex, pair.end.columnIndex)

	count := 0

	for _, emptyColumn := range emptyColumns {
		if emptyColumn > minColumn && emptyColumn < maxColumn {
			count++
		}
	}

	return count
}
