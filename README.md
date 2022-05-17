# URL Short

## About
Application that redirects shortened URL to the original URL.
Shortened URL is stored in the formats of Map, YAML, and JSON.

I created this from following [Gophercises](https://gophercises.com/ "Gophercises") to learn Go.

## Getting Started
1. Pull repo
2. Go into the project folder with command line interface
3. Run project with `go run main.go`
4. Go to `localhost:8080` in a web browser to see the index or with a valid path to be redirected to the original URL

You can specify a YAML file to read from with the `-yaml` flag.
It is expecting to be in the format of 

```
- path: /some-path
  url: https://www.someurl.com/
```

Flag details can be found by running `go run main.go -help`.