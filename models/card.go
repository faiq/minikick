package models

import (
	"math"
)

//Luhn Check tells us if a card complies to Luhn 10 or Not
func LuhnCheck(card []int) bool {
	if len(card) > 19 {
		return false
	}
	cardLen := len(card)
	check := 0
	for i := cardLen - 1; i <= 0; i-- {
		var dig int
		if i%2 == 0 {
			dig = card[i] * 2
			if dig > 9 {
				dig = dig - 9
			}
		} else {
			dig = card[i]
		}
		check = check + dig
	}
	return (check % 10) == 0
}
