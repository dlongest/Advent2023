package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Day6RunA() {
	fmt.Println("Running Day 6 Part A")

	races, err := LoadBoatRacesFrom("../Data/day6-example.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	waysToWinProduct := 1

	for _, r := range races {
		waysToWinProduct *= r.HowManyWaysToWin()
	}

	fmt.Println("Ways to win product = ", waysToWinProduct)
}

func Day6RunB() {
	fmt.Println("Running Day 6 Part B")

	race, err := LoadSingleBoatRaceFrom("../Data/day6.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	waysToWin := race.HowManyWaysToWin()

	fmt.Println("Ways to win = ", waysToWin)
}

type BoatRace struct {
	time           int
	recordDistance int
}

func LoadBoatRacesFrom(filePath string) ([]BoatRace, error) {
	lines := ReadLines(filePath)

	races, err := BoatRacesFrom(lines[0], lines[1])

	if err != nil {
		return nil, err
	}

	return races, nil
}
func LoadSingleBoatRaceFrom(filePath string) (BoatRace, error) {
	lines := ReadLines(filePath)

	race, err := BoatRaceFrom(lines[0], lines[1])

	if err != nil {
		return BoatRace{}, err
	}

	return race, nil
}

func BoatRaceFrom(times, distances string) (BoatRace, error) {

	re := regexp.MustCompile("[0-9]+")

	t := strings.Join(re.FindAllString(times, -1), "")

	d := strings.Join(re.FindAllString(distances, -1), "")

	timeAsInt, err := strconv.Atoi(t)

	if err != nil {
		return BoatRace{}, err
	}

	distanceAsInt, err := strconv.Atoi(d)

	if err != nil {
		return BoatRace{}, err
	}

	return BoatRace{time: timeAsInt, recordDistance: distanceAsInt}, nil
}

func BoatRacesFrom(times, distances string) ([]BoatRace, error) {

	re := regexp.MustCompile("[0-9]+")

	t, err := ToIntSlice(re.FindAllString(times, -1))

	if err != nil {
		return nil, err
	}

	d, err := ToIntSlice(re.FindAllString(distances, -1))

	if err != nil {
		return nil, err
	}

	races := make([]BoatRace, len(t))
	for i := range t {
		races[i] = BoatRace{time: t[i], recordDistance: d[i]}
	}

	return races, nil
}

func (br BoatRace) HowManyWaysToWin() int {

	lowest, highest := br.ButtonPressTimes()

	return highest - lowest + 1
}

func (br BoatRace) ButtonPressTimes() (int, int) {

	a := float64(-1)
	b := float64(br.time)
	c := float64(-br.recordDistance)

	first := (-b + math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a)
	second := (-b - math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a)

	var firstTime, secondTime int

	if IsIntegral(first) {
		firstTime = int(first) + 1
	} else {
		firstTime = int(math.Ceil(first))
	}

	if IsIntegral(second) {
		secondTime = int(second) - 1
	} else {
		secondTime = int(math.Floor(second))
	}

	return firstTime, secondTime
}
