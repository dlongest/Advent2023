package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func Day5RunA() {
	fmt.Println("Running Day 5 Run A")

	lines := ReadLines("../Data/day5.txt")

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

	lines := ReadLines("../Data/day5.txt")

	sut, err := NewAlmanacWithSeedRanges(lines)

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

//func Merge(left *CategoryRangeSet, right *CategoryRangeSet, categoryNameFunc func(string, string) string) *CategoryRangeSet {
//newSet := &CategoryRangeSet{}

//newCategoryName := categoryNameFunc(string(left.category), string(right.category))
//newCategoryName := strings.Split(string(c.category), "-")[0] + strings.Split(string(other.category), "-")[1]

// To merge the two range sets together that accomplish A -> B -> C, we need to find the sub-ranges that allow for A -> C directly

// Think to do this, start by computing the sub-ranges in B -> C that also exist in A -> B

// temperature-to-humidity map:
// 0 69 1
// 1 0 69

// humidity-to-location map:
// 60 56 37
// 56 93 4

// temp-humidity     defines the ranges [69] -> [0] and [0, 68] -> [1, 69]
// humidity-location defines the ranges [56, 92] -> [60, 96] and [93, 96] -> [56, 59]

// First we ask: where does the RHS of temp-humidity fit into the LHS of humidity-location?
// * [69] -> [0] doesn't align anywhere as [0] is not a source range for humidity-location
// * [56, 69] (the sub-range of [1-69]) fits into the larger range [56, 92]
//
// * [56, 92] humidity maps to [60, 96] location
// * What's the source range on the temperature side matching [56, 92] on the humidity side? [55, 91]
//
// So the temperature-location combined map is [55, 91] -> [60, 96], expressed as 60, 55, 37
//
// Maybe this should be defined on the GardenRange itself:
// * Given 2 GardenRanges where one is said to be a direct precursor to the other, return the GardenRanges that
// 	combine them together.

// 	return nil
// }

type GardenRange struct {
	destinationStart, sourceStart, rangeSize int
	destinationEnd, sourceEnd                int
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

func (gr GardenRange) Flip() GardenRange {
	return GardenRange{destinationStart: gr.sourceStart, destinationEnd: gr.sourceEnd, sourceStart: gr.destinationStart,
		sourceEnd: gr.destinationEnd, rangeSize: gr.rangeSize}
}

func (gr GardenRange) DestinationFor(sourceValue int) int {

	// If sourceValue is outside the contained range, the destination is the same as the source value
	if sourceValue < gr.sourceStart || sourceValue > gr.sourceEnd {
		return sourceValue
	}

	offset := sourceValue - gr.sourceStart

	return gr.destinationStart + offset
}

func (gr GardenRange) SourceFor(destinationValue int) int {
	return gr.Flip().DestinationFor(destinationValue)
	// if destinationValue < gr.destinationStart || destinationValue > gr.destinationEnd {
	// 	return destinationValue
	// }

	// offset := destinationValue - gr.destinationStart

	// return gr.sourceStart + offset
}

func NewGardenRange(destinationStart int, sourceStart int, rangeSize int) GardenRange {
	destinationEnd := destinationStart + rangeSize - 1
	sourceEnd := sourceStart + rangeSize - 1

	return GardenRange{destinationStart: destinationStart, sourceStart: sourceStart, rangeSize: rangeSize,
		destinationEnd: destinationEnd, sourceEnd: sourceEnd}
}

func Merge(prior GardenRange, next GardenRange) (GardenRange, bool) {

	// The entire range of prior is before the range of next
	if prior.destinationEnd < next.sourceStart {
		return GardenRange{}, false
	}

	// The entire range of prior is after the range of next
	if prior.destinationStart > next.sourceEnd {
		return GardenRange{}, false
	}

	// We know there is some overlap of the prior.destination and next.source ranges. We need
	// to find the specific sub-range.
	mergeStart := Maximum(prior.destinationStart, next.sourceStart)
	mergeEnd := Minimum(prior.destinationEnd, next.sourceEnd)

	// Now we need the prior.source values that correspond with mergeStart and mergeEnd in prior.
	newPriorSourceStart := prior.SourceFor(mergeStart)
	newRangeSize := mergeEnd - mergeStart + 1

	// We also need next.destination values that corresspond
	newNextDestinationStart := next.DestinationFor(mergeStart)

	return NewGardenRange(newNextDestinationStart, newPriorSourceStart, newRangeSize), true
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

func NewAlmanacWithSeedRanges(lines []string) (*Almanac, error) {

	almanac := &Almanac{}

	// This line has to change to be seed ranges
	seedRangePairs, err := ToIntSlice(regexp.MustCompile("[0-9]+").FindAllString(lines[0], -1))

	seeds := []int{}

	// We'll see if this runs...
	for i := 0; i < len(seedRangePairs)-1; i += 2 {

		startingSeed, seedRangeSize := seedRangePairs[i], seedRangePairs[i+1]
		for i := startingSeed; i < startingSeed+seedRangeSize; i++ {
			seeds = append(seeds, i)
		}
	}

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
