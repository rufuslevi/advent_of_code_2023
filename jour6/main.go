package jour6

import (
	"advent_of_code_2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func Main() string {
	inputLines := utils.ReadFileLines("jour6/input.txt")
	// inputLines = utils.ReadFileLines("jour6/test_input.txt")
	return utils.LogDayAnswers(inputLines, []func([]string) (string, error){part1, part2})
}

func getNumberOfBestTimes(time int, distance int) int {
	best := 0
	for i := 1; i < time; i++ {
		newBestDistance := (time - i) * i
		if newBestDistance > distance {
			best++
		}
	}

	return best
}

func part1(inputLines []string) (string, error) {
	timesStr := strings.Fields(inputLines[0])[1:]
	distancesStr := strings.Fields(inputLines[1])[1:]
	var times []int
	var distances []int
	for pos, timeStr := range timesStr {
		time, _ := strconv.Atoi(timeStr)
		times = append(times, time)

		distance, _ := strconv.Atoi(distancesStr[pos])
		distances = append(distances, distance)
	}

	mul := getNumberOfBestTimes(times[0], distances[0])
	for i := 1; i < len(times); i++ {
		mul *= getNumberOfBestTimes(times[i], distances[i])
	}

	return fmt.Sprint(mul), nil
}

func part2(inputLines []string) (string, error) {
	timesStr := strings.Fields(inputLines[0])[1:]
	distancesStr := strings.Fields(inputLines[1])[1:]
	time, _ := strconv.Atoi(timesStr[0] + timesStr[1] + timesStr[2] + timesStr[3])
	distance, _ := strconv.Atoi(distancesStr[0] + distancesStr[1] + distancesStr[2] + distancesStr[3])

	answer := getNumberOfBestTimes(time, distance)

	return fmt.Sprint(answer), nil
}
