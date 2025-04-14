package main

import (
	"fmt"
	"testing"
)

func Test_cleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  helloworld  ",
			expected: []string{"helloworld"},
		},
		{
			input:    "  hello  world  sdasda",
			expected: []string{"hello", "world", "sdasda"},
		},
		{
			input:    "             ",
			expected: nil,
		},
	}
	for i, c := range cases {

		sliceStr := cleanInput(c.input)

		if sliceStr == nil {
			t.Error("Failed to get slice")
			return
		}

		if len(sliceStr) != len(c.expected) {
			fmt.Printf("sliceLen:%d : expLen:%d", len(sliceStr), len(c.expected))
			t.Errorf("%d -- Not the same len", i+1)
			return
		}

		for i := range sliceStr {
			if sliceStr[i] != c.expected[i] {
				fmt.Printf("GET:%s, EXP:%s", sliceStr[i], c.expected[i])
				t.Error("Input not match")
				return
			}

		}
	}
}
