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
			ticks := num / 100
			password += ticks

			num %= 100
			if num != 0 {
				if currentNumber != 0 {
					currentNumber -= num
					if currentNumber < 0 {
						currentNumber += 100
						password++
					} else if currentNumber == 0 {
						password++
					}
				} else {
					currentNumber -= num
					if currentNumber < 0 {
						currentNumber += 100
					}
				}
			}
		case 'R':
			ticks := num / 100
			password += ticks

			num %= 100
			if num != 0 {
				if currentNumber != 0 {
					currentNumber += num
					if currentNumber >= 100 {
						currentNumber -= 100
						password++
					}
				} else {
					currentNumber += num
				}
			}
		default:
			return 0, fmt.Errorf("invalid direction: %c", dir)
		}
	}

	return password, nil
}
