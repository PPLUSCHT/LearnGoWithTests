package ch3iteration

import (
	"testing"

	testutil "github.com/ppluscht/learngowithtests/test-utilities"
)

func TestRepeat(t *testing.T) {
	output := RepeatN("1", 5)
	expected := "11111"
	testutil.AssertStringsAreCorrect(t, output, expected)
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatN("@", 5)
	}
}
