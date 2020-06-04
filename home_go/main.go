package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

type Body struct {
	Heading    string
	Paragraphs []string
}

type Link struct {
	Url  string
	Name string
}

type IndexPage struct {
	Title   string
	Heading string
	Main    Body
	Links   []Link
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	raw, err := ioutil.ReadFile(path.Join("data", "index.json"))
	var data IndexPage
	err = json.Unmarshal(raw, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
