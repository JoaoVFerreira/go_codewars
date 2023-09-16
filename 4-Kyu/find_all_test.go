package _4Kyu

import "testing"

func TestFindAll(t *testing.T) {
	got := FindAll(10, 3)
	want := []int{19, 109, 280}

	for index, value := range got {
		if value != want[index] {
			t.Errorf("want %q got %q", want[index], value)
		}
	}
}