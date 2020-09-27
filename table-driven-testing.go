package main

import (
	"strings"
	"testing"
)

// table-driven testing
func TestSplit(t *testing.T) {
	var tests = []struct {
		text string
		sep  string
		want int
	}{
		{"a:b:c", ":b:", 2},
		{"a:b:c:d", ":", 4},
		{"a;b;c;d", ";", 4},
	}
	for _, test := range tests {
		text, sep := test.text, test.sep
		words := strings.Split(text, sep)
		if got, want := len(words), test.want; got != want {
			t.Errorf("Split(%q, %q) returned %d words, want %d",
				text, sep, got, want)
		}
	}
}
