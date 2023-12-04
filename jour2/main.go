package jour2

import (
	"advent_of_code_2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func Main() string {
	inputLines := utils.ReadFileLines("jour2/input.txt")
	return utils.LogDayAnswers(inputLines, []func([]string) (string, error){part1, part2})
}

func part1(inputLines []string) (string, error) {
	sum := 0

	for _, line := range inputLines {
		colors := map[string]int{"red": 0, "green": 0, "blue": 0}
		words := strings.Split(line, " ")

		gameId, _ := strconv.Atoi(words[1][:len(words[1])-1])

		for i := 2; i < len(words); i += 2 {
			number, err := strconv.Atoi(words[i])
			utils.CheckErr(err, "Couldn't find the number of cubes")

			if (colors[words[i+1][:len(words[i+1])-1]]) < number {
				colors[words[i+1][:len(words[i+1])-1]] = number
			}
		}

		if colors["red"] <= 12 && colors["green"] <= 13 && colors["blue"] <= 14 {
			sum += gameId
		}
	}

	return fmt.Sprint(sum), nil
}

func part2(inputLines []string) (string, error) {
	sum := 0

	for _, line := range inputLines {
		colors := map[string]int{"red": 0, "green": 0, "blue": 0}
		words := strings.Split(line, " ")

		for i := 2; i < len(words); i += 2 {
			number, err := strconv.Atoi(words[i])
			utils.CheckErr(err, "Couldn't find the number of cubes")

			if (colors[words[i+1][:len(words[i+1])-1]]) < number {
				colors[words[i+1][:len(words[i+1])-1]] = number
			}
		}

		sum += colors["red"] * colors["green"] * colors["blue"]
	}

	return fmt.Sprint(sum), nil
}
