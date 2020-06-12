package main

import (
	"encoding/json"
	"html/template"
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
	fp := path.Join("templates", "html", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func main() {
	http.HandleFunc("/test", templateHandler)
	http.HandleFunc("/", DefaultHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
