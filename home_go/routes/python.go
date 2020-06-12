package routes

import "net/http"

func PythonRouteResponder(r *http.Request) (ResponseConfig, *Error){
	var responseConfig = ResponseConfig{Path: "/python"}
	if r.Method != "GET" {
		httpError := MethodNotAllowedError(r.Method, responseConfig)
		return responseConfig, &httpError
	}
	responseConfig.TemplateName = "index"
	responseConfig.TemplateDataName = "python"
	responseConfig.TemplateMarshal = &IndexPage{}
	return responseConfig, nil
}
