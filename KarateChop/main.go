package KarateChop

import "fmt"

var _pos int = 0
var _break bool = false
var _source []int
var _chop int

// ChopChop BinarySearch implementation
func ChopChop(input []int, chop int) int {

	fmt.Println("Got Array: %v, chopping for: %v", input, chop)
	if isValid(input) {
		return -1
	}

	_source = input
	_chop = chop

	return -1
}

func chop(input []int) bool {
	return false
}

func chopLow() []int {
	return []int{}
}

func chopHigh() []int {
	return []int{}
}

func resultFound() bool {
	return false
}

func isValid(input []int) bool {
	if input != nil {
		return false
	} else {
		return true
	}
}