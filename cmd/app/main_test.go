package main

import (
	"os"
	"testing"
)

func TestGetenv(t *testing.T) {
	testcases := []struct {
		name     string
		env      string
		def      string
		expected string
	}{
		{
			name:     "empty env",
			env:      "",
			def:      "default",
			expected: "default",
		},
		{
			name:     "empty def",
			env:      "env",
			def:      "",
			expected: "env",
		},
		{
			name:     "env",
			env:      "env",
			def:      "default",
			expected: "env",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv("TEST_ENV", tc.env)
			actual := getEnv("TEST_ENV", tc.def)
			if actual != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, actual)
			}
		})
	}
}
