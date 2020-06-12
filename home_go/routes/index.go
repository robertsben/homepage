package routes

import "net/http"

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

type ResponseConfig struct {
	TemplateName     string
	TemplateDataName string
	TemplateMarshal  interface{}
	Path             string
}

func IndexRouteResponder(r *http.Request) (ResponseConfig, *Error){
	var responseConfig = ResponseConfig{Path: "/"}
	if r.Method != "GET" {
		httpError := MethodNotAllowedError(r.Method, responseConfig)
		return responseConfig, &httpError
	}
	responseConfig.TemplateName = "index"
	responseConfig.TemplateDataName = "index"
	responseConfig.TemplateMarshal = &IndexPage{}
	return responseConfig, nil
}
