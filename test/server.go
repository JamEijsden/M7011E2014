package main

import (
	//"fmt"
	"net/http"
	"os"
	"path/filepath"
)

var templatesPath = "templates"

// our init function to get our templates path depending on where we are.
func init() {
	dir, _ := os.Getwd() // gives us the source path if we haven't installed.
	templatesPath = filepath.Join(dir, templatesPath)
}

// calls our handlers and starts our server.
func main() {
	//	fmt.Println("Error starting server!")
	//	os.Exit(1)
	//}
	http.ListenAndServe(":8080", nil) // listen on all interfaces on port 8080
}
