package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"math/big"
	"net/http"
)

type User struct {
	UserID    big.Int `json:"userID"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
}

func main() {
	m := mux.NewRouter()
	// fetch all users
	m.Handle("/users", handler(getAllUsers)).Methods("GET")

	//fetch a user
	m.Handle("/user/{key}", handler(getUser)).Methods("GET")

	//delete user
	m.Handle("/user/{key}", handler(deleteUser)).Methods("DELETE")

	//add user
	m.Handle("/user/{key}", handler(addUser)).Methods("PUT")

	http.Handle("/", m)

	http.ListenAndServe("localhost:8080", nil)

	// everything else fails
	//m.HandleFunc("/{path:.*}", g)
}
func deleteUser(rw http.ResponseWriter, req *http.Request) {

}

func getAllUsers(w http.ResponseWriter, r *http.Request) string {
	//vars := mux.Vars(r)
	//id := vars["key"]
	//fmt.Printf(w, "Page for displaying users")

	return string("all the uses")
}

func getUser(rw http.ResponseWriter, req *http.Request) string {
	vars := mux.Vars(req)
	id := vars["key"]

	/*
		retrive the user from the db
		and put in a user struct
	*/
	var usr User
	//add the other lines
	// then
	ans := usr.jsonResponseUser()
	//fmt.Printf(rw, string(ans))
	return string(ans)

}

func addUser(rw http.ResponseWriter, req *http.Request) {
	// read body
	data, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}
	var usr User
	err = json.Unmarshal(data, &usr)
	if err != nil {
		http.Error(rw, err.Error(), 500)
	}

	//vars := mux.Vars(req)
	//id := vars["key"]
	//d := user
	//d.userID = id

	/*
		here we want to add this to the database
		using usr ass the parameters
	*/

}

func (u *User) jsonResponseUser() []byte {
	jsonencoded, err := json.Marshal(u)
	if err != nil {
		fmt.Println("error int jsonAnswerUser")
		return nil
	}
	return jsonencoded
}

func fixParsingRequest(rw http.ResponseWriter, req *http.Request) {
	//Reading the body
	data, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}

	// Unmarshalling the req param
	var usr User
	err = json.Unmarshal(data, &usr)
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}
	return usr

}