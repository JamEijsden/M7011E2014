package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"is/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"

	// måste importera vårat själva api program också
)

type handlerError struct {
	Error   error
	Message string
	Code    int
}

type users struct {
}

type handler func(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError)

func (fn handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// calling handler
	response, err := fn(w, r)

	//errorCheck
	if err != nil {
		log.Printf("ERROR: %v\n", err.Error)
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Message), err.Code)
		return
	}

	if response == nil {
		log.Printf("ERROR: response form method is nill \n")
	}

	//make response as json
	bytes, e := json.Marshal(response)
	if e != nil {
		http.Error(w, "ERROE marshalling JSON", http.StatusInternalServerError)
		return
	}

	//send json
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
	log.Printf("%s %s %s %d", r.RemoteAddr, r.Method, r.URL, 200)

}

func main() {
	fs := http.Dir(*dir)
	fileHandler := http.FileServer(fs)

	//routes setup
	router := mux.NewRouter()
	//retrive all users
	router.Handle("/users", handler(listAllUser)).Methods("GET")
	// add a user
	router.Handle("/user", handler(addUser)).Methods("POST")
	// retrive a user
	router.Handle("/user/{id}", handler(listUser)).Methods("GET")
	//update user info
	router.Handle("/user/{id}", handler(updateUser)).Methods("POST")
	//deletes a user
	router.Handle("/user/{id}", handler(deleteUser)).Methods("DELETE")
	router.Handle("/", router)
	// servses files on path
	router.Path("/response/").Handler(http.StripPrefix("/response", fileHandler))
	http.ListenAndServe("localhost:1337", nil)
}



func listallusers(rw http.ResponseWriter, req *http.Request (interface{}, *handlerError)) {
	param := mux.Vars(r)
}

/*func main() {
	mux := http.NewServeMux()

	http.HandleFunc("GET", "/users", user.GetALLUser)
	http.HandleFunc("POST", "/users", user.PostUser)
	http.HandleFunc("GET", "/users:id", user.GetUser)
	http.HandleFunc("PUT", "/users:id", user.PutUser)
	http.HandleFunc("DELETE", "/user/:id", user.DeleteUser)
	http.ListenAndServe("localhost:1337", nil)

}
*/
