package helloworld

import "testing"

func TestHello(t *testing.T) {
	if got := Hello(); got != "Hello, world!" {
		t.Fatalf("Hello() = %q, want %q", got, "Hello, world!")
	}
}
