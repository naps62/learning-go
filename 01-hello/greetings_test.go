package greetings

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("naps")
		want := "Hello, naps. Welcome!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is supply", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World. Welcome!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
