package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./pages/assets"))))

	http.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("./pages/vendor"))))

	// http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	if r.URL.Path == "" || r.URL.Path == "/" {
	// 		http.ServeFile(w, r, "./index.html")
	// 	} else {
	// 		http.ServeFile(w, r, "./pages/"+r.URL.Path+".html")
	// 	}
	// }))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!\n")
	}))

	http.Handle("/index", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "" || r.URL.Path == "/" {
			http.ServeFile(w, r, "./index.html")
		} else {
			http.ServeFile(w, r, "./index.html")
		}
	}))

	// fs := http.FileServer(http.Dir("/"))
	// http.Handle("/", fs)

	// Build the Server
	log.Printf("Listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
