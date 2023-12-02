package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Day2RunA() {
	fmt.Println("Running Day 2 Part A")

	lines := ReadLines("../Data/day2.txt")

	spec := GameSpecification{red: 12, green: 13, blue: 14}

	validTally := 0
	validIDSum := 0

	for _, line := range lines {

		game, err := NewGameFromString(line)

		if err != nil {
			fmt.Println("Error processing " + line)
		}

		if spec.IsValid(game) {
			validTally++
			validIDSum += game.id
		}
	}

	fmt.Printf("Valid Games = %v\n", validTally)
	fmt.Printf("Valid Game ID Sum = %v\n", validIDSum)
}

func Day2RunB() {
	fmt.Println("Running Day 2 Part B")

	lines := ReadLines("../Data/day2.txt")

	powerSum := 0

	for _, line := range lines {
		game, err := NewGameFromString(line)

		if err != nil {
			fmt.Println("Error creating game for input '" + line + "'")
		}

		spec := NewGameSpecificationSatisfying(game)
		power := spec.Power()
		powerSum += power
	}

	fmt.Println("Total Power = ", powerSum)
}

type CubeSet struct {
	red, blue, green int
}

func NewCubeSet(red, blue, green int) CubeSet {
	return CubeSet{red: red, blue: blue, green: green}
}

// NewCubeSetsFromString will receive a semi-colon delimited list of segments matching the below format:
// [<B> blue,][<R> red,][<G> green]
// Each segment corresponds to the cube set for a single game. Each cubeset will have at least 1 color and at most 3.
// The colors can be in any order within the cube set (blue isn't always first). Here's an example with 3 cubesets:
// 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
func NewCubeSetsFromString(s string) ([]CubeSet, error) {
	cubeSets := make([]CubeSet, 0)

	for _, cubeSetString := range strings.Split(s, "; ") {

		cubeSet, err := NewCubeSetFromString(cubeSetString)

		if err != nil {
			return cubeSets, errors.New("Unable to create cube set from provided string")
		}

		cubeSets = append(cubeSets, cubeSet)
	}

	return cubeSets, nil
}

// NewCubeSetFromString returns a CubeSet matching the provided input string in a format resembling the below:
// 1 green, 3 red, 6 blue
// The 3 colors can appear in any order. There will always be at least 1 color, but never more than 3.
func NewCubeSetFromString(s string) (CubeSet, error) {

	green := 0
	blue := 0
	red := 0

	for _, colors := range strings.Split(s, ", ") {

		singleColorSlice := strings.Split(colors, " ")

		quantity, err := strconv.Atoi(singleColorSlice[0])

		if err != nil {
			return CubeSet{}, errors.New("failed to correctly parse quantity for color in cube set")
		}

		switch singleColorSlice[1] {
		case "red":
			red = quantity
		case "blue":
			blue = quantity
		case "green":
			green = quantity
		}
	}

	return NewCubeSet(red, blue, green), nil
}

type Game struct {
	id       int
	cubeSets []CubeSet
}

// NewGameFromString creates a new Game assuming the string is the in following format (<X> is a placeholder for value X):
// Game <ID>: [<B> blue,][<R> red,][<G> green][;...]
func NewGameFromString(s string) (Game, error) {

	parts := strings.Split(s, ": ")

	id, err := strconv.Atoi(strings.Split(parts[0], " ")[1])

	if err != nil {
		return Game{}, errors.New("Unable to parse Game ID from input '" + s + "'")
	}

	cubeSets, err := NewCubeSetsFromString(parts[1])

	if err != nil {
		return Game{}, errors.New("Unable to create cube sets from input '" + s + "'")
	}

	return NewGame(id, cubeSets...), nil
}

func NewGame(id int, cubeSets ...CubeSet) Game {
	return Game{id: id, cubeSets: cubeSets}
}

type GameSpecification struct {
	red, blue, green int
}

func (spec GameSpecification) IsValid(game Game) bool {

	for _, set := range game.cubeSets {
		if spec.blue < set.blue || spec.red < set.red || spec.green < set.green {
			return false
		}
	}

	return true
}

func (spec GameSpecification) Power() int {
	return spec.red * spec.green * spec.blue
}

// NewGameSpecificationSatisfying finds the minimum specification to satisfy the game. In other words,
// a game consisting of 1 or more cube sets is satisfied by the minimum specification where red, green, and blue
// is the maximum of each color across all cube sets in the game.
func NewGameSpecificationSatisfying(game Game) GameSpecification {

	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	for _, cubeSet := range game.cubeSets {
		if cubeSet.red > maxRed {
			maxRed = cubeSet.red
		}

		if cubeSet.green > maxGreen {
			maxGreen = cubeSet.green
		}

		if cubeSet.blue > maxBlue {
			maxBlue = cubeSet.blue
		}
	}

	return GameSpecification{red: maxRed, green: maxGreen, blue: maxBlue}
}
