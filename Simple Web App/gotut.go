package main

import (
	"net/http"
	"fmt"
	"html/template"
)

type NewsAggPage struct {
	Title string
	News string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title:"Amazing News Aggregator", News:"some news"}
	t, err := template.ParseFiles("basictemplating.html")
	if err != nil {

	}
	t.Execute(w,p)
	//fmt.Fprintf(w, "AAA")
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
<h1>Whoa, Go is neat</h1>
<p>Whoa, Go is fast !</p>
`)

	fmt.Fprintf(w, "<h1>Whoa, Go is neat</h1>")
	fmt.Fprintf(w, "<p>Whoa, Go is fast!</p>")
	fmt.Fprintf(w, "<h1>Whoa, Go is neat</h1>")
	fmt.Fprintf(w, "<p>You even %s add %s</p>", "can", "<strong>variables</strong>")
}

func about_handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Created by Aditya Alif Nugraha")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}
