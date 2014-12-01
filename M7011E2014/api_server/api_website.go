//website for api
package api

import (
	"net/http"
)

http.HandleFunc("/", servRest)
	http.ListenAndServe("localhost:1337", nil)