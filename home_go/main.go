package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
)

type Body struct {
	Heading string
}

type Link struct {
	Path string
	Name string
}

type Paragraph struct {
	Body string
}

type MainPage struct {
	Title string
	Heading string
	Main Body
	Links []Link
	Paragraphs []Paragraph
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	data := MainPage{
		Title: "test page",
		Heading: "It's a me, a test a page",
		Main: Body{Heading: "ummm yep"},
		Links: []Link{{Path: "/", Name: "Home"}, {Path: "/test", Name: "self"}, {Path: "/about", Name: "About"}},
		Paragraphs: []Paragraph{{Body: "Yeah whatever"}, {Body: "somethign something whatever"}},
	}
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("Hello, %s world!", r.URL.String()))
}

func main() {
	http.HandleFunc("/test", templateHandler)
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
