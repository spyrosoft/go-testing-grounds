package main

import (
	"fmt"
	"net/http"
	"strings"
	"path"
)

type StaticHandler struct {
	http.Dir
}

func (sh *StaticHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path
	if !strings.HasPrefix(filePath, "/") {
		filePath = "/" + filePath
		r.URL.Path = filePath
	}
	cleanpath := path.Clean(filePath)
	f, err := sh.Open(cleanpath)
	if err != nil {
		fmt.Fprint(w, "placeholder for 404 html")
		return
	}
	defer f.Close()
	
	d, err := f.Stat()
	if err != nil {
		//return 404
		return
	}
	
	if d.IsDir() {
		//Check for index.html or 404
		return
	}
	
	http.ServeContent(w, r, d.Name(), d.ModTime(), f)
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/example/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "welcome example")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}

// The goal: API 1st, static files second, 404 last

func main() {
	http.HandleFunc("/example/", exampleHandler)
	staticHandler := StaticHandler{"/tmp/delete"}
	http.HandleFunc("/", staticHandler.ServeHttp)
	http.ListenAndServe(":12345", nil)
}