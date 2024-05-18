package Leetcode_go

import (
	"slices"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	slices.Sort(potions)
	ans := make([]int, len(spells))
	for i := 0; i < len(spells); i++ {
		y := findClosest(spells[i], potions, success)
		ans[i] = len(potions) - y - 1
	}
	return ans
}

func findClosest(spell int, potions []int, success int64) int {
	lo := 0
	hi := len(potions) - 1
	for lo < hi {
		mid := lo + (hi-lo)/2
		temp := potions[mid] * spell
		if int(success) <= temp {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return hi
}
