package main

import (
	"fmt"
	"homepage/routes"
	"homepage/templating"
	"net/http"
	"strings"
)

func getAcceptOptions(r *http.Request) []string {
	accept := r.Header.Get("Accept")
	possible := strings.Split(accept,",")
	options := make([]string, 0)
	for _, potential := range possible {
		option := strings.TrimSpace(potential)
		if option != "*/*" && option != "text/*" {
			options = append(options, option)
		}
	}
	if len(options) > 0 {
		return options
	}
	return []string{"text/html"}
}

func acceptableContentType(options []string) (string, error) {
	for _, option := range options {
		if _, ok := templating.MimeExtentionMap[option]; ok {
			return option, nil
		}
	}
	return "", fmt.Errorf("no acceptable content types found in %s", options)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	contentTypeOptions := getAcceptOptions(r)
	defaultContentType, err := acceptableContentType(contentTypeOptions)
	if err != nil {
		routes.RenderError(w, "text/html", routes.ErrorFactory(http.StatusNotAcceptable, err.Error()))
		return
	}
	if r.Method != "GET" {
		routes.RenderError(w, defaultContentType, routes.ErrorFactory(http.StatusMethodNotAllowed, fmt.Sprintf("%s method is not allowed for route %s", r.Method, r.URL.Path)))
		return
	}
	if r.URL.Path != "/" {
		routes.RenderError(w, defaultContentType, routes.ErrorFactory(http.StatusNotFound, "Route not found"))
		return
	}
	routes.RenderIndex(w, defaultContentType)
}
