package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := hello("John")

	http.Handle("/", templ.Handler(component))

	lisAddr := ":3030"

	fmt.Printf("listening on %s\n", lisAddr)
	err := http.ListenAndServe(lisAddr, nil)
	if err != nil {
		panic(err)
	}
}