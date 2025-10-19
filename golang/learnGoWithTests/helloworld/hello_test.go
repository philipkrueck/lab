package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Phil", "")
		want := "Hello, Phil"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying 'Hello World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Sergi", "Spanish")
		want := "Hola, Sergi"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Hugo", "French")
		want := "Bonjour, Hugo"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in German", func(t *testing.T) {
		got := Hello("Philip", "German")
		want := "Hallo, Philip"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
