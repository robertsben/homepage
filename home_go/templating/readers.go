package templating

import (
	"fmt"
	"io/ioutil"
	"path"
	"sync"
	"text/template"
)

var MimeExtentionMap = map[string]string{
	"text/plain":       "txt",
	"text/html":        "html",
	"application/json": "json",
}

var dataCache map[string][]byte
var templateCache map[string]map[string]*template.Template

var dataCacheOnce sync.Once
var templateCacheOnce sync.Once

func initDataCache() {
	dataCache = make(map[string][]byte)
}

func initTemlpateCache() {
	templateCache = map[string]map[string]*template.Template{}
	for k := range MimeExtentionMap {
		templateCache[k] = make(map[string]*template.Template)
	}
}

func ReadTemplateData(route string) ([]byte, error) {
	dataCacheOnce.Do(initDataCache)
	if val, exists := dataCache[route]; exists {
		return val, nil
	}
	raw, err := ioutil.ReadFile(path.Join("data", fmt.Sprintf("%s.json", route)))
	if err != nil {
		return nil, err
	}
	dataCache[route] = raw
	return raw, nil
}

func ReadTemplate(name string, mime string) (*template.Template, error) {
	templateCacheOnce.Do(initTemlpateCache)
	if val, exists := templateCache[mime][name]; exists {
		return val, nil
	}
	ext := MimeExtentionMap[mime]
	fp := path.Join("templates", ext, fmt.Sprintf("%s.%s", name, ext))
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		return nil, err
	}
	templateCache[mime][name] = tmpl
	return tmpl, nil
}
