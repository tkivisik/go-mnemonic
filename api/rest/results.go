package rest

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/tkivisik/go-mnemonic/mnemo"
)

var ResultTmpl = template.Must(template.New("").Parse(`{{define "result"}}
    <html>
	  <head>
		<title>Mnemoharjutused</title>
		<meta charset="UTF-8">
	  </head>
	  <body>
	    Challenge: {{.Question}}<br>
	    Answer: {{.Answer}}<br>
		Assessment: {{.Assessment}}<br>
		Total Assessment: {{.AssessmentTotal}}<br>
		<button onclick="window.location.href = '/numberchallenge';" autofocus>Click Here</button>
	  </body>
	</html>{{end}}`))

func answer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	challenge := r.FormValue("challenge")
	answer := r.FormValue("answer")
	assessment := mnemo.AssessFromNumber(challenge, answer)
	assessmentTotal := map[string]string{}

	cookie, err := r.Cookie("mnemo-challenge")
	if err != nil {
		b64 := base64.StdEncoding.EncodeToString([]byte("1"))
		cookie = &http.Cookie{Name: "mnemo-challenge", Value: b64}
	}
	valueB, err := base64.StdEncoding.DecodeString(cookie.Value)
	value := string(valueB)
	if counter, err := strconv.Atoi(value); counter == 0 {
		if err != nil {
			fmt.Println("unexpected cookie value")
		}
		assessmentTotal = assessment
	}

	cookie = &http.Cookie{Name: "assessmentPast", Value: ""}
	cookie, err = r.Cookie("assessmentPast")
	if err != nil {
		cookie = &http.Cookie{Name: "assessmentPast", Value: ""}
	}
	valueB, err = base64.StdEncoding.DecodeString(cookie.Value)
	value = string(valueB)
	assessments := strings.Split(value, ";")
	for i := 0; i < len(assessments); i++ {
		parts := strings.Split(assessments[i], ":")
		if len(parts) > 1 {
			assessmentTotal[parts[0]] = parts[1]
		}
	}
	for k, v := range assessment {
		assessmentTotal[k] = fmt.Sprintf("%s%s", assessmentTotal[k], v)
	}
	builder := strings.Builder{}
	for k, v := range assessmentTotal {
		builder.WriteString(fmt.Sprintf("%s:%s;", k, v))
	}
	b64 := base64.StdEncoding.EncodeToString([]byte(builder.String()))
	cookie.Value = b64
	http.SetCookie(w, cookie)

	resultData := Challenge{
		Question:        challenge,
		Answer:          answer,
		Assessment:      assessment,
		AssessmentTotal: assessmentTotal,
	}

	rd := ResultData{w: w}
	rd.Show(ResultTmpl, &resultData)

	// err = ResultTmpl.ExecuteTemplate(w, "result", resultData)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
