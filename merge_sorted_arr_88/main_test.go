package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type mergeStruct struct {
		casename string
		nums1    []int
		nums2    []int
		m        int
		n        int
		expected []int
	}

	testcases := []mergeStruct{
		{
			casename: "case 1",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			nums2:    []int{2, 5, 6},
			m:        3,
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			casename: "case 2",
			nums1:    []int{1},
			nums2:    []int{},
			m:        1,
			n:        0,
			expected: []int{1},
		},
		{
			casename: "case 3",
			nums1:    []int{0},
			nums2:    []int{1},
			m:        0,
			n:        1,
			expected: []int{1},
		},
	}

	for _, tc := range testcases {
		actual := merge(tc.nums1, tc.m, tc.nums2, tc.n)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, actual)
		}
	}
}
