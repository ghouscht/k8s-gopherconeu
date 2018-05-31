package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ghouscht/k8s-gopherconeu/version"
	"github.com/gorilla/mux"
)

// BaseRouter returns a mux.Router with some basic routes setup.
func BaseRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", homeHandler()).Methods(http.MethodGet)
	r.HandleFunc("/release", releaseHandler()).Methods(http.MethodGet)

	return r
}

func homeHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request: %s", r.URL.Path)
		fmt.Fprint(w, "Hello :-)\n")
	}
}

func releaseHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "commit %q, time %q, release %q", version.Commit, version.BuildTime, version.Release)
	}
}
