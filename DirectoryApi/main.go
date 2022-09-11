package main

import (
	. "DirectoryApi/handler"

	"net/http"
)

func main() {

	http.HandleFunc("/users", Handler)

	http.HandleFunc("/users/signup", SignupHandler)

	http.ListenAndServe(":9090", nil)

}
