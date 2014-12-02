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
	if t, err := template.ParseFiles("_index.html", "templates/index.html"); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if t, err := template.ParseFiles("_Base.html", "templates/about.html"); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}

}

func userHandler(w http.ResponseWriter, r *http.Request) {
	if t, err := template.ParseFiles("_Base.html", "templates/user.html"); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}

}

func mapHandler(w http.ResponseWriter, r *http.Request) {

	if t, err := template.ParseFiles("_Base.html", "templates/map.html"); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}

}

func addStairsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("YOLO POPUP")

	if t, err := template.ParseFiles("_add.html", "templates/addstairs.html"); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}

}

func stairsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("YOLO POPUP")

	if t, err := template.ParseFiles("_stair.html", "templates/stairs.html"); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}

}

func photosHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("YOLO POPUP")

	if t, err := template.ParseFiles("_stair.html", "templates/stairs.html"); err != nil {
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
	http.HandleFunc("/home/", homeHandler)
	http.HandleFunc("/user/", userHandler)
	http.HandleFunc("/map/", mapHandler)
	http.HandleFunc("/map/addStairs/", addStairsHandler)
	http.HandleFunc("/map/stairs/", stairsHandler)
	http.HandleFunc("/map/stairs/photos", photosHandler)

	var i int
	fmt.Println("Run server on:\n1. localhost:9999\n2. 192.168.1.230:9999\nChoose a connection(1 or 2).")
	fmt.Scan(&i)
	if i == 1 {
		fmt.Println("Server running on localhost:9999")
		http.ListenAndServe("localhost:9999", nil)
	} else if i == 2 {
		fmt.Println("Server running on 192.168.1.6:9999")
		http.ListenAndServe("192.168.1.6:9999", nil)
	} else {
		fmt.Println("Not a valid option")
	}

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
