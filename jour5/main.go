package jour5

import (
	"advent_of_code_2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func Main() string {
	// inputLines := utils.ReadFileLines("jour5/input.txt")
	inputLines := utils.ReadFileLines("jour5/test_input.txt")
	return utils.LogDayAnswers(inputLines, []func([]string) (string, error){part1, part2})
}

func getSeeds(seedLine string) []int {
	seedStrings := strings.Split(seedLine[7:], " ")

	var seeds []int
	for _, seedString := range seedStrings {
		seed, _ := strconv.Atoi(seedString)
		seeds = append(seeds, seed)
	}

	return seeds
}

func splitMaps(mapsLines []string) [][][]int {
	maps := make([][][]int, 7)

	mapId := 0
	for i := 0; i < len(mapsLines); i++ {
		line := mapsLines[i]

		if len(line) == 1 {
			mapId++
			i++
			continue
		}

		for _, itemString := range strings.Split(mapsLines[i], " ") {
			if itemString == "" {
				continue
			}
			fmt.Println(itemString)
		}

	}

	return maps
}

func part1(inputLines []string) (string, error) {
	// seeds := getSeeds(inputLines[0])
	maps := splitMaps(inputLines[3:])
	fmt.Println(maps)

	return "", nil
}

func part2(inputLines []string) (string, error) {

	return "", nil
}
