package _5Kyu

import "math"

/*
	Given a recipe, should return how many recipes it is possible to make with the ingredients available
*/

func Cakes(recipe, available map[string]int) int {
	min := math.MaxInt
	for material, need := range recipe {
		amount, ok := available[material]
		if !ok {
			return 0
		}
		cakesPossible := amount / need 
		if cakesPossible < min {
			min = cakesPossible
		} 
	}
	return min
}
