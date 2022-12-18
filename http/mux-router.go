package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

func NewMuxRouter() Router {
	return &muxRouter{}
}

var (
	muxDispatcher *mux.Router = mux.NewRouter().StrictSlash(true)
)

func (*muxRouter) GET(url string, handler func(http.ResponseWriter, *http.Request)) {
	muxDispatcher.HandleFunc(url, handler).Methods("GET")
}
func (*muxRouter) POST(url string, handler func(http.ResponseWriter, *http.Request)) {
	muxDispatcher.HandleFunc(url, handler).Methods("POST")
}
func (*muxRouter) PATCH(url string, handler func(http.ResponseWriter, *http.Request)) {
	muxDispatcher.HandleFunc(url, handler).Methods("PATCH")
}
func (*muxRouter) PUT(url string, handler func(http.ResponseWriter, *http.Request)) {
	muxDispatcher.HandleFunc(url, handler).Methods("PUT")
}
func (*muxRouter) DELETE(url string, handler func(http.ResponseWriter, *http.Request)) {
	muxDispatcher.HandleFunc(url, handler).Methods("DELETE")
}
func (*muxRouter) SERVE(port string) {
	log.Println(fmt.Sprintf("Starting server on port %s", port))
	log.Fatalln(http.ListenAndServe(port, muxDispatcher))
}
