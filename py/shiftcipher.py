#!/usr/bin/env python3
import string

CAPACITY = 26

input = "attack the north wall"
shift = -512
outarr = []


lower = list(string.ascii_lowercase)
upper = list(string.ascii_uppercase)

def encodeChar(char):
    dec = ord(char)
    # lowercase
    if dec >= 97 and dec <= 122:
        # get index of current letter
        index = lower.index(char)
        encChar = lower[(index + shift)%CAPACITY]
        return encChar
    # uppercase
    elif dec >= 65 and dec <= 96:
        index = upper.index(char)
        encChar = upper[(index + shift)%CAPACITY]
        return encChar
    else:
        return char


for c in input:
    outarr.append(encodeChar(c))

encMsg = ''.join(outarr)
print(encMsg)
