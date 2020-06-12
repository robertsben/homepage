package routes

import (
	"encoding/json"
	"homepage/templating"
	"net/http"
)

type Error struct {
	Code        int		`json:"code"`
	Status      string	`json:"status"`
	Description string	`json:"description,omitempty"`
}

func ErrorFactory(code int, description string) Error {
	return Error{Code: code, Status: http.StatusText(code), Description: description}
}

func RenderError(w http.ResponseWriter, contentType string, httpError Error) {
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


