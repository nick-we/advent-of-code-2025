package day02

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func SumOfInvalidIdsPart2(input string) (int, error) {
	sum := 0

	for idRange := range strings.SplitSeq(input, ",") {
		ids := slices.Collect(strings.SplitSeq(idRange, "-"))
		if len(ids) != 2 {
			return 0, fmt.Errorf("invalid amount of ids (must be 2): %s", idRange)
		}

		id0, err := strconv.Atoi(strings.TrimSpace(ids[0]))
		if err != nil {
			return 0, fmt.Errorf("failed conversion to int: %s", ids[0])
		}
		id1, err := strconv.Atoi(strings.TrimSpace(ids[1]))
		if err != nil {
			return 0, fmt.Errorf("failed conversion to int: %s", ids[1])
		}

		for num := id0; num <= id1; num++ {
			if IsInvalidIdPart2(strconv.Itoa(num)) {
				sum += num
			}
		}
	}

	return sum, nil
}

func IsInvalidIdPart2(idText string) bool {
	isInvalid := false

	for i := 1; i <= len(idText)/2; i++ {
		if len(idText)%i != 0 {
			continue
		}

		subStr := idText[:i]
		isInvalid := true

		for j := i; j < len(idText); j += i {
			if subStr != idText[j:j+i] {
				isInvalid = false
				break
			}
		}
		if isInvalid {
			return true
		}
	}

	return isInvalid
}
