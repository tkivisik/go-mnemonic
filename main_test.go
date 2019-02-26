package main

import (
	"fmt"
	"testing"
)

func TestString2Num(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "empty input",
			input:    "",
			expected: "",
		},
		{
			desc:     "letter not in map",
			input:    "a",
			expected: "",
		},
		{
			desc:     "number as input",
			input:    "1",
			expected: "",
		},
		{
			desc:     "correct input lower case",
			input:    "luuk",
			expected: "12",
		},
		{
			desc:     "correct input upper-lower case",
			input:    "LuUk",
			expected: "12",
		},
		{
			desc:     "letter order",
			input:    "kuul",
			expected: "21",
		},
		{
			desc:     "correct input",
			input:    "amet",
			expected: "34",
		},
		{
			desc:     "correct input",
			input:    "vaip",
			expected: "56",
		},
		{
			desc:     "correct input",
			input:    "saar",
			expected: "78",
		},
		{
			desc:     "correct input",
			input:    "hani",
			expected: "90",
		},
		{
			desc:     "correct input double same letter",
			input:    "lukk",
			expected: "122",
		},
		{
			desc:     "sentence as input",
			input:    "hello world!",
			expected: "9115819",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if output := String2Num(tC.input); output != tC.expected {
				t.Fatalf("expected f(%s) = %s, got: %s", tC.input, tC.expected, output)
			}
		})
	}
}

func TestMarkTask(t *testing.T) {
	testCases := []struct {
		desc      string
		challenge string
		response  string
		expected  map[string]string
	}{
		{
			desc:      "all correct",
			challenge: "lukk",
			response:  "122",
			expected:  map[string]string{"1": "+", "2": "++"},
		},
		{
			desc:      "all wrong",
			challenge: "lukk",
			response:  "000",
			expected:  map[string]string{"1": "-", "2": "--", "0": "---"},
		},
		{
			desc:      "partially wrong",
			challenge: "lukk",
			response:  "132",
			expected:  map[string]string{"1": "+", "2": "-+", "3": "-"},
		},
		{
			desc:      "no response",
			challenge: "lukk",
			response:  "",
			expected:  map[string]string{"1": "-", "2": "--"},
		},
		{
			desc:      "too long",
			challenge: "luu",
			response:  "123456",
			expected:  map[string]string{"1": "+", "2": "-", "3": "-", "4": "-", "5": "-", "6": "-"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assessment := map[string]string{}
			if MarkTask(tC.challenge, tC.response, assessment); fmt.Sprint(assessment) != fmt.Sprint(tC.expected) {
				t.Fatalf("expected f(%s, %s) = %s, got: %s", tC.challenge, tC.response, tC.expected, assessment)
			}
		})
	}
}
