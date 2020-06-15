package controller

import (
	"homepage/model"
	"homepage/view"
	"net/http"
)

type pythonController struct {
	ControllerConfiguration
}

func (route pythonController) get(w http.ResponseWriter, r *http.Request, contentType string) {
	response := StaticResponse{DataFilename: "python", TemplateFilename: "index", DataModel: &model.IndexPage{}}
	StaticTemplateResponder(w, contentType, response)
}

func (route pythonController) Respond(w http.ResponseWriter, r *http.Request, contentType string) {
	switch r.Method {
	case http.MethodGet:
		route.get(w, r, contentType)
	default:
		ErrorResponder(w, contentType, MethodNotAllowedError(r.Method, route.acceptablePath))
	}
}

func PythonController() *pythonController {
	route := &pythonController{ControllerConfiguration{acceptablePath: "/python"}}
	route.SetAcceptableContentTypes(view.AllAvailableContentTypes())
	return route
}
