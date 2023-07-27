package main

import (
	"reflect"
	"testing"
)

func TestRemoveDupII(t *testing.T) {
	type testcaseStruct struct {
		casename string
		s        string
		t        string
		expected bool
	}

	testcases := []testcaseStruct{
		{
			casename: "case 1",
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			casename: "case 2",
			s:        "rat",
			t:        "car",
			expected: false,
		},
	}

	for _, tc := range testcases {
		actual := isAnagram(tc.s, tc.t)

		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("testcase: %s - Expected %v, but got %v", tc.casename, tc.expected, actual)
		}
	}
}
