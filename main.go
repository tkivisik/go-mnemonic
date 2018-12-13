package main

import (
	"fmt"
)

var mapping = map[string]int{
	"n": 0,
	"l": 1,
	"k": 2,
	"g": 2,
	"m": 3,
	"t": 4,
	"v": 5,
	"w": 5,
	"f": 5,
	"p": 6,
	"b": 6,
	"s": 7,
	"š": 7,
	"z": 7,
	"ž": 7,
	"r": 8,
	"j": 9,
	"d": 9,
	"h": 9,
}

// String2Num converts strings to numbers
func String2Num(str string) string {
	//for i := 0; i < len(str); i++ {
	//	str[i] in mapping

	//}
	return str
}

func main() {

	var result string
	var allikas = "kala"
	result = String2Num(allikas)

	fmt.Println(result)
	fmt.Println(mapping)
}
