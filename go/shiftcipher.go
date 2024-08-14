package main

import (
	"fmt"
	"os"
	"strconv"
)

const MAXLEN int = 26

type RingBuffer struct {
	lower [MAXLEN]rune
	upper [MAXLEN]rune
}

func main() {
	var input string
	var shift int
	if len(os.Args) == 3 {
		input = os.Args[1]

		shift_64, err := strconv.ParseInt(os.Args[2], 10, 32)
		shift = int(shift_64)
		// shits gonna crash if you try to strconv anything but an int
		// but go is pretty memory safe and this should catch it.
		if err != nil {
			fmt.Println("Something bad happened")
			os.Exit(1)
		}
	} else {
		input = "ya goofd - enter 1. a string to shift and 2. number to shift by"
		shift = 0
	}
	var ciphertext string
	var rb RingBuffer
	rb = populate(rb)

	for _, char := range input {
		ciphertext += string(encodeChar(char, shift, rb))
	}

	fmt.Println(ciphertext)
}

func populate(rb RingBuffer) RingBuffer {
	for i := 0; i < MAXLEN; i++ {
		rb.lower[i] = rune('a' + i)
	}

	for i := 0; i < MAXLEN; i++ {
		rb.upper[i] = rune('A' + i)
	}

	return RingBuffer{
		lower: rb.lower,
		upper: rb.upper,
	}
}

func encodeChar(input rune, shift int, rb RingBuffer) rune {
	var buff [MAXLEN]rune
	if input >= 'a' && input <= 'z' {
		buff = rb.lower
	} else if input >= 'A' && input <= 'Z' {
		buff = rb.upper
	} else {
		// not a letter - dont shift
		return input
	}

	getInd := index(input, &buff)

	getShift := (getInd + shift) % MAXLEN
	if getShift < 0 {
		getShift = getShift + MAXLEN
	}

	return buff[getShift]
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
