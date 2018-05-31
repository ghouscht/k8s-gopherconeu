package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// BaseRouter returns a mux.Router with some basic routes setup.
func BaseRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", homeHandler()).Methods(http.MethodGet)

	return r
}

func homeHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request: %s", r.URL.Path)
		fmt.Fprint(w, "Hello :-)\n")
	}
}
