package rest

import (
	"fmt"
	"net/http"
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

	resultData := Challenge{
		Question:   challenge,
		Answer:     answer,
		Assessment: assessment,
	}

	rd := ResultData{w: w}
	rd.Show(ResultTmpl, &resultData)

	// err = ResultTmpl.ExecuteTemplate(w, "result", resultData)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
