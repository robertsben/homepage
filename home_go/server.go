package main

import (
	"encoding/json"
	"fmt"
	"homepage/routes"
	"homepage/templating"
	"log"
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

func RenderError(w http.ResponseWriter, contentType string, httpError routes.Error) {
	if contentType == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		resp, _ := json.Marshal(httpError)
		w.WriteHeader(httpError.Code)
		w.Write(resp)
		return
	}
	tmpl, _ := templating.ReadTemplate("error", contentType)
	w.WriteHeader(httpError.Code)
	if _err := tmpl.Execute(w, httpError); _err != nil {
		http.Error(w, _err.Error(), http.StatusInternalServerError)
	}
}

func RenderRoute(w http.ResponseWriter, contentType string, config routes.ResponseConfig) {
	rawData, err := templating.ReadTemplateData(config.TemplateDataName)
	if err != nil {
		RenderError(w, contentType, routes.ErrorFactory(http.StatusInternalServerError, err.Error()))
	}
	err = json.Unmarshal(rawData, config.TemplateMarshal)
	if err != nil {
		RenderError(w, contentType, routes.ErrorFactory(http.StatusInternalServerError, err.Error()))
	}
	if contentType == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		resp, _ := json.Marshal(config.TemplateMarshal)
		w.Write(resp)
		return
	}
	tmpl, err := templating.ReadTemplate(config.TemplateName, contentType)
	if err != nil {
		RenderError(w, contentType, routes.ErrorFactory(http.StatusInternalServerError, err.Error()))
	}
	if err := tmpl.Execute(w, config.TemplateMarshal); err != nil {
		RenderError(w, contentType, routes.ErrorFactory(http.StatusInternalServerError, err.Error()))
	}
}

func RoutingHandler(responder func(*http.Request) (routes.ResponseConfig, *routes.Error)) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		contentTypeOptions := getAcceptOptions(r)
		defaultContentType, err := acceptableContentType(contentTypeOptions)
		if err != nil {
			RenderError(w, "text/html", routes.ErrorFactory(http.StatusNotAcceptable, err.Error()))
			return
		}
		responseConfig, httpErr := responder(r)
		if httpErr != nil {
			RenderError(w, defaultContentType, *httpErr)
			return
		}
		if r.URL.Path != responseConfig.Path {
			RenderError(w, defaultContentType, routes.ErrorFactory(http.StatusNotFound, "Route not found"))
			return
		}
		RenderRoute(w, defaultContentType, responseConfig)
	}
}

func Serve() {
	http.HandleFunc("/", RoutingHandler(routes.IndexRouteResponder))
	http.HandleFunc("/python", RoutingHandler(routes.PythonRouteResponder))
	http.HandleFunc("/about", RoutingHandler(routes.AboutRouteResponder))
	log.Fatal(http.ListenAndServe(":8000", nil))
}



