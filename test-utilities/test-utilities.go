package testutilities

import (
	"reflect"
	"testing"
)

type primitive interface {
	bool | float64 | string | int
}

func AssertEqual[T primitive](t testing.TB, output, expected T) {
	if output != expected {
		t.Errorf("Output: %v, Expected: %v", output, expected)
	}
}

func AssertSlicesAreEqual[T any](t testing.TB, output, expected []T) {
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Output: %v, Expected: %v", output, expected)
	}
}
