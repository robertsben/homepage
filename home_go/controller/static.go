package controller

import (
	"encoding/json"
	"homepage/model"
	"homepage/view"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", view.JSON)
	resp, _ := json.Marshal(data)
	_, err := w.Write(resp)
	if err != nil {
		ErrorResponder(w, view.JSON, ErrorFactory(http.StatusInternalServerError, err.Error()))
	}
}

func templatedResponse(w http.ResponseWriter, contentType string, response StaticResponse) {
	tmpl, err := view.ReadTemplate(response.TemplateFilename, contentType)
	if err != nil {
		ErrorResponder(w, contentType, ErrorFactory(http.StatusInternalServerError, err.Error()))
	}
	if err := tmpl.Execute(w, response.DataModel); err != nil {
		ErrorResponder(w, contentType, ErrorFactory(http.StatusInternalServerError, err.Error()))
	}
}

func StaticTemplateResponder(w http.ResponseWriter, contentType string, response StaticResponse) {
	rawData, err := model.ReadTemplateData(response.DataFilename)
	if err != nil {
		ErrorResponder(w, contentType, ErrorFactory(http.StatusInternalServerError, err.Error()))
	}

	err = json.Unmarshal(rawData, response.DataModel)
	if err != nil {
		ErrorResponder(w, contentType, ErrorFactory(http.StatusInternalServerError, err.Error()))
	}

	switch contentType {
	case view.JSON:
		jsonResponse(w, response.DataModel)
	default:
		templatedResponse(w, contentType, response)
	}
}