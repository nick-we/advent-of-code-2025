package day01

import (
	"fmt"
	"strings"
)

func PasswordSolverPart2(input string) (int, error) {
	password := 0
	currentNumber := 50

	for line := range strings.FieldsSeq(input) {
		var dir rune
		var num int
		if _, err := fmt.Sscanf(line, "%c%d", &dir, &num); err != nil {
			return 0, fmt.Errorf("invalid input format: %s", line)
		}

		switch dir {
		case 'L':
			password += num / 100
			rem := num % 100
			if currentNumber > 0 && currentNumber-rem <= 0 {
				password++
			}
			currentNumber = (currentNumber - rem + 100) % 100
		case 'R':
			password += num / 100
			rem := num % 100
			if currentNumber+rem >= 100 {
				password++
			}
			currentNumber = (currentNumber + rem) % 100
		default:
			return 0, fmt.Errorf("invalid direction: %c", dir)
		}
	}

	return password, nil
}
