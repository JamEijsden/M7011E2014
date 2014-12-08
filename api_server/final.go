package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/ziutek/mymysql/godrv"
	"io/ioutil"
	"log"
	//"math"
	"net/http"
	"strconv"
	"time"

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
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IdToken   string `json:"id"`
	Photo     string `json:"photo"`
}

//Stair struct
type Stair struct {
	Id          uint64 `json:"id"`
	Position    string `json:"position"`
	Name        string `json:"stairname"`
	User        uint64 `json:"user"`
	Photo       string `json:"photo"`
	Description string `json:"description"`
}

//Comment struct
type Comment struct {
	CommentId   uint64    `json:"commentId"`
	CommentText string    `json:"commentText"`
	CommentDate time.Time `json:"commentDate"`
	IdStair     uint64    `json:"idStair"`
	IdToken     string    `json:"idToken"`
}

// GLOBAL VARIABLE FOR CONNECTING TO DB
//var db *sql.db

// connect to db using standard Go database/sql API
// use whatever database/sql driver you wish
//db, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")

/**
func connect() {
	username := "root"
	password := "jaam"
	database := "M7011E"
	fmt.Println("HALLO")
	db = mysql.New("tcp","localhost:3306", username, password, database)
	fmt.Println("HALLO2")
	err := db.Connect()
	fmt.Println("HALLO3")
	if err != nil {
		panic(err)
	}

}
**/
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(bytes)
	log.Printf("%s %s %s %d", r.RemoteAddr, r.Method, r.URL, 200)
}

/*
	List all users in the db
	!READY FOR TESTING!
*/
func listAllUsers(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		return nil, &handlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer con.Close()

	rows, err := con.Query("select name, uid from users")
	if err != nil {
		return nil, &handlerError{err, "Error in DB", http.StatusInternalServerError}
		//log.Printf("No user with that ID")
	}

	var result []User // create an array of stairs
	var uid uint64
	var name string

	for rows.Next() {
		user := new(User)
		err = rows.Scan(&name, &uid)
		if err != nil {
			return result, &handlerError{err, "Error in DB", http.StatusInternalServerError}
		}
		user.FirstName = name
		user.UserID = uid
		result = append(result, *user)
	}

	return result, nil
}

/*
	Get a user from the db
	!DONE FOR TESTING!

*/
func getUser(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	//mux.Vars(r)["id"] grabs variables from the path
	param := mux.Vars(r)["id"]
	fmt.Println(param)
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	row, err := con.Query("select * from users where idToken =?", param)
	if err == sql.ErrNoRows {
		log.Printf("No user with that ID")
	}

	if err != nil {
		panic(err)
	}

	user := new(User)
	for row.Next() {
		var idToken string
		var uid uint64
		var name, lastname, photo string

		fmt.Println(row)

		if err := row.Scan(&uid, &name, &lastname, &idToken, &photo); err != nil {
			log.Fatal(err)
		}
		user.IdToken = idToken
		user.UserID = uid
		user.FirstName = name
		user.LastName = lastname
		user.Photo = photo
	}
	//fmt.Println(row)
	//	user.UserID = row[0]
	//	user.FirstName = row[1]
	//	user.LastName = row[2]

	//returnable := json.Marshal(user)
	//returnable := json.Marshal(user)

	return user, nil
}

/*
	ADD USER TO DB
	!DONE for TESTING!

*/
func addUser(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {

	data, e := ioutil.ReadAll(r.Body)

	fmt.Println("BEFORE UNMARSHAL" + string(data))
	if e != nil {
		fmt.Println("AJAJAJ 1111")
		fmt.Println(string(data))
		return nil, &handlerError{e, "Can't read request", http.StatusBadRequest}
	}

	// create new user called payload
	var payload User
	e = json.Unmarshal(data, &payload)

	if e != nil {
		fmt.Println("SATAN")
		fmt.Println(e)
		fmt.Println("kunde inte unmarshla detta:")
		fmt.Println(payload)
		return Stair{}, &handlerError{e, "Could'nt parse JSON", http.StatusInternalServerError}
	}
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		fmt.Println("Kunde inte öppna DB")
		return nil, &handlerError{err, "Internal server error", http.StatusInternalServerError}
	}
	defer con.Close()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Detta är token:")
	fmt.Println(payload.IdToken)
	fmt.Println("")
	fmt.Println("")
	row, _ := con.Query("select count(*) from users where idToken=?", payload.IdToken)
	var count int
	for row.Next() {
		row.Scan(&count)
	}

	if count == 1 {
		return nil, &handlerError{nil, "User already exists", http.StatusFound}

	}

	_, err = con.Exec("insert into users( name, lastname, idToken, photo) values(?,?,?,?)", payload.FirstName, payload.LastName, payload.IdToken, payload.Photo)

	if err != nil {
		fmt.Println("Kunde inte lägga till :/")
		return nil, &handlerError{err, "Error adding to DB", http.StatusInternalServerError}
	}

	return payload, nil
	//row, err := con.Query("select * from users where uid =?", param)
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

/*
	Add stair to DB
	!Done for testing!

*/
func addStair(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	data, e := ioutil.ReadAll(req.Body)

	fmt.Println("BEFORE MARSHAL " + string(data))
	if e != nil {
		fmt.Println("AJAJAJ 1111")
		fmt.Println(string(data))
		return nil, &handlerError{e, "Can't read request", http.StatusBadRequest}
	}
	var payload Stair
	e = json.Unmarshal(data, &payload)
	fmt.Print("AFTER UNMARSHAL ")
	fmt.Println(payload)
	if e != nil {
		fmt.Println("SATAN")
		fmt.Println(e)
		fmt.Println("kunde inte unmarshla detta:")
		fmt.Println(payload)
		return Stair{}, &handlerError{e, "Could'nt parse JSON", http.StatusInternalServerError}
	}
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		fmt.Println("Kunde inte öppna DB")
		return nil, &handlerError{err, "Internal server error", http.StatusInternalServerError}
	}
	defer con.Close()
	//jimmie vill ha statiskt...... fultfultfult!!
	//fixar ickestatiskt när han har bättre tankar:P
	//typ om dobbe eller dobbelina
	// eller kanske linkeboda??
	//		OOOOOOOOO
	_, err = con.Exec("insert into Stairs(position, stairname, description, uid, photo) values(?,?,?,?,?)", payload.Position, payload.Name, payload.Description, payload.User, payload.Photo)

	if err != nil {
		fmt.Println("Kunde inte lägga till :/")
		return nil, &handlerError{err, "Error adding to DB", http.StatusInternalServerError}
	}

	return payload, nil
	//row, err := con.Query("select * from users where uid =?", param)
}

/*
	Get stair from DB
	!READY FOR TESTING!
*/

func getUserStairs(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	param := mux.Vars(req)["id"]
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		return nil, &handlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer con.Close()

	row, err := con.Query("select * from Stairs where uid =?", param)
	if err == sql.ErrNoRows {
		return nil, &handlerError{err, "Error no stairs found", http.StatusBadRequest}
		//log.Printf("No user with that ID")
	}

	if err != nil {
		return nil, &handlerError{err, "Internal Error when req DB", http.StatusInternalServerError}
		//panic(err)
	}

	stair := new(Stair)
	for row.Next() {
		var position, stairname, photo, description string
		var uid, id uint64

		fmt.Println(row)

		if err := row.Scan(&id, &position, &stairname, &description, &uid, &photo); err != nil {
			return nil, &handlerError{err, "Internal Error when reading req from DB", http.StatusInternalServerError}
			//log.Fatal(err)
		}

		stair.Id = id
		stair.Name = stairname
		stair.Photo = photo
		stair.User = uid
		stair.Description = description
		stair.Position = position

	}

	return stair, nil

}
func getStair(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	param := mux.Vars(req)["id"]
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		return nil, &handlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer con.Close()

	row, err := con.Query("select * from Stairs where id =?", param)
	if err == sql.ErrNoRows {
		return nil, &handlerError{err, "Error stair not found", http.StatusBadRequest}
		//log.Printf("No user with that ID")
	}

	if err != nil {
		return nil, &handlerError{err, "Internal Error when req DB", http.StatusInternalServerError}
		//panic(err)
	}

	stair := new(Stair)
	for row.Next() {
		var position, stairname, photo, description string
		var uid, id uint64

		fmt.Println(row)

		if err := row.Scan(&id, &position, &stairname, &description, &uid, &photo); err != nil {
			return nil, &handlerError{err, "Internal Error when reading req from DB", http.StatusInternalServerError}
			//log.Fatal(err)
		}

		stair.Id = id
		stair.Name = stairname
		stair.Photo = photo
		stair.User = uid
		stair.Description = description
		stair.Position = position

	}

	return stair, nil
}

/*
	Get all stairs from DB
	!READY FOR TESTING!

*/
func getAllStairs(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		return nil, &handlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer con.Close()

	rows, err := con.Query("select id, position, stairname, description, photo from Stairs")
	if err != nil {
		return nil, &handlerError{err, "Error in DB", http.StatusInternalServerError}
		//log.Printf("No user with that ID")
	}

	var result []Stair // create an array of stairs
	var id uint64
	var position, stairname, description, photo string

	for rows.Next() {
		stair := new(Stair)
		err = rows.Scan(&id, &position, &stairname, &description, &photo)
		if err != nil {
			return result, &handlerError{err, "Error in DB", http.StatusInternalServerError}
		}
		stair.Id = id
		stair.Position = position
		stair.Name = stairname
		stair.Photo = photo
		stair.Description = description
		result = append(result, *stair)
	}

	return result, nil
}

/*
	Get comment for a specific stairid
*/
func getComments(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	param := mux.Vars(req)["id"]
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		return nil, &handlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer con.Close()

	row, err := con.Query("select * from Comments where idStair =?", param)
	if err == sql.ErrNoRows {
		return nil, &handlerError{err, "Error commenting on Stair", http.StatusBadRequest}

	}

	if err != nil {
		return nil, &handlerError{err, "Internal Error when req DB", http.StatusInternalServerError}
	}
	var result []Comment
	var commentText, idToken string
	var commentDate time.Time
	var commentId, idStair uint64

	for row.Next() {
		comment := new(Comment)

		fmt.Println(row)

		if err := row.Scan(&commentId, &commentText, &commentDate, &idStair, &idToken); err != nil {
			return nil, &handlerError{err, "Internal Error when reading req from DB", http.StatusInternalServerError}
		}

		comment.CommentId = commentId
		comment.CommentText = commentText
		comment.CommentDate = commentDate
		comment.IdStair = idStair
		comment.IdToken = idToken

	}

	return result, nil

}

/*
	Add commment to db

*/
func addComment(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	data, e := ioutil.ReadAll(req.Body)

	fmt.Println(string(data))
	if e != nil {
		fmt.Println("AJAJAJ 1111")
		fmt.Println(string(data))
		return nil, &handlerError{e, "Can't read request", http.StatusBadRequest}
	}
	var payload Comment
	e = json.Unmarshal(data, &payload)
	payload.CommentDate = time.Now()

	if e != nil {
		fmt.Println("SATAN")
		fmt.Println(e)
		fmt.Println("kunde inte unmarshla detta:")
		fmt.Println(payload)
		return Comment{}, &handlerError{e, "Could'nt parse JSON", http.StatusInternalServerError}
	}
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		fmt.Println("Kunde inte öppna DB")
		return nil, &handlerError{err, "Internal server error", http.StatusInternalServerError}
	}
	defer con.Close()

	_, err = con.Exec("insert into Comments(commentText,commentDate,idStair,idToken) values(?,?,?,?)", payload.CommentText, payload.CommentDate, payload.IdStair, payload.IdToken)

	if err != nil {
		fmt.Println("Kunde inte lägga till :/")
		return nil, &handlerError{err, "Error adding to DB", http.StatusInternalServerError}
	}

	return payload, nil

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
	router.Handle("/users", handler(listAllUsers)).Methods("GET")
	// hämta ut infon för att lägga till ny
	router.Handle("/users", handler(addUser)).Methods("POST")
	router.Handle("/users/{id}", handler(getUser)).Methods("GET")
	router.Handle("/users/{id}", handler(removeUser)).Methods("DELETE")

	// Handlers for stairs
	router.Handle("/stair", handler(addStair)).Methods("POST")
	router.Handle("/stair/{id}", handler(getStair)).Methods("GET")
	router.Handle("/stairs", handler(getAllStairs)).Methods("GET")
	router.Handle("/stairs/{id}", handler(getUserStairs)).Methods("GET")

	// handlers for comments
	router.Handle("/comment", handler(addComment)).Methods("POST")
	router.Handle("/comment/{id}", handler(getComments)).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileHandler))
	http.Handle("/", router)

	log.Printf("Running on port %d\n", *port)

	addr := fmt.Sprintf("192.168.1.6:%d", *port)
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
