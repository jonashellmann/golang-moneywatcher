package main

import (
  // "fmt"
  "net/http"
  "text/template"
)

type Person struct {
	Name string
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5656", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("templates/index.tmpl")
	err := t.Execute(w, "Hello World!")

	if err != nil {
		panic(err)
	}
}
