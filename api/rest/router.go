package rest

import (
	"fmt"
	"io"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

// NewRouter returns a *mux.Router with app's routes
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", greet)
	r.HandleFunc("/numberchallenge", challenge).Methods(http.MethodGet)
	r.HandleFunc("/answer", answer).Methods(http.MethodPost)

	return r
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
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
