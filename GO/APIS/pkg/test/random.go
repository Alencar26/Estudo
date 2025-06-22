package test

import (
	"math/rand/v2"
	"time"
)

func RandomLetter(amount int) string {
	var letters string
	for i := 1; i <= amount; i++ {
		x := time.Now().UnixNano()
		if x%2 == 0 {

			min := 65
			max := 90

			letters += ReturnASCII(rand.IntN(max-min+1) + min)

			continue
		}
		min := 97
		max := 122

		letters += ReturnASCII(rand.IntN(max-min+1) + min)

	}
	return letters
}

func ReturnASCII(n int) string {
	return string(rune(n))
}
