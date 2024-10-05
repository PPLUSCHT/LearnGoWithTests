package ch2integers

import (
	"fmt"
	"testing"

	testutil "github.com/ppluscht/learngowithtests/test-utilities"
)

func TestAdder(t *testing.T) {
	t.Run("Add two integers in bounds", func(t *testing.T) {
		output := Add(2, 3)
		expected := 5
		testutil.AssertEqual(t, output, expected)
	})
}

func ExampleAdd() {
	sum := Add(5, 6)
	fmt.Println(sum)
	// Output: 11
}
