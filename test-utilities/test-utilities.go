package testutilities

import (
	"reflect"
	"testing"
)

type primitive interface {
	bool | ~float64 | string | int
}

func AssertEqualError(t testing.TB, output error, expectedMessage string) {
	t.Helper()
	if output == nil {
		t.Error("Expected error was nil.")
	} else if expectedMessage != output.Error() {
		t.Errorf("Output error message %v did not equal expected message %v", output, expectedMessage)
	}
}

func AssertEqual[T primitive](t testing.TB, output, expected T) {
	t.Helper()
	if output != expected {
		t.Errorf("Output: %v, Expected: %v", output, expected)
	}
}

func AssertSlicesAreEqual[T any](t testing.TB, output, expected []T) {
	t.Helper()
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Output: %v, Expected: %v", output, expected)
	}
}
