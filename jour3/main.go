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
		// println(thisNumber)
	}

	if pos-1 >= 0 && unicode.IsDigit(line[pos-1]) && !slices.Contains(visited[pos-1], lineNumber) {
		thisNumber = fmt.Sprintf("%s%s", getNumber(line, lineNumber, pos-1, visited), thisNumber)
	}
	if pos+1 < len(line) && unicode.IsDigit(line[pos+1]) && !slices.Contains(visited[pos+1], lineNumber) {
		thisNumber = fmt.Sprintf("%s%s", thisNumber, getNumber(line, lineNumber, pos+1, visited))
	}

	return string(thisNumber)
}

func Main() string {
	inputLines := utils.ReadFileLines("jour3/input.txt")
	return utils.LogDayAnswers(inputLines, []func([]string) (string, error){part1})
}

func checkCoord(x int, y int, inputLines []string, visited map[int][]int, sum *int) {
	number := getNumber([]rune(inputLines[y]), y, x, visited)
	newPartNumber, err := strconv.Atoi(number)
	if err != nil {
		println(err.Error())
		utils.CheckErr(err, "Error reading the number")
		return
	}
	// fmt.Printf("adding %d\n", newPartNumber)
	*sum += newPartNumber
}

func part1(inputLines []string) (string, error) {
	sum := 0
	visited := make(map[int][]int)

	for y, line := range inputLines {
		for x, char := range line {
			if isSymbol(char) && char != '\n' {
				fmt.Printf("char at %d %d is a %s\n", x, y, string(char))
				if x-1 >= 0 && y-1 >= 0 && !slices.Contains(visited[x-1], y-1) {
					checkCoord(x-1, y-1, inputLines, visited, &sum)
				}
				if y-1 >= 0 && !slices.Contains(visited[x], y-1) {
					checkCoord(x, y-1, inputLines, visited, &sum)
				}
				if x+1 < len(line) && y-1 >= 0 && !slices.Contains(visited[x+1], y-1) {
					checkCoord(x+1, y-1, inputLines, visited, &sum)
				}
				if x+1 < len(line) && !slices.Contains(visited[x+1], y) {
					checkCoord(x+1, y, inputLines, visited, &sum)
				}
				if x+1 < len(line) && y+1 < len(inputLines) && !slices.Contains(visited[x+1], y+1) {
					checkCoord(x+1, y+1, inputLines, visited, &sum)
				}
				if y+1 < len(inputLines) && !slices.Contains(visited[x], y+1) {
					checkCoord(x, y+1, inputLines, visited, &sum)
				}
				if x-1 > 0 && y+1 < len(inputLines) && !slices.Contains(visited[x-1], y+1) {
					checkCoord(x-1, y+1, inputLines, visited, &sum)
				}
				if x-1 > 0 && !slices.Contains(visited[x-1], y) {
					checkCoord(x-1, y, inputLines, visited, &sum)
				}
			}
		}
	}

	return fmt.Sprint(sum), nil
}
