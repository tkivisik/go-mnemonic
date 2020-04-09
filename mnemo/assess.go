package mnemo

import (
	"strconv"
	"strings"
)

// AssessFromNumber will be public API, should be tested.
func AssessFromNumber(numChallenge, textAnswer string) (assessment map[string]string) {
	return markTask(numChallenge, textAnswer)
}

// AssessFromText will be public API, should be tested.
func AssessFromText(textChallenge, numAnswer string) (assessment map[string]string) {
	return markTask(textChallenge, numAnswer)
}

// markTask assesses where was a given response correct and where was it
// mistaken.
func markTask(challenge string, response string) (assessment map[string]string) {
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
