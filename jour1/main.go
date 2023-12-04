package jour1

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var numbersNames = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

func Main() (string, error) {
	file, err := os.Open("jour1/input.txt")
	if err != nil {
		return "", errors.New("Coudln't read input file")
	}
	input := bufio.NewReader(file)

	sum := 0

	line, err := input.ReadBytes('\n')
	for string(line) != "" {
		lineNumbers := make(map[int]string)
		for name, number := range numbersNames {
			if strings.Contains(string(line), name) {
				lineNumbers[strings.Index(string(line), name)] = number
				lineNumbers[strings.LastIndex(string(line), name)] = number
			}
		}

		for pos, char := range string(line) {
			num, err := strconv.Atoi(string(char))
			if err == nil {
				lineNumbers[pos] = fmt.Sprint(num)
			}
		}

		var keys []int
		for pos := range lineNumbers {
			keys = append(keys, pos)
		}

		sort.Ints(keys)

		var values []string
		for _, pos := range keys {
			values = append(values, lineNumbers[pos])
		}

		lineNumber, err := strconv.Atoi(fmt.Sprintf("%s%s", values[0], values[len(values)-1]))
		if err != nil {
			return "", err
		}

		sum += lineNumber

		line, err = input.ReadBytes('\n')
	}

	return string(sum), nil
}
