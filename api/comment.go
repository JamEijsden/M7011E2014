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
	"time"

	"github.com/gorilla/mux"
)

//Comment struct
type Comment struct {
	CommentId   uint64    `json:"commentId"`
	CommentText string    `json:"commentText"`
	CommentDate time.Time `json:"commentDate"`
	IdStair     uint64    `json:"idStair"`
	IdToken     string    `json:"idToken"`
}

/*
	Get comment for a specific stairid
*/
func GetComments(rw http.ResponseWriter, req *http.Request) (interface{}, *HandlerError) {
	param := mux.Vars(req)["id"]
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		return nil, &HandlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer con.Close()

	row, err := con.Query("select * from Comments where idStair =?", param)
	if err == sql.ErrNoRows {
		return nil, &HandlerError{err, "Error commenting on Stair", http.StatusBadRequest}

	}

	if err != nil {
		return nil, &HandlerError{err, "Internal Error when req DB", http.StatusInternalServerError}
	}

	var result []Comment
	var commentText, idToken string
	var commentDate time.Time
	var commentId, idStair uint64

	for row.Next() {
		fmt.Println(row)
		comment := new(Comment)

		fmt.Println(row)

		if err := row.Scan(&commentId, &commentText, &commentDate, &idStair, &idToken); err != nil {
			return nil, &HandlerError{err, "Internal Error when reading req from DB", http.StatusInternalServerError}
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
func AddComment(rw http.ResponseWriter, req *http.Request) (interface{}, *HandlerError) {
	data, e := ioutil.ReadAll(req.Body)

	fmt.Println(string(data))
	if e != nil {
		fmt.Println("AJAJAJ 1111")
		fmt.Println(string(data))
		return nil, &HandlerError{e, "Can't read request", http.StatusBadRequest}
	}
	var payload Comment
	e = json.Unmarshal(data, &payload)
	payload.CommentDate = time.Now()

	if e != nil {
		fmt.Println("SATAN")
		fmt.Println(e)
		fmt.Println("kunde inte unmarshla detta:")
		fmt.Println(payload)
		return Comment{}, &HandlerError{e, "Could'nt parse JSON", http.StatusInternalServerError}
	}
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		fmt.Println("Kunde inte öppna DB")
		return nil, &HandlerError{err, "Internal server error", http.StatusInternalServerError}
	}
	defer con.Close()

	_, err = con.Exec("insert into Comments(commentText,commentDate,idStair,idToken) values(?,?,?,?)", payload.CommentText, payload.CommentDate, payload.IdStair, payload.IdToken)

	if err != nil {
		fmt.Println("Kunde inte lägga till :/")
		return nil, &HandlerError{err, "Error adding to DB", http.StatusInternalServerError}
	}

	return payload, nil

}
