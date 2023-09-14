package _6Kyu

/*
	Given an array of numbers, the function should return the positions (indices) of the elements that, when summed, are equal to the target
*/

func TwoSum(numbers []int, target int) [2]int {
	var result [2]int
	for j, number := range numbers {
		for i := j + 1; i < len(numbers); i++ {
			if number+numbers[i] == target {
				result[0] = j
				result[1] = i
				return result
			}
		}
	}
	if result[0] > result[1] {
		result[0], result[1] = result[1], result[0]
	}
	return result
}
