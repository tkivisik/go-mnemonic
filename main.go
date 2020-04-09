package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/tkivisik/go-mnemonic/api/rest"
	"github.com/tkivisik/go-mnemonic/mnemo"
)

func main() {
	port := ":8080"
	flag.Parse()
	input := flag.Args()
	for i := 0; i < len(input); i++ {
		fmt.Printf("%10s ==> %s\n", input[i], mnemo.String2Num(input[i]))
	}

	// fmt.Println(mnemo.TrainFromNumbers(1))
	// fmt.Println(mnemo.TrainFromText(1))

	r := rest.NewRouter()
	fmt.Printf("Server starting now at port %s\n", port)
	log.Println(http.ListenAndServe(port, r))
}
