package main

import (
	"reflect"
	"testing"
)

func TestNewAlmanac(t *testing.T) {

	lines := ReadLines("../Data/day5-example.txt")

	sut, err := NewAlmanac(lines)

	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(sut.seeds, []int{79, 14, 55, 13}) {
		t.Error("Almanac does not have the correct seeds")
	}

	if sut.maps["seed-soil"].ranges[0] != NewGardenRange(50, 98, 2) ||
		sut.maps["seed-soil"].ranges[1] != NewGardenRange(52, 50, 48) {
		t.Error("Incorrect seed-soil range")
	}

	if sut.maps["soil-fertilizer"].ranges[0] != NewGardenRange(0, 15, 37) ||
		sut.maps["soil-fertilizer"].ranges[1] != NewGardenRange(37, 52, 2) ||
		sut.maps["soil-fertilizer"].ranges[2] != NewGardenRange(39, 0, 15) {
		t.Error("Incorrect soil-fertilizer range")
	}

	if sut.maps["fertilizer-water"].ranges[0] != NewGardenRange(49, 53, 8) ||
		sut.maps["fertilizer-water"].ranges[1] != NewGardenRange(0, 11, 42) ||
		sut.maps["fertilizer-water"].ranges[2] != NewGardenRange(42, 0, 7) ||
		sut.maps["fertilizer-water"].ranges[3] != NewGardenRange(57, 7, 4) {
		t.Error("Incorrect fertilizer-water range")
	}

	if sut.maps["water-light"].ranges[0] != NewGardenRange(88, 18, 7) ||
		sut.maps["water-light"].ranges[1] != NewGardenRange(18, 25, 70) {
		t.Error("Incorrect water-light range")
	}

	if sut.maps["light-temperature"].ranges[0] != NewGardenRange(45, 77, 23) ||
		sut.maps["light-temperature"].ranges[1] != NewGardenRange(81, 45, 19) ||
		sut.maps["light-temperature"].ranges[2] != NewGardenRange(68, 64, 13) {
		t.Error("Incorrect light-temperature range")
	}

	if sut.maps["temperature-humidity"].ranges[0] != NewGardenRange(0, 69, 1) ||
		sut.maps["temperature-humidity"].ranges[1] != NewGardenRange(1, 0, 69) {
		t.Error("Incorrect temperature-humidity range")
	}

	if sut.maps["humidity-location"].ranges[0] != NewGardenRange(60, 56, 37) ||
		sut.maps["humidity-location"].ranges[1] != NewGardenRange(56, 93, 4) {
		t.Error("Incorrect humidity-location range")
	}
}

func TestAlmanacMap(t *testing.T) {
	lines := ReadLines("../Data/day5-example.txt")

	sut, err := NewAlmanac(lines)

	if err != nil {
		t.Error(err)
		return
	}

	traversal, err := sut.Map(79)

	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(traversal, []int{79, 81, 81, 81, 74, 78, 78, 82}) {
		t.Error("Traversal path incorrect")
	}
}

func TestAlmanacLocationFor(t *testing.T) {
	lines := ReadLines("../Data/day5-example.txt")

	sut, err := NewAlmanac(lines)

	if err != nil {
		t.Error(err)
		return
	}

	location, err := sut.LocationFor(79)

	if err != nil {
		t.Error(err)
		return
	}

	if location != 82 {
		t.Error("Location not correct")
	}
}

func TestAlmanacLocations(t *testing.T) {
	lines := ReadLines("../Data/day5-example.txt")

	sut, err := NewAlmanac(lines)

	if err != nil {
		t.Error(err)
		return
	}

	locations, err := sut.Locations()

	if err != nil {
		t.Error(err)
		return
	}

	if locations[79] != 82 || locations[14] != 43 || locations[55] != 86 || locations[13] != 35 {
		t.Error("Not all seed locations are correct. Actual: ", locations)
	}
}

func TestAlmanacTargetLocation(t *testing.T) {
	lines := ReadLines("../Data/day5-example.txt")

	sut, err := NewAlmanac(lines)

	if err != nil {
		t.Error(err)
		return
	}

	lowestLocationSeed, lowestLocation, err := sut.TargetLocation()

	if err != nil {
		t.Error(err)
		return
	}

	if lowestLocationSeed != 13 || lowestLocation != 35 {
		t.Error("Lowest location seed or location is not correct")
	}
}

type GardenRangeContainsTestCase struct {
	sut                 GardenRange
	testSourceValue     int
	expectedDestination int
	expectedFound       bool
}

func TestGardenRangeContains(t *testing.T) {

	tc := []GardenRangeContainsTestCase{
		{sut: NewGardenRange(50, 98, 2), testSourceValue: 98, expectedDestination: 50, expectedFound: true},
		{sut: NewGardenRange(50, 98, 2), testSourceValue: 97, expectedFound: false},
		{sut: NewGardenRange(50, 98, 2), testSourceValue: 99, expectedDestination: 51, expectedFound: true},
		{sut: NewGardenRange(50, 98, 2), testSourceValue: 100, expectedFound: false}}

	for _, tt := range tc {

		destination, found := tt.sut.Contains(tt.testSourceValue)

		if tt.expectedFound != found {
			t.Errorf("Source value %v should have been found and wasn't", tt.testSourceValue)
			return
		}

		if !tt.expectedFound {
			// We're done here - the output of destination should not be checked
			return
		}

		if tt.expectedDestination != destination {
			t.Errorf("Expected destination %v doesn't match actual destination %v", tt.expectedDestination, destination)
		}
	}
}

func TestCategoryRangeSetContains(t *testing.T) {

	sut := NewCategoryRangeSet("seed-soil")

	sut.AddRange(NewGardenRange(50, 98, 2))
	sut.AddRange(NewGardenRange(52, 50, 48))

	if sut.Map(10) != 10 {
		t.Error("Source 10 should have mapped to destination 10")
	}

	if sut.Map(98) != 50 {
		t.Error("Source 98 should have mapped to destination 50")
	}

	if sut.Map(99) != 51 {
		t.Error("Source 99 should have mapped to destination 51")
	}

	if sut.Map(50) != 52 {
		t.Error("Source 50 should have mapped to destination 52")
	}

	if sut.Map(49) != 49 {
		t.Error("Source 49 should have mapped to destination 49")
	}

	if sut.Map(79) != 81 {
		t.Error("Source 79 should have mapped to destination 81")
	}
}
