package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/ziutek/mymysql/godrv"
	//"github.com/ziutek/mymysql/mysql"
	//_ "github.com/go-sql-driver/mysql"
)

//username := "root"
//password := "M7017E"
//databaseName:= "m7017e"

//func connect() {
//	db, err := sql.Open("mysql",
//		"user:password@tcp(127.0.0.1:3306)/users")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()

//}

type user struct {
	userID, firstName, lastname string
}

func connect() {
	username := "root"
	password := "M7017E"
	db := mysql.New("tcp", "127.0.0.1:3306", user, password)

	err := db.Connect()
	if err != nil {
		panic(err)
	}
}

// Inputs user into DB
func input() {

}

/*func checkuser() {

	var (
	id int
	name string
)
	rows, err := db.Query("select id, name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
	err := rows.Scan(&id, &name)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(id, name)
}
err = rows.Err()
if err != nil {
	log.Fatal(err)
}
}*/

func checkUser(checksum string) {
	username := "root"
	password := "M7017E"
	databaseName := "m7017e"
	db := mysql.New("tcp", "127.0.0.1:3306", user, password, databaseName)

	err := db.Connect()
	if err != nil {
		panic(err)
	}
	row, res, err := db.Query("SELECT * FROM users WHERE userID = %d", checksum)
	if err != nil {
		panic(err)
	}
	if row == 0 {

		/*										#
							??
			what to return if user not in db???

		#										*/

		return
	} else {
		sendback := encodeJson(row)
		return sendback
	}

}

func sendline(checksum string) {
	username := "root"
	password := "M7017E"
	databaseName := "m7017e"
	con, err := sql.Open("mymusql", databaseName+"/"+username+"/"+password)
	defer con.Close()
	row := con.QueryRow("select * from users where userID=?", checksum)
	answer := encodeJson(row)
	return answer
	//cb := new(SomeThing)
	//err := row.Scan(&cb.Mdpr, &cb.X, &cb.Y, &cb.Z)
	//db := mysql.New("tcp", "127.0.0.1:3306", user, password, databaseName)

	//err := db.Connect()
	//if err != nil {
	//	panic(err)
	//}
	//row, res, err := db.Query("SELECT * FROM users WHERE userID = %d", checksum)
	//if err != nil {
	//	panic(err)
	//}
	//answer := encodeJson(row)
	//return answer

}

func sendmultipleline(trappa int) {
	username := "root"
	password := "M7017E"
	databaseName := "m7017e"
	con, err := sql.Open("mymusql", databaseName+"/"+username+"/"+password)
	defer con.Close()
	rows, err := con.Query("select * where trappa=?", trappa)
	if err != nil {
		fmt.Println("Error in Sendmultipleline func: ", err)
		return nil
	} else {
		answer := encodeJson(rows)
		return answer
	}
	//db := mysql.New("tcp", "127.0.0.1:3306", user, password, databaseName)

	//err := db.Connect()
	//if err != nil {
	//		panic(err)
	//	}
	//	row, res, err := db.Query("SELECT * FROM users WHERE userID = %d", checksum)
	//	if err != nil {
	//		panic(err)
	//	}
	//	answer := encodeJson(row)
	//	return answer

}

func encodeJson(msg string) []byte {
	//encode to json
	jsonenconded, err := json.Marshal(msg)
	if err == nil {
		return jsonenconded
	}
	fmt.Println("Error in encodeJson func: ", err)
	return nil

}
