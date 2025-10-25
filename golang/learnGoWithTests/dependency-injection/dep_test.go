package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "Philip")

	want := "Hello, Philip"
	got := buffer.String()

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
