package jour5

import (
	"advent_of_code_2023/utils"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type categoryMap struct {
	destStart int
	srcStart  int
	rangeLen  int
}

func Main() string {
	inputLines := utils.ReadFileLines("jour5/input.txt")
	// inputLines = utils.ReadFileLines("jour5/test_input.txt")
	return utils.LogDayAnswers(inputLines, []func([]string) (string, error){part1, part2})
}

func getSeeds(seedLine string) []int {
	seedStrings := strings.Split(seedLine[7:len(seedLine)-1], " ")

	var seeds []int
	for _, seedString := range seedStrings {
		seed, _ := strconv.Atoi(seedString)
		seeds = append(seeds, seed)
	}

	return seeds
}

func splitMaps(mapsLines []string) [][]categoryMap {
	maps := make([][]categoryMap, 7)

	mapId := 0
	for i := 0; i < len(mapsLines); i++ {
		line := mapsLines[i][:len(mapsLines[i])-1]

		if len(line) == 0 {
			mapId++
			i++
			continue
		}

		itemString := strings.Split(line, " ")
		destStart, _ := strconv.Atoi(itemString[0])
		srcStart, _ := strconv.Atoi(itemString[1])
		rangeLen, _ := strconv.Atoi(itemString[2])
		maps[mapId] = append(maps[mapId], categoryMap{destStart: destStart, srcStart: srcStart, rangeLen: rangeLen})
	}

	return maps
}

func translate(seed int, categoryMaps [][]categoryMap) int {
	for _, categoryMap := range categoryMaps {
		for _, categoryLine := range categoryMap {
			if categoryLine.srcStart <= seed && seed < categoryLine.srcStart+categoryLine.rangeLen {
				seed = categoryLine.destStart + seed - categoryLine.srcStart
				break
			}
		}
	}

	return seed
}

func part1(inputLines []string) (string, error) {
	seeds := getSeeds(inputLines[0])
	categoryMaps := splitMaps(inputLines[3:])

	lowest := translate(seeds[0], categoryMaps)
	for _, seed := range seeds[0:] {
		num := translate(seed, categoryMaps)
		if num < lowest {
			lowest = num
		}
	}

	return fmt.Sprint(lowest), nil
}

func part2(inputLines []string) (string, error) {
	seedRanges := getSeeds(inputLines[0])
	categoryMaps := splitMaps(inputLines[3:])

	var wg sync.WaitGroup

	totalCompute := 0

	var lowests []int
	for i := 0; i < len(seedRanges); i += 2 {
		wg.Add(1)
		go func(startSeed int, maxRange int, lowests *[]int, categoryMaps [][]categoryMap, compute *int, wg *sync.WaitGroup) {
			localLowest := translate(startSeed, categoryMaps)
			for pos := 0; pos < maxRange; pos++ {
				newTranslation := translate(startSeed+pos, categoryMaps)
				*compute++
				if newTranslation < localLowest {
					localLowest = newTranslation
				}
			}
			*lowests = append(*lowests, localLowest)
			defer wg.Done()
		}(seedRanges[i], seedRanges[i+1], &lowests, categoryMaps, &totalCompute, &wg)
	}

	wg.Wait()

	lowest := lowests[0]
	for _, seed := range lowests {
		if seed < lowest {
			lowest = seed
		}
	}
	fmt.Printf("Computed %d seeds\n", totalCompute)

	return fmt.Sprint(lowest), nil
}
