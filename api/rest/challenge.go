package rest

import (
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

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
