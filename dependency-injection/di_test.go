package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Levi")

	got := buffer.String()
	want := "Hello, Levi"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
