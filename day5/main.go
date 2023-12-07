package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	data, err := readFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := part1(data)
	fmt.Println("Part 1 answer: ", result)

	// TODO: Part 2 is not complete yet
	result = part2(data)
	fmt.Println("Part 2 answer: ", result)
}

func part1(data string) int {
	lines := strings.Split(data, "\n")

	seeds := getSeeds(lines[0])

	mapRanges := getMapRanges(data)

	locations := make([]int, len(seeds))

	// Loop over the seeds
	for si, seed := range seeds {

		locations[si] = seed

		prevLocation := int(^uint(0) >> 1) // set prev to max
		hasChanged := false

		var prevMapVal int
		// Loop over maps
		for _, m := range mapRanges {
			prevMapVal = int(^uint(0) >> 1)

			// Check if seed is within any range
			for _, r := range m {
				destinationStart := r[0]
				sourceStart := r[1]
				rangeLen := r[2]

				if (sourceStart <= locations[si]) && (locations[si] <= sourceStart+rangeLen) {

					diff := sourceStart - destinationStart
					newLocation := locations[si] - diff

					if newLocation < prevMapVal {
						hasChanged = true
						prevLocation = locations[si] - diff
						prevMapVal = prevLocation
						break
					}
				}
			}
			if hasChanged {
				locations[si] = prevLocation
			}
		}
	}

	for _, l := range locations {
		println(l)
	}

	return Min(locations)

}

func part2(data string) int {
	lines := strings.Split(data, "\n")

	seedRanges := getSeedRanges(lines[0])

	mapRanges := getMapRanges(data)

	locations := make([]int, len(seedRanges))

	// Loop over the seeds
	for si, seedRange := range seedRanges {

		fmt.Println(time.Now(), " - Checking seedRange ", si, ": ", seedRange[0], seedRange[1])

		for seed := seedRange[0]; seed <= seedRange[1]; seed++ {

			locations[si] = seed

			prevLocation := int(^uint(0) >> 1) // set prev to max
			hasChanged := false

			var prevMapVal int
			// Loop over maps
			for _, m := range mapRanges {

				prevMapVal = int(^uint(0) >> 1)

				// Check if seed is within any range
				for _, r := range m {
					destinationStart := r[0]
					sourceStart := r[1]
					rangeLen := r[2]

					if (sourceStart <= locations[si]) && (locations[si] <= sourceStart+rangeLen) {

						diff := sourceStart - destinationStart
						newLocation := locations[si] - diff

						if newLocation < prevMapVal {
							hasChanged = true
							prevLocation = locations[si] - diff
							prevMapVal = prevLocation
							break
						}
					}
				}
				if hasChanged {
					locations[si] = prevLocation
				}
			}
		}
	}

	return Min(locations)

}

func Min(nums []int) int {
	minInt := int(^uint(0) >> 1)
	for _, s := range nums {
		if s < minInt {
			minInt = s
		}
	}
	return minInt
}

func getMapRanges(data string) [][][]int {
	//mapsStartIdx := 2
	maps := strings.Split(data, ":")[2:]
	println(len(maps))

	mapRanges := make([][][]int, len(maps))

	for i := 0; i < len(maps); i++ {

		// Get the lines
		mapString := maps[i]
		mapSlices := strings.Split(mapString, "\n")[1:]

		if i < len(maps)-1 {
			mapSlices = mapSlices[0 : len(mapSlices)-2] // remove two last lines
		}

		mapRanges[i] = make([][]int, len(mapSlices))

		for j, m := range mapSlices {
			mapRanges[i][j] = make([]int, len(m))

			for w, numStr := range strings.Fields(m) {

				mapRanges[i][j][w], _ = strconv.Atoi(numStr)
			}
		}
	}
	return mapRanges
}

func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Get seeds
func getSeeds(line string) []int {
	s := strings.Fields(strings.Split(line, ":")[1])

	seeds := make([]int, len(s))
	for i := 0; i < len(seeds); i++ {
		seeds[i], _ = strconv.Atoi(s[i])
	}
	return seeds
}

func getSeedRanges(line string) [][]int {
	s := strings.Fields(strings.Split(line, ":")[1])

	seedRanges := make([][]int, len(s)/2)
	x := 0
	for i := 0; i < len(seedRanges); i++ {
		start, _ := strconv.Atoi(s[x])
		end, _ := strconv.Atoi(s[x+1])
		seedRanges[i] = make([]int, 2)
		if start > end {
			seedRanges[i][0] = end
			seedRanges[i][1] = start
		} else {
			seedRanges[i][0] = start
			seedRanges[i][1] = end
		}
		x += 2
	}
	return seedRanges
}
