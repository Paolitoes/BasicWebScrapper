package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	test := []struct {
		name 		string
		inputURL	string
		expected 	string
	}{
		{
			name: 		"remove scheme",
			inputURL: 		"https://blog.boot.dev/path",
			expected: 	"blog.boot.dev/path",
		},
		{
			name: 		"remove sercure",
			inputURL: 		"https://blog.boot.dev/path",
			expected: 	"http://blog.boot.dev/path",
		},
		{
			name: 		"remove terminal slash",
			inputURL: 		"https://blog.boot.dev/path/",
			expected: 	"https://blog.boot.dev/path",
		},
		{
			name: 		"remove scheme and slash",
			inputURL: 		"https://blog.boot.dev/path/",
			expected: 	"blog.boot.dev/path",
		},
		{
			name: 		"remove secure and slash",
			inputURL: 		"http://blog.boot.dev/path/",
			expected: 	"http://blog.boot.dev/path",
		},
	}

	for i, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name,err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}


}

	


