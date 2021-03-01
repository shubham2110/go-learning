package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello("Shubham")
    want := "Hello, Shubham"
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
	got = Hello("Rahul")
    want = "Hello, Rahul"
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
	
}