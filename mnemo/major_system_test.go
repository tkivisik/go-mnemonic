package mnemo_test

import (
	"testing"

	"github.com/tkivisik/go-mnemonic/mnemo"
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
			if output := mnemo.String2Num(tC.input); output != tC.want {
				t.Errorf("String2Num(%q) = %q, want %q", tC.input, output, tC.want)
			}
		})
	}
}
