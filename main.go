package main

import (
	"fmt"
	"strings"
)

var mapping = map[string]string{
	"n": "0",
	"l": "1",
	"k": "2",
	"g": "2",
	"m": "3",
	"t": "4",
	"v": "5",
	"w": "5",
	"f": "5",
	"p": "6",
	"b": "6",
	"s": "7",
	"š": "7",
	"z": "7",
	"ž": "7",
	"r": "8",
	"j": "9",
	"d": "9",
	"h": "9",
}

// String2Num converts strings to numbers
func String2Num(str string) string {
	output := strings.Builder{}
	for _, ch := range str {
		ch := strings.ToLower(string(ch))
		fmt.Fprint(&output, mapping[ch])
	}
	return output.String()
}

func MarkTask(challenge string, response string, assessment map[string]string) {
	correct := String2Num(challenge)
	// all correct
	if correct == response {
		for _, ch := range challenge {
			number := mapping[string(ch)]
			if number == "" {
				continue
			}
			assessment[number] += "+"
		}
		return
	}

	// var err error
	// cReader := strings.NewReader(String2Num(challenge))
	// rReader := strings.NewReader(response)
	// for err == nil {
	// 	chCorr, _, err := cReader.ReadRune()
	// 	if err == io.EOF {
	// 		err = nil
	// 		for chResp := rReader.ReadRune(); err == nil

	// 	}
	// 	if err != nil {

	// 	}

	// 	correct
	// }
	//return map[string]string{}
}

func main() {
	MarkTask("appi", "66", map[string]string{})
	// rand.Seed(time.Now().UTC().UnixNano())
	// for i := 0; i < 25; i++ {
	// 	fmt.Println(rand.Intn(10))
	// }
}
