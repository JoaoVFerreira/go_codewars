package _6Kyu

import "testing"

func TestTwoSum(t *testing.T) {
	got := TwoSum([]int{1, 2, 3}, 4)
	want := [2]int{0, 2}

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("want %q got %q", want, got)
	}
}