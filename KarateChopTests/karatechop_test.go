package KarateChopTests

import (
	"KarateChop_GoLang/KarateChop"
	"fmt"
	"testing"
)

func TestKarateChopping(t *testing.T) {
	tests := []struct {
		input    []int
		chop     int
		expected int
	}{
		{[]int{}, 3, -1},
		{[]int{1}, 3, -1},
		{[]int{1}, 1, 0},

		{[]int{1, 3, 5}, 1, 0},
		{[]int{1, 3, 5}, 3, 1},
		{[]int{1, 3, 5}, 5, 2},
		{[]int{1, 3, 5}, 0, -1},
		{[]int{1, 3, 5}, 2, -1},
		{[]int{1, 3, 5}, 4, -1},
		{[]int{1, 3, 5}, 6, -1},

		{[]int{1, 3, 5, 7}, 1, 0},
		{[]int{1, 3, 5, 7}, 3, 1},
		{[]int{1, 3, 5, 7}, 5, 2},
		{[]int{1, 3, 5, 7}, 7, 3},
		{[]int{1, 3, 5, 7}, 0, -1},
		{[]int{1, 3, 5, 7}, 2, -1},
		{[]int{1, 3, 5, 7}, 4, -1},
		{[]int{1, 3, 5, 7}, 6, -1},
		{[]int{1, 3, 5, 7}, 8, -1},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("KarateChop=%d", i), func(t *testing.T) {
			got := KarateChop.ChopChop(tc.input, tc.chop)
			if got != tc.expected {
				t.Fatalf("got %v; expected %v", got, tc.expected)
			} else {
				t.Logf("Passed !")
			}
		})
	}
}
