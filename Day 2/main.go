package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Red         = "red"
	Green       = "green"
	Blue        = "blue"
	RedAmount   = 12
	GreenAmount = 13
	BlueAmount  = 14
)

func main() {
	input, _ := os.ReadFile("./input")
	content := strings.TrimSpace(string(input))
	lines := strings.Split(content, "\n")
	fmt.Println("Part 1:", partOne(lines))
	fmt.Println("Part 2:", partTwo(lines))
}

func partOne(lines []string) int64 {
	var result int64 = 0

	for _, line := range lines {
		id, isPossible := processLinePartOne(line)

		if isPossible {
			result += id
		}
	}

	return result
}

func processLinePartOne(line string) (int64, bool) {
	id, _ := strconv.ParseInt(strings.Trim(strings.Split(line, " ")[1], ":"), 0, 64)
	plays := strings.Split(line, ": ")[1]
	sets := strings.Split(plays, "; ")

	for _, set := range sets {
		cubes := strings.Split(set, ", ")

		for _, cube := range cubes {
			chunks := strings.Split(cube, " ")
			color := chunks[1]
			n, _ := strconv.ParseInt(chunks[0], 0, 64)

			switch color {
			case Red:
				if n > RedAmount {
					return id, false
				}
			case Green:
				if n > GreenAmount {
					return id, false
				}
			case Blue:
				if n > BlueAmount {
					return id, false
				}
			}
		}
	}

	return id, true
}

func partTwo(lines []string) int64 {
	var result int64 = 0

	for _, line := range lines {
		result += processLinePartTwo(line)
	}

	return result
}

func processLinePartTwo(line string) int64 {
	plays := strings.Split(line, ": ")[1]
	sets := strings.Split(plays, "; ")
	var maxRed int64
	var maxGreen int64
	var maxBlue int64

	for _, set := range sets {
		cubes := strings.Split(set, ", ")

		for _, cube := range cubes {
			chunks := strings.Split(cube, " ")
			color := chunks[1]
			n, _ := strconv.ParseInt(chunks[0], 0, 64)

			switch color {
			case Red:
				if n > maxRed {
					maxRed = n
				}
			case Green:
				if n > maxGreen {
					maxGreen = n
				}
			case Blue:
				if n > maxBlue {
					maxBlue = n
				}
			}
		}
	}

	return maxRed * maxBlue * maxGreen
}
