package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
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
	muxDispatcher.HandleFunc(url, handler).Methods(http.MethodGet)
}
func (*muxRouter) POST(url string, handler func(http.ResponseWriter, *http.Request)) {
	muxDispatcher.HandleFunc(url, handler).Methods(http.MethodPost)
}
func (*muxRouter) PATCH(url string, handler func(http.ResponseWriter, *http.Request)) {
	muxDispatcher.HandleFunc(url, handler).Methods(http.MethodPatch)
}
func (*muxRouter) PUT(url string, handler func(http.ResponseWriter, *http.Request)) {
	muxDispatcher.HandleFunc(url, handler).Methods(http.MethodPut)
}
func (*muxRouter) DELETE(url string, handler func(http.ResponseWriter, *http.Request)) {
	muxDispatcher.HandleFunc(url, handler).Methods(http.MethodDelete)
}
func (*muxRouter) SERVE(port string) {
	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	swaggerHandler := middleware.Redoc(opts, nil)
	muxDispatcher.Handle("/docs", swaggerHandler)
	muxDispatcher.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	log.Println(fmt.Sprintf("Starting server on port %s", port))
	log.Fatalln(http.ListenAndServe(port, muxDispatcher))
}
