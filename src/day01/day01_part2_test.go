package day01

import "testing"

func TestPasswordSolverPart2(t *testing.T) {
	input := `
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	password, err := PasswordSolverPart2(input)
	if err != nil {
		t.Fatal(err)
	}

	if password != 6 {
		t.Errorf("expected password to be 6, got %d", password)
	}
}

func TestPasswordSolverPart2RealInput(t *testing.T) {
	if _, err := PasswordSolverPart2(InputPart2); err != nil {
		t.Fatal(err)
	}
}
