package main

import "fmt"

// Calculates the number of ways to beat the record for a single race.
func waysToBeatRecord(raceTime, recordDistance int) int {
	ways := 0
	for buttonHoldTime := 0; buttonHoldTime < raceTime; buttonHoldTime++ {
		speed := buttonHoldTime
		moveTime := raceTime - buttonHoldTime
		distance := speed * moveTime
		if distance > recordDistance {
			ways++
		}
	}
	return ways
}

func main() {
	// Example races
	races := []struct {
		time   int
		record int
	}{
		{50, 242},
		{74, 1017},
		{86, 1691},
		{85, 1252},
	}

	totalWays := 1
	for _, race := range races {
		ways := waysToBeatRecord(race.time, race.record)
		fmt.Printf("Race (Time: %dms, Record: %dmm): %d ways to win\n", race.time, race.record, ways)
		totalWays *= ways
	}

	fmt.Printf("Total ways to beat all records: %d\n", totalWays)
}
