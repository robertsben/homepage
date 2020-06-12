package routes

import (
	"fmt"
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

func MethodNotAllowedError(method string, config ResponseConfig) Error {
	return ErrorFactory(http.StatusMethodNotAllowed, fmt.Sprintf("%s method is not allowed for route %s", method, config.Path))
}
