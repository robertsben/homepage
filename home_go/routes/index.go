package routes

import (
	"encoding/json"
	"homepage/templating"
	"net/http"
)

type Body struct {
	Heading    string   `json:"heading"`
	Paragraphs []string `json:"paragraphs"`
}

type Link struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type IndexPage struct {
	Title   string `json:"title"`
	Heading string `json:"heading"`
	Main    Body   `json:"main"`
	Links   []Link `json:"links"`
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
	if contentType == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		resp, _ := json.Marshal(data)
		w.Write(resp)
		return
	}
	tmpl, err := templating.ReadTemplate("index", contentType)
	if err != nil {
		RenderError(w, contentType, ErrorFactory(http.StatusInternalServerError, err.Error()))
	}
	if err := tmpl.Execute(w, data); err != nil {
		RenderError(w, contentType, ErrorFactory(http.StatusInternalServerError, err.Error()))
	}
}
