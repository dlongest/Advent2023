package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func Day5RunA() {
	fmt.Println("Running Day 5 Run A")

	lines := ReadLines("../Data/day5-example.txt")

	sut, err := NewAlmanac(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	lowestLocationSeed, lowestLocation, err := sut.TargetLocation()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Lowest location %v was found starting from seed %v", lowestLocation, lowestLocationSeed)
}

func Day5RunB() {
	fmt.Println("Running Day 5 Run B")
}

// const (
// 	seed_soil            int = iota
// 	soil_fertilizer      int
// 	fertilizer_water     int
// 	water_light          int
// 	light_temperature    int
// 	temperature_humidity int
// 	humidity_location    int
// )

type Category string

type CategoryRangeSet struct {
	category Category
	ranges   []GardenRange
}

func NewCategoryRangeSet(category Category) *CategoryRangeSet {
	c := &CategoryRangeSet{category: category}

	return c
}

func (c *CategoryRangeSet) AddRange(gardenRange GardenRange) {
	c.ranges = append(c.ranges, gardenRange)
}

func (c *CategoryRangeSet) AddNewRangeFromSlice(rangeSlice []int) {
	c.AddNewRange(rangeSlice[0], rangeSlice[1], rangeSlice[2])
}

func (c *CategoryRangeSet) AddNewRange(destinationStart, sourceStart, rangeSize int) {
	r := NewGardenRange(destinationStart, sourceStart, rangeSize)

	c.ranges = append(c.ranges, r)
}

func (c *CategoryRangeSet) Map(sourceValue int) int {

	for _, r := range c.ranges {
		destination, found := r.Contains(sourceValue)

		if found {
			return destination
		}
	}

	// If the source is not mapped to a new destination based on the contained ranges,
	// the destination is the same as the source value
	return sourceValue
}

type GardenRange struct {
	destinationStart, sourceStart, rangeSize int
}

func (gr GardenRange) Contains(sourceValue int) (int, bool) {
	// If sourceValue is less than sourceStart, we immediately know it's not in the range
	if sourceValue < gr.sourceStart {
		return -1, false
	}

	// If sourceValue is greater than the last possible allowable value for source in this ragnge,
	// we know it's not in the range
	if sourceValue > gr.sourceStart+gr.rangeSize-1 {
		return -1, false
	}

	// Compute the destination offset for sourceValue
	offset := sourceValue - gr.sourceStart

	return gr.destinationStart + offset, true
}

func NewGardenRange(destinationStart int, sourceStart int, rangeSize int) GardenRange {
	return GardenRange{destinationStart: destinationStart, sourceStart: sourceStart, rangeSize: rangeSize}
}

type Almanac struct {
	seeds      []int
	categories []Category
	maps       map[Category]*CategoryRangeSet
}

func NewAlmanac(lines []string) (*Almanac, error) {

	almanac := &Almanac{}

	seeds, err := ToIntSlice(regexp.MustCompile("[0-9]+").FindAllString(lines[0], -1))

	if err != nil {
		return nil, err
	}

	almanac.seeds = seeds
	almanac.categories = []Category{"seed-soil", "soil-fertilizer", "fertilizer-water", "water-light",
		"light-temperature", "temperature-humidity", "humidity-location"}

	almanac.maps = make(map[Category]*CategoryRangeSet)

	for _, category := range almanac.categories {
		almanac.maps[category] = NewCategoryRangeSet(category)
	}

	currentCategoryIndex := 0
	currentCategoryMapKey := almanac.categories[0]

	for lineIndex := 2; lineIndex < len(lines); lineIndex++ {

		// Once we hit a blank line, we're done processing that category so move to next and start loop again
		if len(lines[lineIndex]) == 0 {
			currentCategoryIndex++
			currentCategoryMapKey = almanac.categories[currentCategoryIndex]
			almanac.maps[currentCategoryMapKey] = NewCategoryRangeSet(almanac.categories[currentCategoryIndex])
			continue
		}

		// The first line of a given category-category map contains the word "map", but we don't
		// need to read these dynamically so just restart the loop
		if strings.Contains(lines[lineIndex], ("map")) {
			continue
		}

		rangeParts, err := extractIntTriplet(lines[lineIndex])

		if err != nil {
			return nil, err
		}

		almanac.maps[currentCategoryMapKey].AddNewRangeFromSlice(rangeParts)

	}

	return almanac, nil
}

func extractIntTriplet(tripletString string) ([]int, error) {
	parts := strings.Split(tripletString, " ")

	return ToIntSlice(parts)
}

func (a *Almanac) TargetLocation() (lowestLocationSeed int, lowestLocation int, err error) {

	lowestLocation = math.MaxInt
	lowestLocationSeed = -1

	for _, seed := range a.seeds {
		location, err := a.LocationFor(seed)

		if err != nil {
			return -1, -1, err
		}

		if location < lowestLocation {
			lowestLocation = location
			lowestLocationSeed = seed
		}
	}

	return lowestLocationSeed, lowestLocation, nil
}

// Locations finds the locations for all the contained seeds using the contained
// category maps.
func (a *Almanac) Locations() (map[int]int, error) {

	locationsBySeed := map[int]int{}

	for _, seed := range a.seeds {
		location, err := a.LocationFor(seed)

		if err != nil {
			return locationsBySeed, err
		}

		locationsBySeed[seed] = location
	}

	return locationsBySeed, nil
}

func (a *Almanac) LocationFor(seed int) (int, error) {
	traversal, err := a.Map(seed)

	if err != nil {
		return -1, err
	}

	return traversal[len(traversal)-1], nil
}

func (a *Almanac) Map(seed int) ([]int, error) {

	traversal := []int{seed}

	currentSourceValue := seed

	for _, category := range a.categories {

		nextDestination := a.maps[category].Map(currentSourceValue)

		traversal = append(traversal, nextDestination)

		currentSourceValue = nextDestination
	}

	return traversal, nil
}
