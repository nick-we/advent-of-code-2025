package day01

import (
	"fmt"
	"strings"
)

func PasswordSolverPart1(input string) (int, error) {
	password := 0
	currentNumber := 50

	for line := range strings.FieldsSeq(input) {
		var dir rune
		var num int
		if _, err := fmt.Sscanf(line, "%c%d", &dir, &num); err != nil {
			return 0, fmt.Errorf("invalid input format: %s", line)
		}

		num %= 100
		switch dir {
		case 'L':
			currentNumber = (currentNumber - num + 100) % 100
		case 'R':
			currentNumber = (currentNumber + num) % 100
		default:
			return 0, fmt.Errorf("invalid direction: %c", dir)
		}

		if currentNumber == 0 {
			password++
		}
	}

	return password, nil
}
