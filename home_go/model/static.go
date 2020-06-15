package model

import (
	"fmt"
	"io/ioutil"
	"path"
	"sync"
)

var dataCache map[string][]byte

var dataCacheOnce sync.Once

func initDataCache() {
	dataCache = make(map[string][]byte)
}

func ReadTemplateData(route string) ([]byte, error) {
	dataCacheOnce.Do(initDataCache)
	if val, exists := dataCache[route]; exists {
		return val, nil
	}
	raw, err := ioutil.ReadFile(path.Join("static", "data", fmt.Sprintf("%s.json", route)))
	if err != nil {
		return nil, err
	}
	dataCache[route] = raw
	return raw, nil
}

