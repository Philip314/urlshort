package handlers

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		if url, ok := pathsToUrls[path]; ok {
			http.Redirect(writer, request, url, http.StatusFound)
		}
		fallback.ServeHTTP(writer, request)
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
func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	yamlSlice, err := createYamlType(yamlBytes)
	if err != nil {
		return nil, err
	}
	yamlMap := parseYamlTypeToMap(yamlSlice)
	return MapHandler(yamlMap, fallback), nil
}

type yamlType struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func createYamlType(yamlByte []byte) ([]yamlType, error) {
	var yamlSlice []yamlType
	err := yaml.Unmarshal(yamlByte, &yamlSlice)
	if err != nil {
		return nil, err
	}
	return yamlSlice, err
}

func parseYamlTypeToMap(yamlSlice []yamlType) map[string]string {
	mapOutput := make(map[string]string)
	for _, v := range yamlSlice {
		mapOutput[v.Path] = v.URL
	}
	return mapOutput
}

func JsonHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	jsonType, err := createJsonType(json)
	if err != nil {
		return nil, err
	}
	return MapHandler(parseJsonTypeToMap(jsonType), fallback), nil
}

type JsonType struct {
	Path string
	URL  string
}

func createJsonType(jsonByte []byte) ([]JsonType, error) {
	var jsonType []JsonType
	err := json.Unmarshal(jsonByte, &jsonType)
	if err != nil {
		return nil, err
	}
	return jsonType, nil
}

func parseJsonTypeToMap(jsonByte []JsonType) map[string]string {
	mapOutput := make(map[string]string)
	for _, v := range jsonByte {
		mapOutput[v.Path] = v.URL
	}
	return mapOutput
}
