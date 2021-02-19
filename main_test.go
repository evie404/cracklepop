package main

import "testing"

func Test_cracklePop(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want string
	}{
		{
			"returns CracklePop for numbers divisible by both 3 and 5",
			[]int{
				0, 15, 30, 45, 60, 90, 105,
			},
			"CracklePop",
		},
		{
			"returns Crackle for numbers divisible by 3 but not 5",
			[]int{
				3, 6, 9, 12, 18, 21,
			},
			"Crackle",
		},
		{
			"returns Pop for numbers divisible by 5 but not 3",
			[]int{
				5, 10, 20, 25, 35, 40, 50,
			},
			"Pop",
		},
		{
			"returns string representation for numbers not divisible by 3 or 5",
			[]int{
				1,
			},
			"1",
		},
		{
			"returns string representation for numbers not divisible by 3 or 5",
			[]int{
				7,
			},
			"7",
		},
		{
			"returns string representation for numbers not divisible by 3 or 5",
			[]int{
				91,
			},
			"91",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, arg := range tt.args {
				if got := cracklePop(arg); got != tt.want {
					t.Errorf("cracklePop() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
