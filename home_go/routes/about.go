package routes

import "net/http"

func AboutRouteResponder(r *http.Request) (ResponseConfig, *Error){
	var responseConfig = ResponseConfig{Path: "/about"}
	if r.Method != "GET" {
		httpError := MethodNotAllowedError(r.Method, responseConfig)
		return responseConfig, &httpError
	}
	responseConfig.TemplateName = "index"
	responseConfig.TemplateDataName = "about"
	responseConfig.TemplateMarshal = &IndexPage{}
	return responseConfig, nil
}
