package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
)

type todo struct {
	ID    string
	Title string
	Done  bool
}

var mapMake = make(map[string]*todo)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("templates/*.html"))

}

func main() {

	http.HandleFunc("/", todos)
	fmt.Println(http.ListenAndServe(":8000", nil))
}

func todos(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmp.Execute(w, mapMake)
		return
	}

	switch r.Method {
	case "POST":

		ramdomID := randomString()
		fmt.Println(ramdomID)

		mapMake[ramdomID] = &todo{ID: ramdomID, Title: r.FormValue("title"), Done: false}
		fmt.Println(mapMake)
		fmt.Println(mapMake[ramdomID])

		tmp.Execute(w, mapMake)

	case "DELETE":
		fmt.Println(r.FormValue("sajid"))
		tmp.Execute(w, "DELETE")
	default:
		tmp.Execute(w, nil)

	}

}

func randomString() string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 10)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
