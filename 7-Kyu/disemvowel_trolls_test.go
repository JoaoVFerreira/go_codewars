package _7Kyu

import "testing"

func TestDisemvowel(t *testing.T) {
	got := Disemvowel("This website is for losers LOL!")
	want := "Ths wbst s fr lsrs LL!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}