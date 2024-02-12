#!/usr/bin/env python

################
#
# (c) Copyright - Trenton VanderWert
# This is Free Software licensed under the GNU General Public License (GPL) V3
#
# PYTHON CAESAR CIPHER ENCODER/DECODER
# 
# Caesar Ciphars are a cryptographically weak and very old cipher used for obscuring text. Used by Julius Caesar to send messages to the millitary about battle plans. In order to decode the recipient must know the ammount of characters shifted and in which direction. Since the Alphabet is only 26 characters this method is VERY easy to brute force by just trial and error.
#
# A Traditional Caesar cipher works like this:
# say you'd like to send the message: "attack the north wall"
# you could shift every character in this message a certain amount up or down the alphabet
# for this example I'll say the shift is 7
# so the first letter is "a" would shift 7 places up (or down) the alphabet
# UP: a(0) b(1) c(2) d(3) e(4) f(5) g(6) h(7)
# DOWN: a(0) z(-1) y(-2) x(-3) w(-4) v(-5) u(-6) t(-7)
# 
# So our first letter is either a+7=h or a-7=t
# repeat this for all the characters in the message and you have your ciphertext
#
# FUN FACT: a 13 shift is the same + and - since it is half the alphabet this is also called a rot13
#
###############


# take user input
# need to use raw_input for the user string so we dont append the new line into the array
# using a try/catch to validate input and loop it so program doesnt crash
while True:
    try:
        user_sting = raw_input("Please enter a string you'd like to shift: ")
        caesar_shift = int(input("How many chars would you like to shift: "))
        if caesar_shift > 26 or caesar_shift < 0:
            print "There are a max of 26 chars in the Alphabet! Enter a value between 0 and 26"
        else:
            break
    except:
        print "invalid input!\n"

# split user input into array
input_array = list(user_sting)

# Define some output arrays to hold output chars
pos_array = []
neg_array = []

# loop through the array of user input chars
for x in input_array:
    # bring all uppercase chars to lower for consistancy
    i = x.lower()
    # convert the char to the decimal equal for math
    dec_i = ord(i)
    # if statment to skip over non letters
    if dec_i >= 97 and dec_i <= 122:
        # This is where the magic happens
        # since we converted the user input into a decimal above we can add or subtract our shift 
        pos_int = dec_i + caesar_shift
        # check to see if our number goes out of range and wrap around if it does 
        if pos_int > 122:
            pos_int = pos_int - 122 + 97 - 1
        # repeat same method for negative
        # just slighly different math for wrapping
        neg_int = dec_i - caesar_shift
        if neg_int < 97:
            neg_int = 122 - (97 - neg_int) + 1
        
        # take the number and convert it back into a char
        # then add it to the output array
        pos_array.append(chr(pos_int))
        neg_array.append(chr(neg_int))
    else:
        # char was not a letter - add the char as-is and move on
        pos_array.append(x)
        neg_array.append(x)


# print the cipher text out to the user
# use the join function to concatenate all of the chars in the array
print "+" + str(caesar_shift)
print ''.join(pos_array)
print "\n"
print "-" + str(caesar_shift)
print ''.join(neg_array)
print "\n"
print "PLEASE NOTE: All chars were shifted to lowercase during the conversion please validate case!!"
