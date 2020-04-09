package rest

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
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
	cookie, err := r.Cookie("mnemo-challenge")
	if err != nil {
		cookie = &http.Cookie{Name: "mnemo-challenge", Value: "1"}
	}
	counter, err := strconv.Atoi(cookie.Value)
	if err != nil {
		cookie = &http.Cookie{Name: "mnemo-challenge", Value: "1"}
	}
	cookie.Value = fmt.Sprintf("%d", counter+1)
	http.SetCookie(w, cookie)

	rand.Seed(time.Now().UnixNano())
	challenge := fmt.Sprintf("%02d", rand.Intn(100))

	err = ChallengeTmpl.ExecuteTemplate(w, "Challenge", challenge)
	if err != nil {
		fmt.Println(err)
	}

}

type Challenge struct {
	Question   string
	Answer     string
	Assessment map[string]string
}
