package controller

import (
	"homepage/model"
	"homepage/view"
	"net/http"
)

type aboutController struct {
	ControllerConfiguration
}

func (route aboutController) get(w http.ResponseWriter, r *http.Request, contentType string) {
	response := StaticResponse{DataFilename: "about", TemplateFilename: "index", DataModel: &model.IndexPage{}}
	StaticTemplateResponder(w, contentType, response)
}

func (route aboutController) Respond(w http.ResponseWriter, r *http.Request, contentType string) {
	switch r.Method {
	case http.MethodGet:
		route.get(w, r, contentType)
	default:
		ErrorResponder(w, contentType, MethodNotAllowedError(r.Method, route.acceptablePath))
	}
}

func AboutController() *aboutController {
	route := &aboutController{ControllerConfiguration{acceptablePath: "/about"}}
	route.SetAcceptableContentTypes(view.AllAvailableContentTypes())
	return route
}
