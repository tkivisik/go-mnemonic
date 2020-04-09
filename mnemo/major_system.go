package mnemo

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Mapping is the N:1 mapping of letters to numbers.
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

// String2Num converts strings to numbers in string
func String2Num(str string) string {
	output := strings.Builder{}
	for _, ch := range str {
		ch := strings.ToLower(string(ch))
		fmt.Fprint(&output, Mapping[ch])
	}
	return output.String()
}

func nextCh(r *strings.Reader) (char string, isFinished bool) {
	nextChar, _, err := r.ReadRune()
	if err != nil {
		if err == io.EOF {
			isFinished = true
			return "", isFinished
		}
		log.Println(err)
		return "", false
	}
	return string(nextChar), isFinished
}

func TrainFromNumbers(n int) (totalAssessment map[string]string) {
	reader := bufio.NewReader(os.Stdin)
	totalAssessment = map[string]string{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		// TrainFromNumber()
		challenge := fmt.Sprintf("%02d", rand.Intn(100))
		fmt.Printf("Mis on %q tähtedena? ", challenge)
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		response = strings.TrimSpace(response)
		assessment := AssessFromNumber(challenge, response)
		for k, v := range assessment {
			totalAssessment[k] += v
		}
	}
	return totalAssessment
}

func TrainFromText(n int) (totalAssessment map[string]string) {
	reader := bufio.NewReader(os.Stdin)
	totalAssessment = map[string]string{}
	rand.Seed(time.Now().UnixNano())
	alphabet := "nlkgmtvwfpbsrjdh"
	for i := 0; i < n; i++ {
		// TrainFromText()
		random := rand.Intn(16)

		challenge := fmt.Sprintf("%s", alphabet[random:random+1])
		fmt.Printf("Mis on %s numbrina? ", challenge)
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		response = strings.TrimSpace(response)
		assessment := AssessFromText(challenge, response)
		for k, v := range assessment {
			totalAssessment[k] += v
		}
	}
	return totalAssessment
}
