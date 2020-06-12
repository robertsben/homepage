package routes

import (
	"encoding/json"
	"homepage/templating"
	"net/http"
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

func RenderIndex(w http.ResponseWriter, contentType string) {
	var data IndexPage
	rawData, err := templating.ReadTemplateData("index")
	if err != nil {
		RenderError(w, contentType, ErrorFactory(http.StatusInternalServerError, err.Error()))
	}
	err = json.Unmarshal(rawData, &data)
	if err != nil {
		RenderError(w, contentType, ErrorFactory(http.StatusInternalServerError, err.Error()))
	}
	tmpl, err := templating.ReadTemplate("index", contentType)
	if err != nil {
		RenderError(w, contentType, ErrorFactory(http.StatusInternalServerError, err.Error()))
	}
	if err := tmpl.Execute(w, data); err != nil {
		RenderError(w, contentType, ErrorFactory(http.StatusInternalServerError, err.Error()))
	}
}

