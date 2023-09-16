package _4Kyu

import (
	"strconv"
)

func FindAll(sumDigit, digit int) []int {
	maxNumber := (digit * 100) - 1
	var partialResult []int
	for i := 1; i <= maxNumber; i++ {
		if len(strconv.Itoa(i)) == digit {
			digitSum := sumOfDigits(i)
			if digitSum == sumDigit {
				partialResult = append(partialResult, i)
			}
		}
	}
	var result []int
	lowest := maxNumber
	greatest := 0
	for _, v := range partialResult {
		if v > greatest {
			greatest = v
		}
		if v < lowest {
			lowest = v
		}
	}
	result = append(result, len(partialResult))
	result = append(result, lowest)
	result = append(result, greatest)
	return result
}

func sumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		rigthSideDigit := n % 10
		sum += rigthSideDigit
		n /= 10
	}
	return sum
}