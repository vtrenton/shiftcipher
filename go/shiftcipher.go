package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const MAXLEN int = 26

type RingBuffer struct {
	lower [MAXLEN]rune
	upper [MAXLEN]rune
}

func (rb *RingBuffer) populate() {
	for i := 0; i < MAXLEN; i++ {
		rb.lower[i] = rune('a' + i)
	}

	for i := 0; i < MAXLEN; i++ {
		rb.upper[i] = rune('A' + i)
	}
}

func main() {

	var input string
	var shift int
	var ciphertext string
	var rb RingBuffer

	if len(os.Args) == 3 {
		input = os.Args[1]
		shift_64, err := strconv.ParseInt(os.Args[2], 10, 32)
		if err != nil {
			log.Fatalf("Something bad happened")
		}
		shift = int(shift_64)
	} else {
		log.Fatal("ya goofd - enter 1. a string to shift and 2. number to shift by")
	}

	rb.populate()

	for _, char := range input {
		ciphertext += string(encodeChar(char, shift, rb))
	}

	fmt.Println(ciphertext)
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
	log.Fatalf("Character %c not found in array - this should be impossible", char)
	return -1 // never reached, just to satisfy compiler
}
