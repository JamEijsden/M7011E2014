package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// our user context input type
type UserContextInput struct {
	Tag              string
	Script           string
	Style            string
	AttributeLink    string
	AttributeDouble  string
	AttributeSingle  string
	AttributeOnEvent string
	AttributeSrc     string
}

var templates string // package scoped var to hold templates path

// parses the form values into our custom structure.
func getFormValues(userInput *url.Values) (contextInput *UserContextInput, err error) {
	// create a new UserContextInput struct to hold user input
	contextInput = new(UserContextInput)

	// iterate over the form values assigning each one to our struct.
	for key, value := range *userInput {
		switch key {
		case "tag":
			contextInput.Tag = value[0]
		case "script":
			contextInput.Script = value[0]
		case "style":
			contextInput.Style = value[0]
		case "attr_double":
			contextInput.AttributeDouble = value[0]
		case "attr_single":
			contextInput.AttributeSingle = value[0]
		case "attr_onevent":
			contextInput.AttributeOnEvent = value[0]
		case "attr_src":
			contextInput.AttributeSrc = value[0]
		default:
			return nil, fmt.Errorf("Unknown value passed in form!") // no funny business.
		}
	}
	return contextInput, nil // yes, it is safe to return a pointer in Go
}

// !!!　注意 DO NOT DO THIS (Disables XSS protection in IE/Chrome) 注意　!!!
func disableXSSProtection(function func(http.ResponseWriter,
	*http.Request)) func(http.ResponseWriter, *http.Request) {
	// return a closure wrapping our response with our header added
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "0")
		function(w, r)
	}
}

// our function to map handlers to URLs, also sets package var for our template path.
func ServerFunc(dir string) (err error) {
	templates = dir

	if _, err := os.Stat(templates); err != nil {
		return fmt.Errorf("templates path is incorrect!")
	}
	// map / to our rootHandler
	http.HandleFunc("/", disableXSSProtection(rootHandler))
	// map /display to our postHandler
	http.HandleFunc("/display", disableXSSProtection(postHandler))
	return nil
}

// handles everything in / except what is mapped in ServerFunc.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Parse our root.html template
	if t, err := template.ParseFiles(filepath.Join(templates, "root.html")); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// parse our form results page which injects user input.
	if t, err := template.ParseFiles(filepath.Join(templates, "form_results.html")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// make sure they parse correctly
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form values!", http.StatusInternalServerError)
			return
		}
		// get the form values
		userInput, err := getFormValues(&r.Form)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// inject user input with our implicit context aware encoding thanks to Go.
		t.Execute(w, userInput)
	}
}
