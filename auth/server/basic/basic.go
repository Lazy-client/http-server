package main

import (
	"fmt"
	"net/http"
)

var users = map[string]string{
	"admin": "admin",
}

func main() {
	http.HandleFunc("/", basicHandler)
	http.ListenAndServe("0.0.0.0:8081", nil)
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Header().Set("WWW-Authenticate", `Basic realm="example"`)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Authentication required")
		return
	}
	passwordHash, ok := users[username]
	if !ok {
		w.Header().Set("WWW-Authenticate", `Basic realm="example"`)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Invalid username")
		return
	}
	if password != passwordHash {
		w.Header().Set("WWW-Authenticate", `Basic realm="example"`)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Invalid password")
		return
	}
	fmt.Fprintln(w, "Hello, "+username+"!")
}
