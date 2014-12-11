package main

import (
	"database/sql"
	//	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/ziutek/mymysql/godrv"
	"io/ioutil"
	"log"
	//"math"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"

	// this is test for photos
	"bytes"
	"encoding/base64"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"strings"
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

//Picture struct
type Picture struct {
	PhotoId uint64 `json:"photoId"`
	StairId uint64 `json:"idStair"`
	UserId  uint64 `json:"userID"`
	Picture string `json:"photo"`
	Preview string `json:"preview"`
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

	rows, err := con.Query("select name, uid from Users")
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
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	row, err := con.Query("select * from Users where idToken =?", param)
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

		if err := row.Scan(&uid, &name, &lastname, &idToken, &photo); err != nil {
			log.Fatal(err)
		}
		user.IdToken = idToken
		user.UserID = uid
		user.FirstName = name
		user.LastName = lastname
		user.Photo = photo
	}

	return user, nil
}

/*
	ADD USER TO DB
	!DONE for TESTING!

*/
func addUser(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {

	data, e := ioutil.ReadAll(r.Body)

	if e != nil {
		return nil, &handlerError{e, "Can't read request", http.StatusBadRequest}
	}

	// create new user called payload
	var payload User
	e = json.Unmarshal(data, &payload)

	if e != nil {
		return Stair{}, &handlerError{e, "Could'nt parse JSON", http.StatusInternalServerError}
	}
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		return nil, &handlerError{err, "Internal server error", http.StatusInternalServerError}
	}
	defer con.Close()
	row, _ := con.Query("select count(*) from Users where idToken=?", payload.IdToken)
	var count int
	for row.Next() {
		row.Scan(&count)
	}

	if count == 1 {
		return nil, &handlerError{nil, "User already exists", http.StatusFound}

	}

	_, err = con.Exec("insert into Users( name, lastname, idToken, photo) values(?,?,?,?)", payload.FirstName, payload.LastName, payload.IdToken, payload.Photo)

	if err != nil {
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
	id = id

	returnable := string("removeUser")
	return returnable, nil
}

/*
	Add stair to DB
	!Done for testing!

*/
func addStair(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	data, e := ioutil.ReadAll(req.Body)

	if e != nil {

		return nil, &handlerError{e, "Can't read request", http.StatusBadRequest}
	}
	var payload Stair
	e = json.Unmarshal(data, &payload)

	if e != nil {

		return Stair{}, &handlerError{e, "Could'nt parse JSON", http.StatusInternalServerError}
	}
	//handle photos
	//l, _ := base64.StdEncoding.DecodeString(payload.Photo)
	//reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(payload.Photo))
	a := strings.Split(payload.Photo, ",")
	reader, err := base64.StdEncoding.DecodeString(a[1])
	if err != nil {
		return err, &handlerError{err, "Internal", http.StatusInternalServerError}
	}
	s := string(reader[:])
	photo, _, err := image.Decode(strings.NewReader(s))
	if err != nil {

		return Stair{}, &handlerError{e, "Could'nt fix this image", http.StatusInternalServerError}
	}

	// resize photo
	newphoto := resize.Resize(215, 0, photo, resize.Lanczos3)

	//creates a buffer to save the encoded file to
	buf := new(bytes.Buffer)

	//encodes the image again and saves it to buf
	err = jpeg.Encode(buf, newphoto, nil)
	if err != nil {
		return Stair{}, &handlerError{e, "Could'nt fix this image", http.StatusInternalServerError}
	}

	//encodes the photo to base64 agian
	payload.Photo = base64.StdEncoding.EncodeToString(buf.Bytes())

	// adds the header from the website again
	payload.Photo = a[0] + "," + payload.Photo
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {

		return nil, &handlerError{err, "Internal server error", http.StatusInternalServerError}
	}
	defer con.Close()

	//inputs the stair to the db
	_, err = con.Exec("insert into Stairs(position, stairname, description, uid, photo) values(?,?,?,?,?)", payload.Position, payload.Name, payload.Description, payload.User, payload.Photo)

	if err != nil {

		return nil, &handlerError{err, "Error adding to DB", http.StatusInternalServerError}
	}

	return payload, nil
}

/*
	Add picture to db

*/
func addPicture(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	data, e := ioutil.ReadAll(req.Body)

	if e != nil {

		return nil, &handlerError{e, "Can't read request", http.StatusBadRequest}
	}

	// create new picture called payload
	var payload Picture
	e = json.Unmarshal(data, &payload)

	if e != nil {

		return Comment{}, &handlerError{e, "Could'nt parse JSON", http.StatusInternalServerError}
	}
	//Fixing preview
	a := strings.Split(payload.Picture, ",")
	reader, err := base64.StdEncoding.DecodeString(a[1])
	if err != nil {

		return err, &handlerError{err, "Internal", http.StatusInternalServerError}
	}
	s := string(reader[:])
	photo, _, err := image.Decode(strings.NewReader(s))
	if err != nil {

		return Stair{}, &handlerError{e, "Could'nt fix this image", http.StatusInternalServerError}
	}

	// resize photo
	newphoto := resize.Resize(215, 0, photo, resize.Lanczos3)

	//creates a buffer to save the encoded file to
	buf := new(bytes.Buffer)

	//encodes the image again and saves it to buf
	err = jpeg.Encode(buf, newphoto, nil)
	if err != nil {

		return Stair{}, &handlerError{e, "Could'nt fix this image", http.StatusInternalServerError}
	}

	//encodes the photo to base64 agian
	payload.Preview = base64.StdEncoding.EncodeToString(buf.Bytes())

	// adds the header from the website again
	payload.Preview = a[0] + "," + payload.Preview

	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {

		return nil, &handlerError{err, "Internal server error", http.StatusInternalServerError}
	}
	defer con.Close()

	_, err = con.Exec("insert into Photos(user_id,stair_id,photo_base64,preview) values(?,?,?,?)", payload.UserId, payload.StairId, payload.Picture, payload.Preview)

	if err != nil {

		return nil, &handlerError{err, "Error adding to DB", http.StatusInternalServerError}
	}

	returnvariables := new(Picture)
	returnvariables.PhotoId = payload.PhotoId
	returnvariables.Preview = payload.Preview
	return returnvariables, nil
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

	row, err := con.Query("select id, position, stairname, description, uid from Stairs where id =?", param)
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
		var position, stairname, description string
		var uid, id uint64

		if err := row.Scan(&id, &position, &stairname, &description, &uid); err != nil {
			return nil, &handlerError{err, "Internal Error when reading req from DB", http.StatusInternalServerError}
			//log.Fatal(err)
		}

		stair.Id = id
		stair.Name = stairname
		//stair.Photo = photo
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

	rows, err := con.Query("select id, position, stairname from Stairs")
	if err != nil {
		return nil, &handlerError{err, "Error in DB", http.StatusInternalServerError}
		//log.Printf("No user with that ID")
	}

	var result []Stair // create an array of stairs
	var id uint64
	var position, stairname string

	for rows.Next() {
		stair := new(Stair)
		err = rows.Scan(&id, &position, &stairname)
		if err != nil {
			return result, &handlerError{err, "Error in DB", http.StatusInternalServerError}
		}
		stair.Id = id
		stair.Position = position
		stair.Name = stairname
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

		if err := row.Scan(&commentId, &commentText, &commentDate, &idStair, &idToken); err != nil {
			return nil, &handlerError{err, "Internal Error when reading req from DB", http.StatusInternalServerError}
		}

		comment.CommentId = commentId
		comment.CommentText = commentText
		comment.CommentDate = commentDate
		comment.IdStair = idStair
		comment.IdToken = idToken
		result = append(result, *comment)

	}

	return result, nil

}

/*
	Add commment to db

*/
func addComment(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	data, e := ioutil.ReadAll(req.Body)

	if e != nil {

		return nil, &handlerError{e, "Can't read request", http.StatusBadRequest}
	}
	var payload Comment
	e = json.Unmarshal(data, &payload)
	payload.CommentDate = time.Now()

	if e != nil {

		return Comment{}, &handlerError{e, "Could'nt parse JSON", http.StatusInternalServerError}
	}
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {

		return nil, &handlerError{err, "Internal server error", http.StatusInternalServerError}
	}
	defer con.Close()

	_, err = con.Exec("insert into Comments(commentText,commentDate,idStair,idToken) values(?,?,?,?)", payload.CommentText, payload.CommentDate, payload.IdStair, payload.IdToken)

	if err != nil {

		return nil, &handlerError{err, "Error adding to DB", http.StatusInternalServerError}
	}

	return payload, nil

}

/*
	Get a specific picture from from db

*/
func getPicture(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	param := mux.Vars(req)["id"]
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	row, err := con.Query("select * from Photos where photo_id =?", param)
	if err == sql.ErrNoRows {
		log.Printf("No photo with that ID")
	}

	if err != nil {
		panic(err)
	}

	photo := new(Picture)
	for row.Next() {
		var photo_id, user_id, stair_id uint64
		var photo_base64 string

		if err := row.Scan(&photo_id, &user_id, &stair_id, &photo_base64); err != nil {
			log.Fatal(err)
		}
		photo.PhotoId = photo_id
		photo.UserId = user_id
		photo.StairId = stair_id
		photo.Picture = photo_base64
	}

	return photo, nil
}

/*
	Retrive a users pictures

*/
func retriveUserPictures(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	param := mux.Vars(req)["id"]
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		return nil, &handlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer con.Close()

	row, err := con.Query("select * from Photos where user_id =?", param)
	if err == sql.ErrNoRows {
		return nil, &handlerError{err, "Error commenting on Stair", http.StatusBadRequest}

	}

	if err != nil {
		return nil, &handlerError{err, "Internal Error when req DB", http.StatusInternalServerError}
	}
	var result []Picture
	var photo_id, user_id, stair_id uint64
	var photo_base64 string

	for row.Next() {
		picture := new(Picture)

		if err := row.Scan(&photo_id, &user_id, &stair_id, &photo_base64); err != nil {
			return nil, &handlerError{err, "Internal Error when reading req from DB", http.StatusInternalServerError}
		}

		picture.PhotoId = photo_id
		picture.UserId = user_id
		picture.StairId = stair_id
		picture.Picture = photo_base64
		result = append(result, *picture)

	}

	return result, nil

}

/*
	Retrive a stairs pictures

*/
func retriveStairPictures(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	param := mux.Vars(req)["id"]
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		return nil, &handlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer con.Close()

	row, err := con.Query("select * from Photos where stair_id =?", param)
	if err == sql.ErrNoRows {
		return nil, &handlerError{err, "Error commenting on Stair", http.StatusBadRequest}

	}

	if err != nil {
		return nil, &handlerError{err, "Internal Error when req DB", http.StatusInternalServerError}
	}
	var result []Picture
	var photo_id, user_id, stair_id uint64
	var photo_base64 string

	for row.Next() {
		picture := new(Picture)

		if err := row.Scan(&photo_id, &user_id, &stair_id, &photo_base64); err != nil {
			return nil, &handlerError{err, "Internal Error when reading req from DB", http.StatusInternalServerError}
		}

		picture.PhotoId = photo_id
		picture.UserId = user_id
		picture.StairId = stair_id
		picture.Picture = photo_base64
		result = append(result, *picture)

	}

	return result, nil
}

/*
	Retrive a stairs preview pictures from the db

*/
func retriveStairPreview(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	param := mux.Vars(req)["id"]
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		return nil, &handlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer con.Close()

	row, err := con.Query("select preview, photo_id from Photos where stair_id =?", param)
	if err == sql.ErrNoRows {
		return nil, &handlerError{err, "Error commenting on Stair", http.StatusBadRequest}

	}

	if err != nil {
		return nil, &handlerError{err, "Internal Error when req DB", http.StatusInternalServerError}
	}
	var result []Picture
	var photo_id uint64
	var preview string

	for row.Next() {
		picture := new(Picture)

		if err := row.Scan(&photo_id, &preview); err != nil {
			return nil, &handlerError{err, "Internal Error when reading req from DB", http.StatusInternalServerError}
		}

		picture.Preview = preview
		picture.PhotoId = photo_id
		result = append(result, *picture)

	}

	return result, nil
}

/*
	Retrive a users pictures previews from the db

*/
func retriveUserPicturesPreview(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	param := mux.Vars(req)["id"]
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		return nil, &handlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer con.Close()

	row, err := con.Query("select preview, photo_id from Photos where user_id =?", param)
	if err == sql.ErrNoRows {
		return nil, &handlerError{err, "Error Rertriving photos from user", http.StatusBadRequest}

	}

	if err != nil {
		return nil, &handlerError{err, "Internal Error when req DB", http.StatusInternalServerError}
	}
	var result []Picture
	var photo_id uint64
	var preview string

	for row.Next() {
		picture := new(Picture)

		if err := row.Scan(&photo_id, &preview); err != nil {
			return nil, &handlerError{err, "Internal Error when reading req from DB", http.StatusInternalServerError}
		}

		picture.Preview = preview
		picture.PhotoId = photo_id
		result = append(result, *picture)

	}

	return result, nil
}

/*
	Retrive a stairs photo

*/
func retriveStairPhoto(rw http.ResponseWriter, req *http.Request) (interface{}, *handlerError) {
	param := mux.Vars(req)["id"]
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		return nil, &handlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer con.Close()

	row, err := con.Query("select photo from Stairs where id =?", param)
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
		var photo string

		if err := row.Scan(&photo); err != nil {
			return nil, &handlerError{err, "Internal Error when reading req from DB", http.StatusInternalServerError}
			//log.Fatal(err)
		}
		stair.Photo = photo

	}

	return stair, nil
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
	// hämta alla bilder en användare har laddat upp
	router.Handle("/users/picture/{id}", handler(retriveUserPictures)).Methods("GET")
	router.Handle("/users/picture/preview/{id}", handler(retriveUserPicturesPreview)).Methods("GET")
	// Handlers for stairs
	router.Handle("/stair", handler(addStair)).Methods("POST")
	router.Handle("/stair/{id}", handler(getStair)).Methods("GET")
	router.Handle("/stairs", handler(getAllStairs)).Methods("GET")
	// Get all stairs a user have added..
	router.Handle("/stairs/{id}", handler(getUserStairs)).Methods("GET")
	//Get alla pictures for a stair
	router.Handle("/stair/picture/{id}", handler(retriveStairPictures)).Methods("GET")

	//Get all preview pictures for a stair
	router.Handle("/stair/picture/preview/{id}", handler(retriveStairPreview)).Methods("GET")

	router.Handle("/stair/photo/{id}", handler(retriveStairPhoto)).Methods("GET")

	// handlers for comments
	router.Handle("/comment", handler(addComment)).Methods("POST")
	router.Handle("/comment/{id}", handler(getComments)).Methods("GET")

	// Handlers for pictures
	router.Handle("/picture", handler(addPicture)).Methods("POST")
	router.Handle("/picture/{id}", handler(getPicture)).Methods("GET")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileHandler))
	http.Handle("/", router)

	log.Printf("Running on port %d\n", *port)

	addr := fmt.Sprintf("192.168.1.6:%d", *port)
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
