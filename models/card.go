package models

import (
	"unicode/utf8"
)

//Parse card takes a card string and returns an int array we can run LuhnCheck on
func ParseCard(card string) []int {
	arrlen := utf8.RuneCountInString(card)
	ret := make([]int, arrlen)
	for i, dig := range card {
		ret[i] = dig
	}
	return ret
}

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
