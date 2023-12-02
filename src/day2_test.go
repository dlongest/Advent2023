package main

import (
	"testing"
)

type GameSpecificationTestCase struct {
	spec          GameSpecification
	game          Game
	shouldBeValid bool
}

func TestNewCubeSetFromString(t *testing.T) {
	input := "3 green, 15 blue, 14 red"
	want := CubeSet{red: 14, blue: 15, green: 3}

	got, err := NewCubeSetFromString(input)

	if err != nil {
		t.Error("Received error parsing cube set string into cube set: ", err.Error())
	}

	if got != want {
		t.Errorf("Created cube set %v does not match provided input string %v", got, input)
	}
}

func TestNewCubeSetsFromString(t *testing.T) {

	input := "1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"

	want := []CubeSet{NewCubeSet(3, 6, 1), NewCubeSet(6, 0, 3), NewCubeSet(14, 15, 3)}

	got, err := NewCubeSetsFromString(input)

	if err != nil {
		t.Error("Received error attempting to parse cube set '", input, "'")
	}

	for i := range got {

		if got[i] != want[i] {
			t.Error("Did not properly parse cube set from string '", input, "': got = %v :: want = %v", got[i], want[i])
		}
	}
}

func TestNewGame(t *testing.T) {

	input := "Game 21: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"

	want := Game{id: 21, cubeSets: []CubeSet{NewCubeSet(3, 6, 1), NewCubeSet(6, 0, 3), NewCubeSet(14, 15, 3)}}

	got, err := NewGameFromString(input)

	if err != nil {
		t.Error("Failed to create game")
	}

	if want.id != got.id {
		t.Errorf("Expected game ID %v does not match actual game ID %v", want.id, got.id)
	}

	if len(want.cubeSets) != len(got.cubeSets) {
		t.Error("Games do not have the same number of cube sets")
	}

	for i := range want.cubeSets {
		if got.cubeSets[i] != want.cubeSets[i] {
			t.Errorf("Expected cube set does match actual cube set: Expected %v :: Actual %v", want.cubeSets[i], got.cubeSets[i])
		}
	}
}

func TestGameSpecification(t *testing.T) {

	sut := GameSpecification{red: 5, blue: 5, green: 5}

	var tests = []GameSpecificationTestCase{
		{spec: sut, game: NewGame(1, NewCubeSet(5, 4, 3)), shouldBeValid: true},
		{spec: sut, game: NewGame(2, NewCubeSet(5, 6, 3)), shouldBeValid: false},
		{spec: sut, game: NewGame(3, NewCubeSet(2, 5, 5), NewCubeSet(1, 2, 5)), shouldBeValid: true},
		{spec: sut, game: NewGame(3, NewCubeSet(2, 5, 5), NewCubeSet(1, 2, 8)), shouldBeValid: false}}

	for _, tt := range tests {
		if tt.spec.IsValid(tt.game) != tt.shouldBeValid {
			t.Error("Failure on Game ID ", tt.game.id)
		}
	}
}

func TestNewGameSpecificationSatisfying(t *testing.T) {
	input := Game{id: 21, cubeSets: []CubeSet{NewCubeSet(3, 6, 1), NewCubeSet(6, 0, 3), NewCubeSet(14, 15, 3)}}

	want := GameSpecification{red: 14, green: 3, blue: 15}

	got := NewGameSpecificationSatisfying(input)

	if got != want {
		t.Error("Expected game specification does not match actual game specification")
	}
}

func TestGameSpecificationPower(t *testing.T) {
	want := 240
	spec := GameSpecification{red: 8, blue: 3, green: 10}

	got := spec.Power()

	if want != got {
		t.Error("Expected power calculation doesn't match actual calculation")
	}
}
