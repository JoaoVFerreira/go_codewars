package _6Kyu

import "testing"

func TestCountingDuplicates(t *testing.T) {
	got := CountDuplicates("aabBcde")
	want := 2

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}