package view

import (
	"fmt"
	"path"
	"sync"
	"text/template"
)

const (
	JSON  = "application/json"
	HTML  = "text/html"
	PLAIN = "text/plain"
)

func AllAvailableContentTypes() []string {
	return []string{
		JSON,
		HTML,
		PLAIN,
	}
}

var MimeExtentionMap = map[string]string{
	PLAIN: "txt",
	HTML:  "html",
	JSON:  "json",
}

var templateCache map[string]map[string]*template.Template

var templateCacheOnce sync.Once

func initTemlplateCache() {
	templateCache = map[string]map[string]*template.Template{}
	for k := range MimeExtentionMap {
		templateCache[k] = make(map[string]*template.Template)
	}
}

func ReadTemplate(name string, mime string) (*template.Template, error) {
	templateCacheOnce.Do(initTemlplateCache)
	if val, exists := templateCache[mime][name]; exists {
		return val, nil
	}
	ext := MimeExtentionMap[mime]
	fp := path.Join("static", "templates", ext, fmt.Sprintf("%s.%s", name, ext))
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		return nil, err
	}
	templateCache[mime][name] = tmpl
	return tmpl, nil
}
