package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {

	b := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(b, spySleeper)

	got := b.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}
