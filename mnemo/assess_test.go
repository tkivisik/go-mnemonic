package mnemo_test

import (
	"fmt"
	"testing"

	"github.com/tkivisik/go-mnemonic/mnemo"
)

func TestAssessFromText(t *testing.T) {
	testCases := []struct {
		desc      string
		challenge string
		response  string
		want      map[string]string
	}{
		{desc: "all correct", challenge: "lukk", response: "122", want: map[string]string{"l": "+", "k": "++"}},
		{desc: "all wrong", challenge: "lukk", response: "000", want: map[string]string{"l": "-", "k": "--", "0": "---"}},
		{desc: "partially wrong", challenge: "lukk", response: "132", want: map[string]string{"l": "+", "k": "-+", "3": "-"}},
		{desc: "no response", challenge: "lukk", response: "", want: map[string]string{"l": "-", "k": "--"}},
		{desc: "too long", challenge: "luu", response: "123456", want: map[string]string{"l": "+", "2": "-", "3": "-", "4": "-", "5": "-", "6": "-"}},
		{desc: "should not ignore non-ints", challenge: "lukk", response: "l122", want: map[string]string{"1": "-", "2": "-", "k": "-+", "l": "--"}},
		{desc: "should not ignore vocals", challenge: "l", response: "a1", want: map[string]string{"1": "-", "a": "-", "l": "-"}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if assessment := mnemo.AssessFromText(tC.challenge, tC.response); fmt.Sprint(assessment) != fmt.Sprint(tC.want) {
				t.Errorf("MarkTask(%q, %q) = %v, want: %v", tC.challenge, tC.response, assessment, tC.want)
			}
		})
	}
}

func TestAssessFromNumber(t *testing.T) {
	testCases := []struct {
		desc      string
		challenge string
		response  string
		want      map[string]string
	}{
		{desc: "all correct", challenge: "122", response: "lkk", want: map[string]string{"1": "+", "2": "++"}},
		{desc: "all wrong", challenge: "122", response: "nnn", want: map[string]string{"1": "-", "2": "--", "n": "---"}},
		{desc: "partially wrong", challenge: "122", response: "lmk", want: map[string]string{"1": "+", "2": "-+", "m": "-"}},
		{desc: "no response", challenge: "122", response: "", want: map[string]string{"1": "-", "2": "--"}},
		{desc: "too long", challenge: "1", response: "lkmtvp", want: map[string]string{"1": "+", "k": "-", "m": "-", "t": "-", "v": "-", "p": "-"}},
		{desc: "should not ignore non-ints", challenge: "122", response: "llkk", want: map[string]string{"1": "+", "2": "-+", "k": "-", "l": "-"}},
		{desc: "all wrong - vocals", challenge: "122", response: "aae", want: map[string]string{"1": "-", "2": "--", "a": "--", "e": "-"}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if assessment := mnemo.AssessFromNumber(tC.challenge, tC.response); fmt.Sprint(assessment) != fmt.Sprint(tC.want) {
				t.Errorf("MarkTask(%q, %q) = %v, want: %v", tC.challenge, tC.response, assessment, tC.want)
			}
		})
	}
}
