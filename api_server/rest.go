package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct{
	Stuff Data
}

type Data struct{
	Fruit Fruits
	Veggies Vegetables
}



type fruits map[string]int 
type Vegetables map[string]int
func servRest(w http.ResponseWriter, r *http.Request) {
	responce := getJsonResponse()
}

func main() {
	http.HandleFunc("/", servRest)
	http.ListenAndServe("localhost:1337", nil)
}


func getJsonResponse() {
	fruits := make(map[string]int)
	fruits["Apples"] = 25
	fruits["Oranges"] = 11
	vegetables = make([string]int)
	vegetables]["Carrots"]

}