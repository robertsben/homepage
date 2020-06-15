package controller

import (
	"net/http"
)

type Controller interface {
	AcceptableContentTypes() map[string]struct{}
	AcceptablePath() string
	Respond(w http.ResponseWriter, r *http.Request, contentType string)
}

type StaticResponse struct {
	DataFilename string
	TemplateFilename string
	DataModel interface{}
}

type ControllerConfiguration struct {
	acceptableContentTypes map[string]struct{}
	acceptablePath string
}

func (c *ControllerConfiguration) SetAcceptableContentTypes(contentTypes []string) {
	c.acceptableContentTypes = make(map[string]struct{}, len(contentTypes))
	for _, contentType := range contentTypes {
		c.acceptableContentTypes[contentType] = struct{}{}
	}
}

func (c *ControllerConfiguration) AcceptableContentTypes() map[string]struct{} {
	return c.acceptableContentTypes
}

func (c *ControllerConfiguration) AcceptablePath() string {
	return c.acceptablePath
}