package main

import (
	"io/ioutil"
	"golang.org/x/tools/godoc"
	"net/http"
	"fmt"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	f := p.Title + ".txt"
	return ioutil.WriteFile(f, p.Body, 0600)
}

func load(title string) (*Page, error) {
	f := title + ".txt"
	body,err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return &Page{Title:title, Body:body}, nil
}

func view(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/test/"):]
	p, _ := load(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	p := Page{Title:"Test", Body:[]byte("Welcome to the Test Page!")}
	p.save()
	http.HandleFunc("/test/", view)
	http.ListenAndServe(":8000", nil)
}