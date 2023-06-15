package main

import (
	"reflect"
	"testing"
)

func TestRemoveDup(t *testing.T) {
	type testcaseStruct struct {
		casename string
		nums     []int
		expected int
	}

	testcases := []testcaseStruct{
		{
			casename: "case 1",
			nums:     []int{1, 1, 2},
			expected: 3,
		},
		{
			casename: "case 2",
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
		},
	}

	for _, tc := range testcases {
		actual := removeDuplicates(tc.nums)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, actual)
		}
	}
}
