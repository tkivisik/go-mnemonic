package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"strings"
)

var Mapping = map[string]string{
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
		fmt.Fprint(&output, Mapping[ch])
	}
	return output.String()
}

func MarkTask(challenge string, response string) (assessment map[string]string) {
	chalReader := strings.NewReader(challenge)
	respReader := strings.NewReader(response)
	var chalCh, respCh rune
	var chalFinished, respFinished bool
	var err error
	var want string
	for {
		if chalFinished == false {
			chalCh, _, err = chalReader.ReadRune()
			if err != nil {
				if err == io.EOF {
					chalFinished = true
				} else {
					log.Println(err)
				}
			}

			want = String2Num(string(chalCh))
			if want == "" {
				continue
			}
		}

		if respFinished == false {
			respCh, _, err = respReader.ReadRune()
			if err != nil {
				if err == io.EOF {
					respFinished = true
				} else {
					log.Println(err)
				}
			}
		}

		if chalFinished {
			if respFinished {
				// BREAKS THE FOR{}
				break
			}
			got := string(respCh)
			assessment[got] += "-"
			continue
		}

		if respFinished {
			assessment[want] += "-"
			continue
		}

		if got := string(respCh); want != got {
			assessment[want] += "-"
			assessment[got] += "-"
			continue
		}
		assessment[want] += "+"
	}
	return assessment
}

func main() {
	flag.Parse()
	input := flag.Args()
	for i := 0; i < len(input); i++ {
		fmt.Printf("%10s ==> %s\n", input[i], String2Num(input[i]))
	}
}
