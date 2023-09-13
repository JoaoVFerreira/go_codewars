package _8Kyu

import "testing"

func TestCountSheeps(t * testing.T) {
	got := CountSheeps(Sheeps)
	want := 17

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}