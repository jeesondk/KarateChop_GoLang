package KarateChop

import "fmt"

var _pos int = 0
var _break bool = false
var _source []int
var _chop int

// ChopChop BinarySearch implementation
func ChopChop(input []int, chop int) int {

	fmt.Println("Got Array: %v, chopping for: %v", input, chop)
	if !isValid(input) {
		return -1
	}

	_source = input
	_chop = chop
	_pos = chopPos(len(_source))

	hasMatch := doChop(_source)
	if hasMatch {
		return _pos
	} else {
		return -1
	}
}

func doChop(input []int) bool {

	chop := chopPos(len(input))
	valAtChop := input[chop]

	if valAtChop == _chop {
		_break = true
	}

	if len(input) <= 1 {
		_break = true
	}

	for !_break {
		var intSlice = make([]int, len(input))
		if valAtChop <= _chop {
			intSlice = chopHigh(input)
		} else {
			intSlice = chopLow(input)
		}
		doChop(intSlice)
	}

	return checkValue()
}

func chopPos(arrLen int) int {
	return arrLen / 2
}

func chopLow([]int) []int {
	return []int{}
}

func chopHigh([]int) []int {
	return []int{}
}

func checkValue() bool {
	return _source[_pos] == _chop
}

func isValid(input []int) bool {
	if len(input) == 0 {
		return false
	} else {
		return true
	}
}
