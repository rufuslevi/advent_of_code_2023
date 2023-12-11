package jour7

import "advent_of_code_2023/utils"

func Main() string {
	inputLines := utils.ReadFileLines("jour7/input.txt")
	inputLines = utils.ReadFileLines("jour7/test_input.txt")
	return utils.LogDayAnswers(inputLines, []func([]string) (string, error){part1, part2})
}

func part1(inputLines []string) (string, error) {
	return "", nil
}

func part2(inputLines []string) (string, error) {
	return "", nil
}
