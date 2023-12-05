package jour4

import (
	"advent_of_code_2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Main() string {
	inputLines := utils.ReadFileLines("jour4/input.txt")
	return utils.LogDayAnswers(inputLines, []func([]string) (string, error){part1, part2})
}

func part1(inputLines []string) (string, error) {
	sum := 0

	for _, line := range inputLines {
		lineParts := strings.Split(line[10:], "|")

		var winningNumbers []int
		for _, numberString := range strings.Split(lineParts[0], " ") {
			if numberString == "" || numberString == "0" {
				continue
			}
			number, _ := strconv.Atoi(numberString)

			winningNumbers = append(winningNumbers, number)
		}

		offset := -1
		for _, numberString := range strings.Split(lineParts[1][:len(lineParts[1])-1], " ") {
			if numberString == "" || numberString == "0" {
				continue
			}
			number, _ := strconv.Atoi(numberString)

			if slices.Contains(winningNumbers, number) {
				offset++
			}
		}

		if offset >= 0 {
			sum += 1 << offset
		}
	}

	return fmt.Sprint(sum), nil
}

func part2(inputLines []string) (string, error) {
	return "", nil
}
