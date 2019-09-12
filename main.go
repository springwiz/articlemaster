package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// serves as the starting point of the application
// initializes the mux router and maps the routes to functions
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/{id}", GetArticleByIdHandler()).Methods("GET")
	r.HandleFunc("/articles/{id}", PutArticleHandler()).Methods("PUT")
	r.HandleFunc("/articles", PostArticleHandler()).Methods("POST")
	r.HandleFunc("/tags/{tagName}/{date}", GetArticlesByTagDateHandler()).Methods("GET")
	http.ListenAndServe(":80", r)
}
