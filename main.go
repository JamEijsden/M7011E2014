package main

import "M7011E2014/api"

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/ziutek/mymysql/godrv"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// attach the standard ServeHTTP method to our handler so the http library can call it
func (fn handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// here we could do some prep work before calling the handler if we wanted to

	// call the actual handler
	response, err := fn(w, r)

	// check for errors
	if err != nil {
		log.Printf("ERROR: %v\n", err.Error)
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Message), err.Code)
		return
	}
	if response == nil {
		log.Printf("ERROR: response from method is nil\n")
		http.Error(w, "Internal server error. Check the logs.", http.StatusInternalServerError)
		return
	}

	// turn the response into JSON
	bytes, e := json.Marshal(response)
	if e != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	// send the response and log
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(bytes)
	log.Printf("%s %s %s %d", r.RemoteAddr, r.Method, r.URL, 200)
}

// a custom type that we can use for handling errors and formatting responses
type handler func(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError)

// error response struct
type handlerError struct {
	Error   error
	Message string
	Code    int
}

func main() {
	// command line flags
	port := flag.Int("port", 8888, "port to serve on")
	dir := flag.String("directory", "web/", "directory of web files")
	flag.Parse()

	// connect to database
	//	connect()

	// handle all requests by serving a file of the same name
	fs := http.Dir(*dir)
	fileHandler := http.FileServer(fs)

	// setup routes
	router := mux.NewRouter()
	router.Handle("/", http.RedirectHandler("/static/", 302))

	// Handlers for Users
	router.Handle("/users", handler(api.ListAllUsers)).Methods("GET")
	// hämta ut infon för att lägga till ny
	router.Handle("/users", handler(api.AddUser)).Methods("POST")
	router.Handle("/users/{id}", handler(api.GetUser)).Methods("GET")
	router.Handle("/users/{id}", handler(api.RemoveUser)).Methods("DELETE")
	// hämta alla bilder en användare har laddat upp
	router.Handle("/users/picture/{id}", handler(api.RetriveUserPictures)).Methods("GET")

	// Handlers for stairs
	router.Handle("/stair", handler(api.AddStair)).Methods("POST")
	router.Handle("/stair/{id}", handler(api.GetStair)).Methods("GET")
	router.Handle("/stairs", handler(api.GetAllStairs)).Methods("GET")
	// Get all stairs a user have added..
	router.Handle("/stairs/{id}", handler(api.GetUserStairs)).Methods("GET")
	//Get alla pictures for a stair
	router.Handle("/stair/picture/{id}", handler(api.GetriveStairPictures)).Methods("GET")

	// handlers for comments
	router.Handle("/comment", handler(api.AddComment)).Methods("POST")
	router.Handle("/comment/{id}", handler(api.GetComments)).Methods("GET")

	// Handlers for pictures
	router.Handle("/picture", handler(api.AddPicture)).Methods("POST")
	router.Handle("/picture/{id}", handler(api.GetPicture)).Methods("GET")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileHandler))
	http.Handle("/", router)

	log.Printf("Running on port %d\n", *port)

	addr := fmt.Sprintf("192.168.1.6:%d", *port)
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
