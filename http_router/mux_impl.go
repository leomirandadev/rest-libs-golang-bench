package http_router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type muxImpl struct {
	dispatch *mux.Router
}

func NewMux() Router {
	return &muxImpl{dispatch: mux.NewRouter()}
}

func (m *muxImpl) GET(uri string, resolver ResolveHandler) {

	m.dispatch.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		ctx := NewMuxContext(w, r)
		resolver(ctx)
	}).Methods("GET")
}

func (m *muxImpl) SERVE(port int) {
	portString := strconv.Itoa(port)
	fmt.Println("serving with MUX on ", port)
	log.Fatal(http.ListenAndServe(":"+portString, m.dispatch))
}

func NewMuxContext(w http.ResponseWriter, r *http.Request) ContextRouter {
	return &muxContext{w: w, r: r}
}

type muxContext struct {
	w http.ResponseWriter
	r *http.Request
}

func (c *muxContext) JSON(code int, obj interface{}) error {
	c.w.Header().Set("Content-Type", "application/json")
	c.w.WriteHeader(code)
	return json.NewEncoder(c.w).Encode(obj)
}
