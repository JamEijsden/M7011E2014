package main

import (
	"fmt"
	"html/template"
	"net/http" //package for http based web programs
	//"net/url"
	"path/filepath"
)

var templates string

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside handler")
	// Parse our root.html template
	if t, err := template.ParseFiles(filepath.Join(templates, "index.html")); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if t, err := template.ParseFiles(filepath.Join(templates, "about.html")); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}

}
func tomteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Tomte")
	remPart := r.URL.Path[len("/tomte"):]
	fmt.Fprintf(w, "Yes? This is %s", remPart)

}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside test")

	if t, err := template.ParseFiles(filepath.Join(templates, "test.html")); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}

}

func main() {
	templates = ""
	http.HandleFunc("/", handler)
	http.HandleFunc("/about/", aboutHandler)
	http.HandleFunc("/test/", testHandler)
	http.ListenAndServe("localhost:9999", nil) // listen for connections at port 9999 on the local machine
}
