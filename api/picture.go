package api

import (
	"database/sql"
	//	"encoding/base64"
	"encoding/json"
	//	"flag"
	"fmt"
	_ "github.com/ziutek/mymysql/godrv"
	"io/ioutil"
	"log"
	//"math"
	"net/http"
	//	"strconv"
	//"time"

	"github.com/gorilla/mux"
)

//Picture struct
type Picture struct {
	PhotoId uint64 `json:"photoId"`
	StairId uint64 `json:"idStair"`
	UserId  uint64 `json:"userID"`
	Picture string `json:"photo"`
}

/*
	Get a specific picture from from db

*/
func GetPicture(rw http.ResponseWriter, req *http.Request) (interface{}, *HandlerError) {
	param := mux.Vars(req)["id"]
	fmt.Println(param)
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
	Add picture to db

*/
func AddPicture(rw http.ResponseWriter, req *http.Request) (interface{}, *HandlerError) {
	data, e := ioutil.ReadAll(req.Body)

	if e != nil {
		fmt.Println("AJAJAJ 1111")
		fmt.Println(string(data))
		return nil, &HandlerError{e, "Can't read request", http.StatusBadRequest}
	}

	// create new picture called payload
	var payload Picture
	e = json.Unmarshal(data, &payload)

	if e != nil {
		fmt.Println(e)
		fmt.Println(payload)
		return Comment{}, &HandlerError{e, "Could'nt parse JSON", http.StatusInternalServerError}
	}
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		fmt.Println("Kunde inte öppna DB")
		return nil, &HandlerError{err, "Internal server error", http.StatusInternalServerError}
	}
	defer con.Close()

	_, err = con.Exec("insert into Photos(user_id,stair_id,photo_base64) values(?,?,?,?)", payload.UserId, payload.StairId, payload.Picture)

	if err != nil {
		fmt.Println("Kunde inte lägga till :/")
		return nil, &HandlerError{err, "Error adding to DB", http.StatusInternalServerError}
	}

	return payload, nil
}

/*
	Retrive a users pictures

*/
func RetriveUserPictures(rw http.ResponseWriter, req *http.Request) (interface{}, *HandlerError) {
	param := mux.Vars(req)["id"]
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		return nil, &HandlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer con.Close()

	row, err := con.Query("select * from Photos where user_id =?", param)
	if err == sql.ErrNoRows {
		return nil, &HandlerError{err, "Error commenting on Stair", http.StatusBadRequest}

	}

	if err != nil {
		return nil, &HandlerError{err, "Internal Error when req DB", http.StatusInternalServerError}
	}
	var result []Picture
	var photo_id, user_id, stair_id uint64
	var photo_base64 string

	for row.Next() {
		picture := new(Picture)

		fmt.Println(row)

		if err := row.Scan(&photo_id, &user_id, &stair_id, &photo_base64); err != nil {
			return nil, &HandlerError{err, "Internal Error when reading req from DB", http.StatusInternalServerError}
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
func RetriveStairPictures(rw http.ResponseWriter, req *http.Request) (interface{}, *HandlerError) {
	param := mux.Vars(req)["id"]
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		return nil, &HandlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer con.Close()

	row, err := con.Query("select * from Photos where stair_id =?", param)
	if err == sql.ErrNoRows {
		return nil, &HandlerError{err, "Error commenting on Stair", http.StatusBadRequest}

	}

	if err != nil {
		return nil, &HandlerError{err, "Internal Error when req DB", http.StatusInternalServerError}
	}
	var result []Picture
	var photo_id, user_id, stair_id uint64
	var photo_base64 string

	for row.Next() {
		picture := new(Picture)

		fmt.Println(row)

		if err := row.Scan(&photo_id, &user_id, &stair_id, &photo_base64); err != nil {
			return nil, &HandlerError{err, "Internal Error when reading req from DB", http.StatusInternalServerError}
		}

		picture.PhotoId = photo_id
		picture.UserId = user_id
		picture.StairId = stair_id
		picture.Picture = photo_base64
		result = append(result, *picture)

	}

	return result, nil
}
