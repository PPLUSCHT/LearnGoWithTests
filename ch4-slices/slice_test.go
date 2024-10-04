package ch4slices

import (
	"testing"

	testutil "github.com/ppluscht/learngowithtests/test-utilities"
)

func TestSum(t *testing.T) {
	t.Run("Should sum all the numbers in a slice", func(t *testing.T) {
		output := 10
		expected := Sum([]int{1, 2, 3, 4})
		testutil.AssertIntegersAreEqual(t, output, expected)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Should sum each slice", func(t *testing.T) {
		output := []int{10, 13, 6, 20, 0}
		expected := SumAll([]int{1, 3, 6}, []int{4, 3, 6}, []int{2, 3, 1}, []int{4, 7, 5, 4}, []int{})
		testutil.AssertSlicesAreEqual(t, output, expected)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("Should sum all but the first element of each slice", func(t *testing.T) {
		output := []int{9, 9, 4, 16, 0}
		expected := SumAllTails([]int{1, 3, 6}, []int{4, 3, 6}, []int{2, 3, 1}, []int{4, 7, 5, 4}, []int{})
		testutil.AssertSlicesAreEqual(t, output, expected)
	})
}
