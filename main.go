package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	http.HandleFunc("/", todos)
	fmt.Println(http.ListenAndServe(":8000", nil))
}

func todos(w http.ResponseWriter, r *http.Request) {
	tmp.Execute(w, "Todo app")
}
