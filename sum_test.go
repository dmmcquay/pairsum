package pairsum

import "testing"

func TestDoesItSum(t *testing.T) {
	var tests = []struct {
		a        []int
		s        int
		expected bool
	}{
		{a: []int{1, 3, 5, 9}, s: 8, expected: true},
		{a: []int{1, 4, 5, 9}, s: 8, expected: false},
		{a: []int{1, 3, 4, 4}, s: 8, expected: true},
	}
	for _, rt := range tests {
		actual := DoesItSum(rt.a, rt.s)
		if actual != rt.expected {
			t.Errorf(
				"failed doesItSum \n\texpected: %t\n\t  actual: %t",
				rt.expected,
				actual,
			)
		}
	}
}
