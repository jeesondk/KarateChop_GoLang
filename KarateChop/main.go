package KarateChop

import "fmt"

func main() {
	data := []struct {
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

	for i, d := range data {
		fmt.Sprintf("KarateChop=%d", i)
		res := ChopChop(d.input, d.chop)
		fmt.Sprintf("Got %v; expected %v", res, d.expected)
	}
}

// ChopChop BinarySearch implementation
func ChopChop(input []int, chop int) int {

	fmt.Printf("Got Array: %v, chopping for: %v\n", input, chop)

	if !isValid(&input) {
		return -1
	} else {
		return 0
	}
}

func isValid(i *[]int) bool {
	if len(*i) < 1 {
		return false
	} else {
		return true
	}
}
