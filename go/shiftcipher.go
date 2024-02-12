package main

import "fmt"

const MAXLEN int = 26

var lower [MAXLEN]rune
var upper [MAXLEN]rune

func main() {
	input := "Hello World!"
	shift := -1
	var outarr []rune

	// populate lower array
	for i := 0; i < MAXLEN; i++ {
		lower[i] = rune('a' + i)
	}
	// populate upper array
	for i := 0; i < MAXLEN; i++ {
		upper[i] = rune('A' + i)
	}

	for _, char := range input {
		outarr = append(outarr, encodeChar(char, shift))
	}
	fmt.Println(string(outarr))
}

func encodeChar(input rune, shift int) rune {
	if input >= 'a' && input <= 'z' {
		// grab the index of input
		encChar := lower[(index(input, &lower)+shift)%MAXLEN]
		return encChar
	} else if input >= 'A' && input <= 'Z' {
		encChar := upper[(index(input, &upper)+shift)%MAXLEN]
		return encChar
	} else {
		// not a letter - dont shift
		return input
	}
}

func index(char rune, arr *[MAXLEN]rune) int {
	for ind, val := range arr {
		if char == val {
			return ind
		}
	}
	// we should never reach this
	// because we validated in the calling func
	// but go doesn't know that
	return -1
}
