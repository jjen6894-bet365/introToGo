package bob

import "testing"

func TestHello(t *testing.T) {
    got := Hello("Jack")
    want := "Hello, Jack"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
