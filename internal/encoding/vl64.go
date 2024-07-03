// Copyright 2018 Ewout van Mansom. All rights reserved.
// Use of this source code is governed by a AGPLv3
// license that can be found in the LICENSE file.
package encoding

import (
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

// Length calculates the length of a VL64 sequence.
func Length(firstChar byte) int {
	return (int(firstChar) - int('@')) / 8
}

// Decode transforms the mixed radix representation into an integer.
func Decode(s []byte) (int, int) {
	len := Length(s[0])    // (firstChar - 64) / 8
	total := int(s[0]) % 4 // Base4 value
	inc := 0

	for inc < len { // Increment all b64 symbols to the total
		total += (int(s[inc]) - int(byte('@'))) * int((math.Pow(float64(64), float64(inc)))/16)
		inc++
	}

	positive := s[0]%8 < 4 // Base4 positive/negative
	if positive {
		return total, len
	}

	return -total, len
}
