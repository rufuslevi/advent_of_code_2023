package jour1

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Main() (int, error) {
	file, err := os.Open("jour1/input.txt")
	if err != nil {
		return -1, errors.New("Coudln't read input file")
	}
	input := bufio.NewReader(file)

	sum := 0

	line, err := input.ReadBytes('\n')
	for string(line) != "" {
		var lineNumbers []int
		for _, char := range string(line) {
			num, err := strconv.Atoi(string(char))
			if err == nil {
				lineNumbers = append(lineNumbers, num)
			}
		}
		lineNumber, err := strconv.Atoi(fmt.Sprintf("%d%d", lineNumbers[0], lineNumbers[len(lineNumbers)-1]))
		if err != nil {
			return -1, err
		}

		sum += lineNumber

		line, err = input.ReadBytes('\n')
	}

	return sum, nil
}
