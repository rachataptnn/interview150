package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type mergeStruct struct {
		casename string
		nums     []int
		val      int
		expected int
	}

	testcases := []mergeStruct{
		{
			casename: "case 1",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: 2,
		},
		{
			casename: "case 2",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: 5,
		},
	}

	for _, tc := range testcases {
		actual := removeElement(tc.nums, tc.val)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, actual)
		}
	}
}
