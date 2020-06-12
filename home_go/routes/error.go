package routes

import (
	"homepage/templating"
	"net/http"
)

type Error struct {
	Code        int
	Status      string
	Description string
}

func ErrorFactory(code int, description string) Error {
	return Error{Code: code, Status: http.StatusText(code), Description: description}
}

func RenderError(w http.ResponseWriter, mime string, err Error) {
	tmpl, _ := templating.ReadTemplate("error", mime)
	if _err := tmpl.Execute(w, err); _err != nil {
		http.Error(w, _err.Error(), http.StatusInternalServerError)
	}
}


