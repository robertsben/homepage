package controller

import (
	"homepage/model"
	"homepage/view"
	"net/http"
)

type indexController struct {
	ControllerConfiguration
}

func (route indexController) get(w http.ResponseWriter, r *http.Request, contentType string) {
	response := StaticResponse{DataFilename: "index", TemplateFilename: "index", DataModel: &model.IndexPage{}}
	StaticTemplateResponder(w, contentType, response)
}

func (route indexController) Respond(w http.ResponseWriter, r *http.Request, contentType string) {
	switch r.Method {
	case http.MethodGet:
		route.get(w, r, contentType)
	default:
		ErrorResponder(w, contentType, MethodNotAllowedError(r.Method, route.acceptablePath))
	}
}

func IndexController() *indexController {
	route := &indexController{ControllerConfiguration{acceptablePath: "/"}}
	route.SetAcceptableContentTypes(view.AllAvailableContentTypes())
	return route
}
