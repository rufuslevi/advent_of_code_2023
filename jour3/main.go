package jour3

import (
	"advent_of_code_2023/utils"
	"fmt"
	"slices"
	"strconv"
	"unicode"
)

func isSymbol(input rune) bool {
	return !unicode.IsDigit(input) && !unicode.IsNumber(input) && !unicode.IsLetter(input) && input != '.'
}

func getNumber(line []rune, lineNumber int, pos int, visited map[int][]int) string {
	if !unicode.IsDigit(line[pos]) {
		return "0"
	}

	thisNumber := ""
	if !slices.Contains(visited[pos], lineNumber) {
		visited[pos] = append(visited[pos], lineNumber)
		thisNumber = string(line[pos])
	}

	if pos-1 >= 0 && unicode.IsDigit(line[pos-1]) && !slices.Contains(visited[pos-1], lineNumber) {
		thisNumber = getNumber(line, lineNumber, pos-1, visited) + thisNumber
	}
	if pos+1 < len(line) && unicode.IsDigit(line[pos+1]) && !slices.Contains(visited[pos+1], lineNumber) {
		thisNumber = thisNumber + getNumber(line, lineNumber, pos+1, visited)
	}

	return string(thisNumber)
}

func getNumberLoop(line []rune, lineNumber int, pos int, visited map[int][]int) string {
	if !unicode.IsDigit(line[pos]) {
		return "0"
	}

	thisNumber := ""
	if !slices.Contains(visited[pos], lineNumber) {
		visited[pos] = append(visited[pos], lineNumber)
		thisNumber = string(line[pos])
	}

	offset := 1
	for pos-offset >= 0 && !slices.Contains(visited[pos-offset], lineNumber) && unicode.IsDigit(line[pos-offset]) {
		visited[pos-offset] = append(visited[pos-offset], lineNumber)
		thisNumber = string(line[pos-offset]) + thisNumber
		offset++
	}
	offset = 1
	for pos+offset < len(line) && !slices.Contains(visited[pos+offset], lineNumber) && unicode.IsDigit(line[pos+offset]) {
		visited[pos+offset] = append(visited[pos+offset], lineNumber)
		thisNumber = thisNumber + string(line[pos+offset])
		offset++
	}

	return string(thisNumber)
}

func checkCoord(x int, y int, inputLines []string, visited map[int][]int) int {
	number := getNumberLoop([]rune(inputLines[y]), y, x, visited)
	newPartNumber, err := strconv.Atoi(number)

	if err != nil {
		utils.CheckErr(err, "Error reading the number")
		return -1
	}
	if newPartNumber == 0 {
		return -1
	}

	return newPartNumber
}

func Main() string {
	inputLines := utils.ReadFileLines("jour3/input.txt")
	return utils.LogDayAnswers(inputLines, []func([]string) (string, error){part1, part2})
}

func part1(inputLines []string) (string, error) {
	sum := 0
	visited := make(map[int][]int)

	for y, line := range inputLines {
		for x, char := range line {
			if isSymbol(char) && char != '\n' {
				if x-1 >= 0 && y-1 >= 0 && !slices.Contains(visited[x-1], y-1) {
					newPartNumber := checkCoord(x-1, y-1, inputLines, visited)
					if newPartNumber != -1 {
						sum += newPartNumber
					}
				}
				if y-1 >= 0 && !slices.Contains(visited[x], y-1) {
					newPartNumber := checkCoord(x, y-1, inputLines, visited)
					if newPartNumber != -1 {
						sum += newPartNumber
					}
				}
				if x+1 < len(line) && y-1 >= 0 && !slices.Contains(visited[x+1], y-1) {
					newPartNumber := checkCoord(x+1, y-1, inputLines, visited)
					if newPartNumber != -1 {
						sum += newPartNumber
					}
				}
				if x+1 < len(line) && !slices.Contains(visited[x+1], y) {
					newPartNumber := checkCoord(x+1, y, inputLines, visited)
					if newPartNumber != -1 {
						sum += newPartNumber
					}
				}
				if x+1 < len(line) && y+1 < len(inputLines) && !slices.Contains(visited[x+1], y+1) {
					newPartNumber := checkCoord(x+1, y+1, inputLines, visited)
					if newPartNumber != -1 {
						sum += newPartNumber
					}
				}
				if y+1 < len(inputLines) && !slices.Contains(visited[x], y+1) {
					newPartNumber := checkCoord(x, y+1, inputLines, visited)
					if newPartNumber != -1 {
						sum += newPartNumber
					}
				}
				if x-1 > 0 && y+1 < len(inputLines) && !slices.Contains(visited[x-1], y+1) {
					newPartNumber := checkCoord(x-1, y+1, inputLines, visited)
					if newPartNumber != -1 {
						sum += newPartNumber
					}
				}
				if x-1 > 0 && !slices.Contains(visited[x-1], y) {
					newPartNumber := checkCoord(x-1, y, inputLines, visited)
					if newPartNumber != -1 {
						sum += newPartNumber
					}
				}
			}
		}
	}

	return fmt.Sprint(sum), nil
}

func part2(inputLines []string) (string, error) {
	sum := 0
	visited := make(map[int][]int)
	gears := make(map[int][]int)

	symbolId := 0
	for y, line := range inputLines {
		for x, char := range line {
			if isSymbol(char) && char != '\n' {
				if x-1 >= 0 && y-1 >= 0 && !slices.Contains(visited[x-1], y-1) {
					newPartNumber := checkCoord(x-1, y-1, inputLines, visited)
					if newPartNumber != -1 {
						gears[symbolId] = append(gears[symbolId], newPartNumber)
					}
				}
				if y-1 >= 0 && !slices.Contains(visited[x], y-1) {
					newPartNumber := checkCoord(x, y-1, inputLines, visited)
					if newPartNumber != -1 {
						gears[symbolId] = append(gears[symbolId], newPartNumber)
					}
				}
				if x+1 < len(line) && y-1 >= 0 && !slices.Contains(visited[x+1], y-1) {
					newPartNumber := checkCoord(x+1, y-1, inputLines, visited)
					if newPartNumber != -1 {
						gears[symbolId] = append(gears[symbolId], newPartNumber)
					}
				}
				if x+1 < len(line) && !slices.Contains(visited[x+1], y) {
					newPartNumber := checkCoord(x+1, y, inputLines, visited)
					if newPartNumber != -1 {
						gears[symbolId] = append(gears[symbolId], newPartNumber)
					}
				}
				if x+1 < len(line) && y+1 < len(inputLines) && !slices.Contains(visited[x+1], y+1) {
					newPartNumber := checkCoord(x+1, y+1, inputLines, visited)
					if newPartNumber != -1 {
						gears[symbolId] = append(gears[symbolId], newPartNumber)
					}
				}
				if y+1 < len(inputLines) && !slices.Contains(visited[x], y+1) {
					newPartNumber := checkCoord(x, y+1, inputLines, visited)
					if newPartNumber != -1 {
						gears[symbolId] = append(gears[symbolId], newPartNumber)
					}
				}
				if x-1 > 0 && y+1 < len(inputLines) && !slices.Contains(visited[x-1], y+1) {
					newPartNumber := checkCoord(x-1, y+1, inputLines, visited)
					if newPartNumber != -1 {
						gears[symbolId] = append(gears[symbolId], newPartNumber)
					}
				}
				if x-1 > 0 && !slices.Contains(visited[x-1], y) {
					newPartNumber := checkCoord(x-1, y, inputLines, visited)
					if newPartNumber != -1 {
						gears[symbolId] = append(gears[symbolId], newPartNumber)
					}
				}
				symbolId += 1
			}
		}
	}

	for _, ratios := range gears {
		if len(ratios) == 2 {
			sum += ratios[0] * ratios[1]
		}
	}

	return fmt.Sprint(sum), nil

}
