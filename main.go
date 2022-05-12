package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/philip314/urlshort/handlers"
)

func main() {
	fileRead := true

	yamlFilename := flag.String("yaml", "link.yaml", "YAML file containing paths and URLs")

	flag.Parse()

	file, err := os.Open(*yamlFilename)
	if err != nil {
		fmt.Println("Error opening file")
		fileRead = false
	}

	fileSize, err := getFileSize(file)
	if err != nil {
		fmt.Println("Error getting filesize")
		fileRead = false
	}

	data := make([]byte, fileSize)
	count, err := file.Read(data)
	if err != nil || count == 0 {
		fmt.Println("Error reading file")
		fileRead = false
	}

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := handlers.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	// Fallback to default data if reading file fails
	if !fileRead {
		fmt.Println("Using default data")
		data = []byte(yaml)
	}

	yamlHandler, err := handlers.YAMLHandler([]byte(data), mapHandler)
	if err != nil {
		panic(err)
	}
	_ = yamlHandler

	json := `
	[{
		"path": "/urlshort",
		"url": "https://github.com/gophercises/urlshort"
	}, {
		"path": "/urlshort-final",
		"url": "https://github.com/gophercises/urlshort/tree/solution"
	}]
`
	jsonHandler, err := handlers.JsonHandler([]byte(json), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func getFileSize(file *os.File) (int, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}
	return int(fileInfo.Size()), nil
}
