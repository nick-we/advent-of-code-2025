package day02

import "testing"

func TestIsInvalidIdPart2(t *testing.T) {
	type given struct {
		idText string
	}
	type want struct {
		isInvalid bool
	}

	tests := []struct {
		name  string
		given given
		want  want
	}{
		{
			name: "id is valid",
			given: given{
				idText: "101",
			},
			want: want{
				isInvalid: false,
			},
		},
		{
			name: "id is invalid",
			given: given{
				idText: "1111",
			},
			want: want{
				isInvalid: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isInvalid := IsInvalidIdPart2(tt.given.idText)
			if isInvalid != tt.want.isInvalid {
				t.Errorf("want: %t, but got: %t", tt.want.isInvalid, isInvalid)
			}
		})
	}
}

func TestSumOfInvalidIdsPart2(t *testing.T) {
	type given struct {
		input string
	}
	type want struct {
		sum int
	}

	tests := []struct {
		name  string
		given given
		want  want
	}{
		{
			name: "test case 1",
			given: given{
				input: `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124`,
			},
			want: want{
				sum: 4174379265,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSum, err := SumOfInvalidIdsPart2(tt.given.input)
			if err != nil {
				t.Fatal(err)
			}
			if gotSum != tt.want.sum {
				t.Errorf("want: %d, but got: %d", tt.want.sum, gotSum)
			}
		})
	}
}
