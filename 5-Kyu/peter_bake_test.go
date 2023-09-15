package _5Kyu

import "testing"

func TestCake(t *testing.T) {
	got := Cakes(map[string]int{"apples": 3, "flour": 300, "sugar": 150, "milk": 100, "oil": 100},map[string]int{"sugar": 500, "flour": 2000, "milk": 2000})
	want := 0

	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}