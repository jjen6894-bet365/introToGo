package bob

import "testing"

func assertCorrectMessage(t *testing.T, got, want string) {
 t.Helper()
 if got != want {
	t.Errorf("got: %q wanted: %q", got, want)
 } else {
	 println("PASS:: Expected and got were identical!")
 }
}


func TestHello(t *testing.T) {
    t.Run("saying hello to people", func(t *testing.T) {
   	 got := Hello("Jack")
	 want := "Hello, Jack"
    	 assertCorrectMessage(t, got, want)
    })
    
    t.Run("say 'Hello, World' when an empty string is the argument supplied", func(t *testing.T) {
	got := Hello("")
	expect := "Hello, World"
	assertCorrectMessage(t, got, expect)
    })
}
