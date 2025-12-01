package main

import (
	"fmt"

	"github.com/nick-we/advent-of-code-2025/src/day01"
)

func main() {
	password, err := day01.PasswordSolverPart2(day01.InputPart2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Password: %d\n", password)
}
