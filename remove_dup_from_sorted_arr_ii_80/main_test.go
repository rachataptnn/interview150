package main

import (
	"reflect"
	"testing"
)

func TestRemoveDupII(t *testing.T) {
	type testcaseStruct struct {
		casename     string
		nums         []int
		expected     int
		modifiedNums []int
	}

	testcases := []testcaseStruct{
		{
			casename:     "case 1",
			nums:         []int{1, 1, 1, 2, 2, 3},
			expected:     5,
			modifiedNums: []int{1, 1, 2, 2, 3},
		},
		{
			casename:     "case 2",
			nums:         []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expected:     7,
			modifiedNums: []int{0, 0, 1, 1, 2, 3, 3},
		},
	}

	for _, tc := range testcases {
		actual, modNums := removeDuplicates(tc.nums)

		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, actual)
		}

		if !reflect.DeepEqual(modNums, tc.modifiedNums) {
			t.Errorf("Expected %v, but got %v", tc.modifiedNums, modNums)
		}
	}
}
