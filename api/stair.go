package api

import (
	"database/sql"
	"encoding/json"
	//	"flag"
	"fmt"
	_ "github.com/ziutek/mymysql/godrv"
	"io/ioutil"
	"log"
	"net/http"
	//	"strconv"
	//	"time"

	"github.com/gorilla/mux"
)

//Stair struct
type Stair struct {
	Id          uint64 `json:"id"`
	Position    string `json:"position"`
	Name        string `json:"stairname"`
	User        uint64 `json:"user"`
	Photo       string `json:"photo"`
	Description string `json:"description"`
}

/*
	Add stair to DB
	!Done for testing!

*/
func AddStair(rw http.ResponseWriter, req *http.Request) (interface{}, *HandlerError) {
	data, e := ioutil.ReadAll(req.Body)

	fmt.Println("BEFORE MARSHAL " + string(data))
	if e != nil {
		fmt.Println("AJAJAJ 1111")
		fmt.Println(string(data))
		return nil, &HandlerError{e, "Can't read request", http.StatusBadRequest}
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
		return Stair{}, &HandlerError{e, "Could'nt parse JSON", http.StatusInternalServerError}
	}
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		fmt.Println("Kunde inte öppna DB")
		return nil, &HandlerError{err, "Internal server error", http.StatusInternalServerError}
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
		return nil, &HandlerError{err, "Error adding to DB", http.StatusInternalServerError}
	}

	return payload, nil
	//row, err := con.Query("select * from users where uid =?", param)
}

/*
	Get stair from DB
	!READY FOR TESTING!
*/

func GetUserStairs(rw http.ResponseWriter, req *http.Request) (interface{}, *HandlerError) {
	param := mux.Vars(req)["id"]
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		return nil, &HandlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer con.Close()

	row, err := con.Query("select * from Stairs where uid =?", param)
	if err == sql.ErrNoRows {
		return nil, &HandlerError{err, "Error no stairs found", http.StatusBadRequest}
		//log.Printf("No user with that ID")
	}

	if err != nil {
		return nil, &HandlerError{err, "Internal Error when req DB", http.StatusInternalServerError}
		//panic(err)
	}

	stair := new(Stair)
	for row.Next() {
		var position, stairname, photo, description string
		var uid, id uint64

		fmt.Println(row)

		if err := row.Scan(&id, &position, &stairname, &description, &uid, &photo); err != nil {
			return nil, &HandlerError{err, "Internal Error when reading req from DB", http.StatusInternalServerError}
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
func GetStair(rw http.ResponseWriter, req *http.Request) (interface{}, *HandlerError) {
	param := mux.Vars(req)["id"]
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		return nil, &HandlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer con.Close()

	row, err := con.Query("select * from Stairs where id =?", param)
	if err == sql.ErrNoRows {
		return nil, &HandlerError{err, "Error stair not found", http.StatusBadRequest}
		//log.Printf("No user with that ID")
	}

	if err != nil {
		return nil, &HandlerError{err, "Internal Error when req DB", http.StatusInternalServerError}
		//panic(err)
	}

	stair := new(Stair)
	for row.Next() {
		var position, stairname, photo, description string
		var uid, id uint64

		fmt.Println(row)

		if err := row.Scan(&id, &position, &stairname, &description, &uid, &photo); err != nil {
			return nil, &HandlerError{err, "Internal Error when reading req from DB", http.StatusInternalServerError}
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
func GetAllStairs(rw http.ResponseWriter, req *http.Request) (interface{}, *HandlerError) {
	con, err := sql.Open("mymysql", "tcp:localhost:3306*M7011E/root/jaam")
	if err != nil {
		return nil, &HandlerError{err, "Local error opening DB", http.StatusInternalServerError}
		log.Fatal(err)
	}
	defer con.Close()

	rows, err := con.Query("select id, position, stairname, description, photo from Stairs")
	if err != nil {
		return nil, &HandlerError{err, "Error in DB", http.StatusInternalServerError}
		//log.Printf("No user with that ID")
	}

	var result []Stair // create an array of stairs
	var id uint64
	var position, stairname, description, photo string

	for rows.Next() {
		stair := new(Stair)
		err = rows.Scan(&id, &position, &stairname, &description, &photo)
		if err != nil {
			return result, &HandlerError{err, "Error in DB", http.StatusInternalServerError}
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
