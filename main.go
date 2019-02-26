package main

import (
	"fmt"
	"io"
	"log"
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
	respReader := strings.NewReader(response)
	for _, ch := range challenge {
		number := String2Num(string(ch))
		if number == "" {
			continue
		}

		respCh, _, err := respReader.ReadRune()
		if err != nil {
			if err == io.EOF {
				assessment[number] += "-"
				continue
			}
			log.Fatalln(err)
		}
		if number == string(respCh) {
			assessment[number] += "+"
			continue
		}
		assessment[number] += "-"
	}
	return
}

func main() {
	MarkTask("appi", "66", map[string]string{})

	// rand.Seed(time.Now().UTC().UnixNano())
	// for i := 0; i < 25; i++ {
	// 	fmt.Println(rand.Intn(10))
	// }
}
