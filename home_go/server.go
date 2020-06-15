package main

import (
	"fmt"
	"homepage/controller"
	"homepage/view"
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

func controllerAllowedContentType(options []string, c controller.Controller) (string, error) {
	allowedContentTypes := c.AcceptableContentTypes()
	for _, option := range options {
		if _, ok := allowedContentTypes[option]; ok {
			return option, nil
		}
	}
	return "", fmt.Errorf("no acceptable content types found in %s", options)
}

func RoutingHandler(c controller.Controller) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		contentTypeOptions := getAcceptOptions(r)
		contentType, err := controllerAllowedContentType(contentTypeOptions, c)
		if err != nil {
			controller.ErrorResponder(w, view.HTML, controller.ErrorFactory(http.StatusNotAcceptable, err.Error()))
			return
		}
		if r.URL.Path != c.AcceptablePath() {
			controller.ErrorResponder(w, contentType, controller.ErrorFactory(http.StatusNotFound, "Route not found"))
			return
		}
		c.Respond(w, r, contentType)
	}
}

func Serve() {
	controllers := []controller.Controller{
		controller.IndexController(),
		controller.AboutController(),
		controller.PythonController(),
	}

	for _, c := range controllers {
		http.HandleFunc(c.AcceptablePath(), RoutingHandler(c))
	}

	log.Fatal(http.ListenAndServe(":8000", nil))
}



