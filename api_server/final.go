package main

import (
	"encoding/json"
	"flag"
	"fmt"
	//"io/ioutil"
	"log"
	//"math"
	"database/sql"
	"github.com/ziutek/mymysql/godrv"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// error response struct
type handlerError struct {
	Error   error
	Message string
	Code    int
}

// user struct

type User struct {
	UserID    uint64 `json:"userID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// GLOBAL VARIABLE FOR CONNECTING TO DB
//var db *sql.DB

// connect to db using standard Go database/sql API
// use whatever database/sql driver you wish
db, err := sql.Open("mymysql", "tcp:localhost:3306*mydb/myuser/mypassword")

func connect() {
	username := "root"
	password := "M7017E"
	db = mysql.New("tcp", "127.0.0.1:3306", user, password)

	err := db.Connect()
	if err != nil {
		panic(err)
	}
}

// a custom type that we can use for handling errors and formatting responses
type handler func(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError)

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
	w.Write(bytes)
	log.Printf("%s %s %s %d", r.RemoteAddr, r.Method, r.URL, 200)
}

/*
	List all users in the db

*/
func listUsers(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	var hej string
	hej = string("list all users for the db")
	//return books, nil
	return hej, nil
}

/*
	Get a user from the db

*/
func getUser(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	mux.Vars(r)["id"] //grabs variables from the path
	param := mux.Vars(r)["id"]
	row, res, err := db.QueryRow("select * from users where userID =?", param)
	if err == sql.ErrNoRows {
		log.Printf("No user with that ID")
	}

	if err != nil {
		panic(err)
	}
	user := new (User);
	user.UserID = row[0]
	user.FirstName = row[1]
	user.LastName = row[2]




	//returnable := json.Marshal(user)

	return user, nil
}

/*
	ADD USER TO DB

*/
func addUser(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {

	returnable := string("adduser maby?")
	return returnable, nil
}

/*
	Remove user from DB

*/

func removeUser(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	param := mux.Vars(r)["id"]
	id, e := strconv.Atoi(param)
	if e != nil {
		return nil, &handlerError{e, "Id should be an integer", http.StatusBadRequest}
	}
	fmt.Println(id)
	// this is jsut to check to see if the book exists

	returnable := string("removeUser")
	return returnable, nil
}

func main() {
	// command line flags
	port := flag.Int("port", 8888, "port to serve on")
	dir := flag.String("directory", "web/", "directory of web files")
	flag.Parse()

	// connect to database
	connect()

	// handle all requests by serving a file of the same name
	fs := http.Dir(*dir)
	fileHandler := http.FileServer(fs)

	// setup routes
	router := mux.NewRouter()
	router.Handle("/", http.RedirectHandler("/static/", 302))
	router.Handle("/users", handler(listUsers)).Methods("GET")
	// hämta ut infon för att lägga till ny
	router.Handle("/users/{info}", handler(addUser)).Methods("POST")
	router.Handle("/users/{id}", handler(getUser)).Methods("GET")
	router.Handle("/users/{id}", handler(removeUser)).Methods("DELETE")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileHandler))
	http.Handle("/", router)

	// bootstrap some data
	//books = append(books, book{"Ender's Game", "Orson Scott Card", getNextId()})
	//books = append(books, book{"Code Complete", "Steve McConnell", getNextId()})
	//books = append(books, book{"World War Z", "Max Brooks", getNextId()})
	//books = append(books, book{"Pragmatic Programmer", "David Thomas", getNextId()})

	log.Printf("Running on port %d\n", *port)

	addr := fmt.Sprintf("192.168.1.230:%d", *port)
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
