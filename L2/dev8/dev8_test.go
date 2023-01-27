package main

import "testing"

func TestShall(t *testing.T) {
	type testCase struct {
		description string
		input       string
		expected    string
	}
	testCases := []testCase{
		{
			description: "Checking path ",
			input:       pwd(),
			expected:    "/mnt/d/WIldberries_internship/L2/dev8",
		},
	}
	for _, tc := range testCases {
		if pwd() != tc.expected {
			t.Errorf("Expected %s but got %s ", tc.expected, pwd())
		}
	}
}
