package day01

import "testing"

func TestPasswordSolver(t *testing.T) {
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

	password, err := PasswordSolverPart1(input)
	if err != nil {
		t.Fatal(err)
	}

	if password != 3 {
		t.Errorf("expected password to be 3, got %d", password)
	}
}

func TestPasswordSolverPart1RealInput(t *testing.T) {
	if _, err := PasswordSolverPart1(InputPart1); err != nil {
		t.Fatal(err)
	}
}
