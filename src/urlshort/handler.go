package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

type Mapping struct {
	Path string `yaml:"path" json:"path"`
	Url  string `yaml:"url" json:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	mux := http.NewServeMux()
	for k, v := range pathsToUrls {
		mux.HandleFunc(k, buildHandler(v))
	}
	mux.HandleFunc("/", fallback.ServeHTTP)
	return mux.ServeHTTP
}

func buildHandler(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, http.StatusFound)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathsToUrls, err := parseMarkup(yaml.Unmarshal, yml)
	if err != nil {
		return nil, err
	}
	return MapHandler(pathsToUrls, fallback), nil
}

func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathsToUrls, err := parseMarkup(json.Unmarshal, jsonBytes)
	if err != nil {
		return nil, err
	}
	return MapHandler(pathsToUrls, fallback), nil
}

func parseMarkup(unmarshalFunc func([]byte, interface{}) error, in []byte) (map[string]string, error) {
	retMap := make(map[string]string)
	var mappingData []Mapping
	err := unmarshalFunc(in, &mappingData)
	if err != nil {
		return nil, err
	}
	for _, m := range mappingData {
		retMap[m.Path] = m.Url
	}
	return retMap, nil
}
