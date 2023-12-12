package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Transformation represents a mapping from a source range to a destination range.
type Transformation struct {
	DestStart   int
	SourceStart int
	Length      int
}

// TransformNumber transforms a number based on the provided transformations.
func TransformNumber(number int, transformations []Transformation) int {
	for _, t := range transformations {
		if number >= t.SourceStart && number < t.SourceStart+t.Length {
			return t.DestStart + (number - t.SourceStart)
		}
	}
	return number
}

// readMapAfterTitle reads the transformations from the file starting after the specified title until a blank line.
func readMapAfterTitle(scanner *bufio.Scanner, title string) []Transformation {
	var transformations []Transformation
	// Skip lines until the title is found
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == title {
			break
		}
	}
	// Read transformations until a blank line
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break // Stop reading at a blank line
		}
		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			fmt.Printf("Invalid line format: '%s'\n", line)
			continue
		}
		destStart, _ := strconv.Atoi(parts[0])
		sourceStart, _ := strconv.Atoi(parts[1])
		length, _ := strconv.Atoi(parts[2])
		transformations = append(transformations, Transformation{destStart, sourceStart, length})
	}
	return transformations
}

func main() {
	file, err := os.Open("day5.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read and generate seeds
	scanner.Scan()
	seedsLine := strings.TrimPrefix(scanner.Text(), "seeds: ")
	seedStrs := strings.Split(seedsLine, " ")
	var seeds []int
	for i := 0; i < len(seedStrs); i += 2 {
		start, _ := strconv.Atoi(seedStrs[i])
		length, _ := strconv.Atoi(seedStrs[i+1])
		for j := 0; j < length; j++ {
			seeds = append(seeds, start+j)
		}
	}

	// Read transformation maps
	seedToSoil := readMapAfterTitle(scanner, "seed-to-soil map:")
	soilToFertilizer := readMapAfterTitle(scanner, "soil-to-fertilizer map:")
	fertilizerToWater := readMapAfterTitle(scanner, "fertilizer-to-water map:")
	waterToLight := readMapAfterTitle(scanner, "water-to-light map:")
	lightToTemperature := readMapAfterTitle(scanner, "light-to-temperature map:")
	temperatureToHumidity := readMapAfterTitle(scanner, "temperature-to-humidity map:")
	humidityToLocation := readMapAfterTitle(scanner, "humidity-to-location map:")

	// Transform seeds to locations and find the lowest location number
	minLocation := -1
	for _, seed := range seeds {
		soil := TransformNumber(seed, seedToSoil)
		fertilizer := TransformNumber(soil, soilToFertilizer)
		water := TransformNumber(fertilizer, fertilizerToWater)
		light := TransformNumber(water, waterToLight)
		temperature := TransformNumber(light, lightToTemperature)
		humidity := TransformNumber(temperature, temperatureToHumidity)
		location := TransformNumber(humidity, humidityToLocation)

		if minLocation == -1 || location < minLocation {
			minLocation = location
		}
	}

	fmt.Println("The lowest location number is:", minLocation)
}
