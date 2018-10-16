package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]
		ext := filepath.Ext(path)
		mime := "text/html"
		switch ext {
		case ".png":
			mime = "img/png"
		case ".js":
			mime = "application/javascript"
		case ".css":
			mime = "text/css"
		}

		w.Header().Add("Content-Type", mime)

		log.Println(fmt.Sprintf("Path: %s, ext/mime: %s/%s, %+v", path, ext, mime, w.Header()))
		http.ServeFile(w, r, path)
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	go log.Fatal(http.ListenAndServeTLS(":5000", "./cert/cert.pem", "./cert/key.pem", nil))
	log.Fatal(http.ListenAndServe(":5000", nil))
}
