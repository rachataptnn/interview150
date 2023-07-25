package main

import (
	"reflect"
	"testing"
)

func TestRemoveDupII(t *testing.T) {
	type testcaseStruct struct {
		casename string
		input    []int
		pos      int
		expected bool
	}

	testcases := []testcaseStruct{
		{
			casename: "case 1",
			input:    []int{3, 2, 0, -4},
			pos:      1,
			expected: true,
		},
		{
			casename: "case 2",
			input:    []int{1, 2},
			pos:      0,
			expected: true,
		},
		{
			casename: "case 3",
			input:    []int{1},
			pos:      -1,
			expected: false,
		},
	}

	for _, tc := range testcases {
		actual := hasCycle(createCyclicLinkedList(tc.input, tc.pos))

		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("testcase: %s - Expected %v, but got %v", tc.casename, tc.expected, actual)
		}
	}
}
