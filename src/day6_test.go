package main

import (
	"testing"
)

func TestBoatRace(t *testing.T) {

	races, err := LoadBoatRacesFrom("../Data/day6-example.txt")

	if err != nil {
		t.Error(err)
		return
	}

	expected := []int{4, 8, 9}

	for i, r := range races {
		actual := r.HowManyWaysToWin()

		if expected[i] != actual {
			t.Errorf("Expected %v ways to win and got actual %v", expected[i], actual)
		}
	}
}

func TestLoadSingleBoatRaceFrom(t *testing.T) {

	race, err := LoadSingleBoatRaceFrom("../Data/day6-example.txt")

	if err != nil {
		t.Error(err)
		return
	}

	waysToWin := race.HowManyWaysToWin()

	if waysToWin != 71503 {
		t.Errorf("Expected 71503 ways to win and only got %v\n", waysToWin)
	}
}

func TestBoatRaceButtonPressTimes(t *testing.T) {
	sut := BoatRace{time: 7, recordDistance: 9}

	first, second := sut.ButtonPressTimes()

	if first != 2 || second != 5 {
		t.Error("Button press times are not correct - should be 2 and 5")
	}
}

func TestBoatRaceHowManyWaysToWin(t *testing.T) {

	type testCase struct {
		time, distance, expectedWaysToWin int
	}

	testCases := []testCase{{time: 7, distance: 9, expectedWaysToWin: 4},
		{time: 15, distance: 40, expectedWaysToWin: 8}, {time: 30, distance: 200, expectedWaysToWin: 9}}

	//	testCases := []testCase{{time: 30, distance: 200, expectedWaysToWin: 9}}

	for _, tc := range testCases {
		actual := BoatRace{time: tc.time, recordDistance: tc.distance}.HowManyWaysToWin()

		if actual != tc.expectedWaysToWin {
			t.Errorf("Expected %v ways to win and got %v", tc.expectedWaysToWin, actual)
		}
	}
}
