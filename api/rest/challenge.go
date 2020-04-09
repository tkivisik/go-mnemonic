package rest

import (
	"encoding/base64"
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

	cookie := IncrementCookie("mnemo-challenge", r)
	http.SetCookie(w, cookie)

	rand.Seed(time.Now().UnixNano())
	challenge := fmt.Sprintf("%02d", rand.Intn(100))

	err := ChallengeTmpl.ExecuteTemplate(w, "Challenge", challenge)
	if err != nil {
		fmt.Println(err)
	}

}

func IncrementCookie(name string, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("mnemo-challenge")
	if err != nil {
		b64 := base64.StdEncoding.EncodeToString([]byte("0"))
		cookie = &http.Cookie{Name: "mnemo-challenge", Value: b64}
	}
	valueB, err := base64.StdEncoding.DecodeString(cookie.Value)
	value := string(valueB)

	counter, err := strconv.Atoi(value)
	if err != nil {
		b64 := base64.StdEncoding.EncodeToString([]byte("0"))
		cookie = &http.Cookie{Name: "mnemo-challenge", Value: b64}
	}
	b64 := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", counter+1)))
	cookie.Value = b64
	return cookie
}

type Challenge struct {
	Question        string
	Answer          string
	Assessment      map[string]string
	AssessmentTotal map[string]string
}
