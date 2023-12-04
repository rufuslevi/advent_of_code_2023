package utils

import (
	"bufio"
	"io"
	"os"
)

func ReadFileLines(filePath string) []string {
	file, err := os.Open(filePath)
	CheckErr(err, "Couldn't read input file")

	input := bufio.NewReader(file)

	lines := []string{}

	line, err := input.ReadBytes('\n')
	CheckErr(err, "Error reading line")
	for string(line) != "" {
		lines = append(lines, string(line))

		line, err = input.ReadBytes('\n')
		if err != io.EOF {
			CheckErr(err, "")
		}
	}

	return lines
}
