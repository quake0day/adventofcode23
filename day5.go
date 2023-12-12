package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Transformation represents a single map, e.g., seed-to-soil.
type Transformation struct {
	DestStart   int
	SourceStart int
	Length      int
}

// TransformNumber transforms a given number based on a list of transformations.
func TransformNumber(number int, transformations []Transformation) int {
	for _, t := range transformations {
		if number >= t.SourceStart && number < t.SourceStart+t.Length {
			// Calculate the offset within the source range
			offset := number - t.SourceStart
			// Return the corresponding number in the destination range
			return t.DestStart + offset
		}
	}
	// If the number does not match any range, return it unchanged
	return number
}

// ReadTransformations reads and parses transformation maps from a file.
func ReadTransformations(scanner *bufio.Scanner) []Transformation {
	var transformations []Transformation
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			continue // Skip lines that don't have exactly three parts
		}
		destStart, _ := strconv.Atoi(parts[0])
		sourceStart, _ := strconv.Atoi(parts[1])
		length, _ := strconv.Atoi(parts[2])
		transformations = append(transformations, Transformation{destStart, sourceStart, length})
	}
	return transformations
}

// readNextMap reads the next transformation map, skipping any headers and blank lines.
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
		destStart, err1 := strconv.Atoi(parts[0])
		sourceStart, err2 := strconv.Atoi(parts[1])
		length, err3 := strconv.Atoi(parts[2])
		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Printf("Error parsing numbers in line: '%s'\n", line)
			continue
		}
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

	// Read seeds
	scanner.Scan()
	seedsLine := strings.TrimPrefix(scanner.Text(), "seeds: ")
	seedStrs := strings.Split(seedsLine, " ")
	var seeds []int
	for _, s := range seedStrs {
		seed, _ := strconv.Atoi(s)
		seeds = append(seeds, seed)
	}

	// Read transformation maps
	seedToSoil := readMapAfterTitle(scanner, "seed-to-soil map:")
	soilToFertilizer := readMapAfterTitle(scanner, "soil-to-fertilizer map:")
	fertilizerToWater := readMapAfterTitle(scanner, "fertilizer-to-water map:")
	waterToLight := readMapAfterTitle(scanner, "water-to-light map:")
	lightToTemperature := readMapAfterTitle(scanner, "light-to-temperature map:")
	temperatureToHumidity := readMapAfterTitle(scanner, "temperature-to-humidity map:")
	humidityToLocation := readMapAfterTitle(scanner, "humidity-to-location map:")

	fmt.Println(seedToSoil)
	fmt.Println(soilToFertilizer)
	// Transform seeds to locations
	var locations []int
	for _, seed := range seeds {
		fmt.Printf("Seed: %d\n", seed)

		soil := TransformNumber(seed, seedToSoil)
		fmt.Printf("  Soil: %d\n", soil)

		fertilizer := TransformNumber(soil, soilToFertilizer)
		fmt.Printf("  Fertilizer: %d\n", fertilizer)

		water := TransformNumber(fertilizer, fertilizerToWater)
		fmt.Printf("  Water: %d\n", water)

		light := TransformNumber(water, waterToLight)
		fmt.Printf("  Light: %d\n", light)

		temperature := TransformNumber(light, lightToTemperature)
		fmt.Printf("  Temperature: %d\n", temperature)

		humidity := TransformNumber(temperature, temperatureToHumidity)
		fmt.Printf("  Humidity: %d\n", humidity)

		location := TransformNumber(humidity, humidityToLocation)
		fmt.Printf("  Location: %d\n", location)

		fmt.Println("--------------------------------")

		locations = append(locations, location)
	}
	fmt.Println(locations)
	// Find the lowest location number
	lowestLocation := locations[0]
	for _, loc := range locations {
		if loc < lowestLocation {
			lowestLocation = loc
		}
	}

	fmt.Println("The lowest location number is:", lowestLocation)
}
