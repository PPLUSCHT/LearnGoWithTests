package ch5structs

import (
	"math"
	"testing"

	testutil "github.com/ppluscht/learngowithtests/test-utilities"
)

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape  Shape
		expect float64
	}{
		{Rectangle{11, 5}, 32.0},
		{Circle{10}, 20 * math.Pi},
	}

	for _, test := range perimeterTests {
		testutil.AssertEqual(t, test.shape.Perimeter(), test.expect)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape  Shape
		expect float64
	}{
		{Rectangle{11, 5}, 55.0},
		{Circle{10}, 100 * math.Pi},
	}

	for _, test := range areaTests {
		testutil.AssertEqual(t, test.shape.Area(), test.expect)
	}
}
