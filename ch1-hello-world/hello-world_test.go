package ch1helloworld

import (
	"testing"

	testutil "github.com/ppluscht/learngowithtests/test-utilities"
)

func TestHello(t *testing.T) {
	t.Run("say hello to a non-empty name with English as the specified language", func(t *testing.T) {
		output := Hello("Chris", English)
		expected := "Hello, Chris"
		testutil.AssertEqual(t, output, expected)
	})
	t.Run("say 'Hello, World' when an empty input is provided with English as the specified language", func(t *testing.T) {
		output := Hello("", English)
		expected := "Hello, World"
		testutil.AssertEqual(t, output, expected)
	})

	t.Run("should say 'Hola, Jean' when Spanish is the specified language", func(t *testing.T) {
		output := Hello("Jean", Spanish)
		expected := "Hola, Jean"
		testutil.AssertEqual(t, output, expected)
	})

	t.Run("should say 'Hola, Mundo' when Spanish is the specified language and an empty name is provided", func(t *testing.T) {
		output := Hello("", Spanish)
		expected := "Hola, Mundo"
		testutil.AssertEqual(t, output, expected)
	})
}
