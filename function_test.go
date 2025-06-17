package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	recieved := add(2, 3)
	actual := 5

	if recieved != actual {
		t.Errorf("Add(2,3) = %d ; actual = %d", recieved, actual)
	}
}

func TestAdd2(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Add positive numbers", 2, 3, 5},
		{"Add negative numbers", -1, -2, -3},
		{"Add zero", 0, 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := add(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Test %s failed. Expected %d, got %d", tc.name, tc.expected, result)
			}
		})
	}
}
