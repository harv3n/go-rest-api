package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	muxDispatcher = mux.NewRouter()
)

type MuxRouter struct{}

func NewMuxRouter() Router {
	return &MuxRouter{}
}

func (*MuxRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*MuxRouter) Post(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*MuxRouter) Serve(port string) {
	fmt.Printf("Mux HTTP Server runnig on port %v", port)
	http.ListenAndServe(port, muxDispatcher)
}
