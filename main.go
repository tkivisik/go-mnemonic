package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tkivisik/go-mnemonic/mnemo"
)

func main() {
	flag.Parse()
	input := flag.Args()
	for i := 0; i < len(input); i++ {
		fmt.Printf("%10s ==> %s\n", input[i], mnemo.String2Num(input[i]))
	}

	// fmt.Println("TOTAL\n", sumAssessment)

	r := mux.NewRouter()
	r.HandleFunc("/", greet)
	r.HandleFunc("/numberchallenge", challenge).Methods(http.MethodGet)
	r.HandleFunc("/answer", answer).Methods(http.MethodPost)
	log.Println(http.ListenAndServe(":8080", r))

	fmt.Println(mnemo.TrainFromNumbers(1))
	fmt.Println(mnemo.TrainFromText(1))

}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

var ChallengeTmpl = template.Must(template.New("").Parse(`{{define "Challenge"}}
  <html>
    <head>
      <title>Mnemoharjutused</title>
      <meta charset="UTF-8">
    </head>
    <body>
      Mis on {{.}} tähtedena?
      <form action="/answer" method="post">
        <input type="text" name="answer" autofocus>
        <input type="hidden" name="challenge" value="{{.}}"><br>
        <input type="submit" value="Submit">
	  </form>
	</body>
  </html>
{{end}}`))

func challenge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	rand.Seed(time.Now().UnixNano())
	challenge := fmt.Sprintf("%02d", rand.Intn(100))

	err := ChallengeTmpl.ExecuteTemplate(w, "Challenge", challenge)
	if err != nil {
		fmt.Println(err)
	}

}

type Challenge struct {
	Question   string
	Answer     string
	Assessment map[string]string
}

type Resulter interface {
	Show(*template.Template, *Challenge)
}

type ResultData struct {
	w io.Writer
}

func (r ResultData) Show(t *template.Template, c *Challenge) {
	err := t.ExecuteTemplate(r.w, "result", c)
	if err != nil {
		fmt.Println(err)
	}
}

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
