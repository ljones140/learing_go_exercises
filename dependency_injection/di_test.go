package main

import "bytes"
import "testing"

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	// &buffer is passing a reference to the buffer
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
