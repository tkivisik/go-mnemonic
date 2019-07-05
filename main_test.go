package main

import (
	"fmt"
	"testing"
)

func TestString2Num(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  string
	}{
		{desc: "empty input", input: "", want: ""},
		{desc: "letter not in map", input: "a", want: ""},
		{desc: "number as input", input: "1", want: ""},
		{desc: "correct input lower case", input: "luuk", want: "12"},
		{desc: "correct input mixed case", input: "LuUk", want: "12"},
		{desc: "letter order 21", input: "kuul", want: "21"},
		{desc: "correct input 34", input: "amet", want: "34"},
		{desc: "correct input 56", input: "vaip", want: "56"},
		{desc: "correct input 78", input: "saar", want: "78"},
		{desc: "correct input 90", input: "hani", want: "90"},
		{desc: "correct input double same letter", input: "lukk", want: "122"},
		{desc: "sentence as input", input: "hello world!", want: "9115819"},
		{desc: "common symbols success", input: "!@#$s%^&*m()_+s", want: "737"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if output := String2Num(tC.input); output != tC.want {
				t.Errorf("String2Num(%q) = %q, want %q", tC.input, output, tC.want)
				// t.Fatalf("want f(%s) = %s, got: %s", tC.input, tC.want, output)
			}
		})
	}
}

func TestMarkTask(t *testing.T) {
	testCases := []struct {
		desc      string
		challenge string
		response  string
		want      map[string]string
	}{
		{desc: "all correct", challenge: "lukk", response: "122", want: map[string]string{"1": "+", "2": "++"}},
		{desc: "all wrong", challenge: "lukk", response: "000", want: map[string]string{"1": "-", "2": "--", "0": "---"}},
		{desc: "partially wrong", challenge: "lukk", response: "132", want: map[string]string{"1": "+", "2": "-+", "3": "-"}},
		{desc: "no response", challenge: "lukk", response: "", want: map[string]string{"1": "-", "2": "--"}},
		{desc: "too long", challenge: "luu", response: "123456", want: map[string]string{"1": "+", "2": "-", "3": "-", "4": "-", "5": "-", "6": "-"}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assessment := map[string]string{}
			if MarkTask(tC.challenge, tC.response, assessment); fmt.Sprint(assessment) != fmt.Sprint(tC.want) {
				// t.Errorf("MarkTask(%q) = %q, want %q", tC.challenge, assessment, tC.want)
				t.Errorf("MarkTask(%q, %q) = %v, want: %v", tC.challenge, tC.response, assessment, tC.want)
			}
		})
	}
}
