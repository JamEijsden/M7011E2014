package main

import (
<<<<<<< HEAD
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

func hejHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Hej")
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

func main() {
	templates = ""
	http.HandleFunc("/", handler)
	http.HandleFunc("/hej/", hejHandler)
	http.ListenAndServe("localhost:9999", nil) // listen for connections at port 9999 on the local machine
=======
	"github.com/go-martini/martini"
)

/*
Kollade in martini lite. Värkar ganska trevligt.
för att få in martini i sitt golang så behöver
man hämta ner det.

Detta görs lättast med att skriva
go get github.com/go-martini/martini
från terminalen/ cmd
*/

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
>>>>>>> origin/master
}
