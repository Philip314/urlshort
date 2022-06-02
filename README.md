# URL Short

## About

Application that redirects shortened URL to the original URL.
Shortened URL is stored in the formats of Map, YAML, and JSON.

I created this from following [Gophercises](https://gophercises.com/ "Gophercises") to learn Go.

## Getting Started

1. Clone repo
2. Go into the project folder with command line interface
3. Run project with `go run main.go`
4. Go to `localhost:8080` in a web browser to see the index or with a valid path to be redirected to the original URL

## Command-line Arguments

```
$ go run main.go -h
  -json string
        JSON file containing paths and URLs (default "link.json")
  -type string
        Type to decode URLS from (map, yaml, json) (default "json")
  -yaml string
        YAML file containing paths and URLs (default "link.yaml")
```

## File Formats

It is expecting the YAML file to be in the following format:

```
- path: /some-path
  url: https://www.someurl.com/
```

For JSON:

```
[{
    "path": "/some-path",
    "url": "https://www.someurl.com/"
}]

```