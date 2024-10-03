package ch2integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Add two integers in bounds", func(t *testing.T) {
		output := Add(2, 3)
		expected := 5
		assertIntegersAreEqual(t, output, expected)
	})	
}

func assertIntegersAreEqual(t testing.TB, output, expected int) {
	t.Helper()
	if output != expected {
		t.Errorf("Output: %d Expected: %d", output, expected)
	}
}

func ExampleAdd() {
	sum := Add(5,6)
	fmt.Println(sum)
	// Output: 11
}

