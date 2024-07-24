// Copyright 2018 Ewout van Mansom. All rights reserved.
package vl64

import (
	"errors"
	"fmt"
	"math"
)

const (
	// Negative constant is a char code for negative representation.
	Negative int = 72

	// Positive constant is a char code for positive representation.
	Positive int = 73

	// MaxInteger as 32-bit integer causes VL64 to have max length of 6.
	MaxInteger int = 6
)

// Encode transforms integer i into a mixed radix representation.
func Encode(i int) []byte {
	vl64 := make([]byte, 6)          // 32-bit integer causes VL64 to have max length of 6.
	num := int(math.Abs(float64(i))) // Operate on normalized, positive integer
	len := 1                         // Length indicator, updated during encode
	j := 1
	var indicator int
	if i < 0 {
		indicator = 'D'
	} else {
		indicator = '@'
	}

	vl64[0] = byte(num%4 + indicator) // Base4 char, positive(+64)/negative(+68) indicator
	num /= 4                          // Base4 processed, prepare for remaining b64 symbols

	for j < 6 {
		vl64[j] = byte(num%64 + '@') // b64
		num /= 64
		if vl64[j] != '@' {
			len = j + 1 // @ = padding / zero symbol
		}
		j++
	}

	vl64[0] = byte(int(vl64[0]) + len*8)
	return vl64[:len]
}

func Length(firstChar byte) int {
	return (int(firstChar) - int('@')) / 8
}

func Decode(s []byte) (int, int, error) {
	if len(s) == 0 {
		return 0, 0, errors.New("len(s) is 0")
	}

	vlen := Length(s[0])   // (firstChar - 64) / 8
	total := int(s[0]) % 4 // Base4 value
	inc := 0

	if vlen > len(s) {
		return 0, 0, fmt.Errorf("len(buf) = %d smaller than vl64 length %d", len(s), vlen)
	}

	for inc < vlen { // Increment all b64 symbols to the total
		total += (int(s[inc]) - int(byte('@'))) * int((math.Pow(float64(64), float64(inc)))/16)
		inc++
	}

	positive := s[0]%8 < 4 // Base4 positive/negative
	if positive {
		return total, vlen, nil
	}

	return -total, vlen, nil
}
