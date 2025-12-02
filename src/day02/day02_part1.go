package day02

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func SumOfInvalidIds(input string) (int, error) {
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
			if IsInvalidId(strconv.Itoa(num)) {
				sum += num
			}
		}
	}

	return sum, nil
}

func IsInvalidId(idText string) bool {
	if len(idText)%2 == 1 {
		return false
	}

	lengthHalf := len(idText) / 2

	return idText[:lengthHalf] == idText[lengthHalf:]
}
