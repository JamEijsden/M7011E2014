package main

import (
	"fmt"
	"html/template"
	"net/http" //package for http based web programs
	//"net/url"
	"os"
	//"path/filepath"
)

var templates string

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside handler")
	// Parse our root.html template
	if t, err := template.ParseFiles("_testBase.html", "templates/index.html"); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside about")
	if t, err := template.ParseFiles("_testBase.html", "templates/about.html"); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}

}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside test")

	if t, err := template.ParseFiles("_testBase.html", "templates/test.html"); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}

}

func backHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {

	//fs := justFilesFilesystem{http.Dir("/static")}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", handler)
	http.HandleFunc("/about/", aboutHandler)
	http.HandleFunc("/test/", testHandler)

	http.ListenAndServe("localhost:9999", nil) // listen for connections at port 9999 on the local machine
}

// HIDE DICRECTORY
type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

type neuteredReaddirFile struct {
	http.File
}

func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}
