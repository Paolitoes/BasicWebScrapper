package main

import (
	"testing"
	"reflect"
)


func TestgetURLsFromHTML() {
	test := []struct {
		name		string
		inputURL	string
		inputBody	string
		expected	[]string
	}{
		{
			name: 		"absolute and relative URLs",
			inputURL: 	"https://blog.boot.dev",
			inputBody: '
				<html>
					<body>
						<a href="/path/one">
							<span>Boot.dev</span>
						</a>
						<a href="https://other.com/path/one">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
				',
			expected: 	[]string{"https://blog.boot.dev/path/one","https://other.com/path/one"},
		},{
			name: 		"absolute and relative URLs",
			inputURL: 	"https://blog.boot.dev",
			inputBody: '
				<html>
					<body>
						<a href="/path/one">
							<span>Boot.dev</span>
						</a>
						<a href="https://other.com/path/one">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
				',
			expected: 	[]string{"https://blog.boot.dev/path/one","https://other.com/path/one"},
		},
	}

	for i, tc := range test {
		t.Run(tc.name, func (t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody,tc.inputURL)
			if err != nil {
				t.Error("Test %v - '%s' FAIL: unexpected error %v", i, tc.name, err)
			}
			if DeepEqual(actual,tc.expected){
				t.Error("Test %v - %s FAIL: unexpected URL(s): %v, actual %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
