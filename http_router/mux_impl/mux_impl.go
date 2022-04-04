package mux_impl

import (
	"fmt"
	"log"
	"net/http"
	"rest-bench/http_router"
	"strconv"

	"github.com/gorilla/mux"
)

type muxImpl struct {
	dispatch *mux.Router
}

func New() http_router.Router {
	return &muxImpl{dispatch: mux.NewRouter()}
}

func (m *muxImpl) GET(uri string, resolver http_router.ResolveHandler) {

	m.dispatch.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		ctx := newContext(w, r)
		_ = resolver(ctx)
	}).Methods("GET")
}

func (m *muxImpl) SERVE(port int) {
	portString := strconv.Itoa(port)
	fmt.Println("serving with MUX on ", port)
	log.Fatal(http.ListenAndServe(":"+portString, m.dispatch))
}
