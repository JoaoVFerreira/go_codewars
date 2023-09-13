package _8Kyu

var Sheeps = []bool{
	true, true, true, false,
	true, true, true, true,
	true, false, true, false,
	true, false, false, true,
	true, true, true, true,
	false, false, true, true,
}

/*
	Receive sheeps (True) and should return the qunatity of sheeps
*/

func CountSheeps(batchSheeps []bool) int {
	sheeps := 0
	for _, sheep := range batchSheeps {
		if sheep {
			sheeps++
		}
	}
	return sheeps
}
