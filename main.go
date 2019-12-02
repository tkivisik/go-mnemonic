package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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

// MarkTask assesses where was a given response correct and where was it
// mistaken.
// TODO: upgrade to collect data required for a ROC curve.
func MarkTask(challenge string, response string) (assessment map[string]string) {
	isFromIntChallenge := true
	_, err := strconv.Atoi(challenge)
	if err != nil {
		isFromIntChallenge = false
	}

	assessment = map[string]string{}
	chalReader := strings.NewReader(challenge)
	respReader := strings.NewReader(response)
	// var chalCh, respCh string
	var chalFinished, respFinished bool
	var want, got string
	for {
		// Read new characters
		if !chalFinished {
			want, chalFinished = nextCh(chalReader)

			if want == "" || !isFromIntChallenge && String2Num(want) == "" {
				continue
			}
		}

		if !respFinished {
			got, respFinished = nextCh(respReader)
		}

		if chalFinished && respFinished {
			// BREAKS THE FOR{}
			break
		}

		// Assess
		if got == "" {
			assessment[want] += "-"
			continue
		}
		if want == "" {
			assessment[got] += "-"
			continue
		}

		if !isFromIntChallenge {
			if String2Num(want) != got {
				assessment[want] += "-"
				assessment[got] += "-"
				continue
			}
		}
		if isFromIntChallenge {
			if want != String2Num(got) {
				assessment[want] += "-"
				assessment[got] += "-"
				continue
			}
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

	// fmt.Println(TrainFromNumbers(5))
	fmt.Println(TrainFromText(5))

	// reader := bufio.NewReader(os.Stdin)

	// //
	// // assessment := map[string]string{}
	// numChallenge := "0123"
	// fmt.Printf("Mis on %q tähtedena?\n", numChallenge)
	// textAnswer, _ := reader.ReadString('\n')
	// textAnswer = strings.TrimSpace(textAnswer)
	// assessment := AssessFromNumber(numChallenge, textAnswer)
	// fmt.Println(assessment)

	// sumAssessment := map[string]string{}
	// fmt.Println("EMPTY", sumAssessment)
	// for k, v := range assessment {
	// 	sumAssessment[k] = v
	// }
	// fmt.Println("FIRST", sumAssessment)

	// //
	// textChallenge := "kalamaja"
	// fmt.Printf("Mis on %q numbrina?\n", textChallenge)
	// numAnswer, _ := reader.ReadString('\n')
	// numAnswer = strings.TrimSpace(numAnswer)
	// assessment = AssessFromText(textChallenge, numAnswer)
	// fmt.Println(assessment)

	// for k, v := range assessment {
	// 	sumAssessment[k] += v
	// }

	// fmt.Println("TOTAL\n", sumAssessment)

}

// AssessFromNumber will be public API, should be tested.
func AssessFromNumber(numChallenge, textAnswer string) (assessment map[string]string) {
	return MarkTask(numChallenge, textAnswer)
}

// AssessFromText will be public API, should be tested.
func AssessFromText(textChallenge, numAnswer string) (assessment map[string]string) {
	return MarkTask(textChallenge, numAnswer)
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
