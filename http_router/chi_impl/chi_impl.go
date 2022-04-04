package chi_impl

import (
	"fmt"
	"log"
	"net/http"
	"rest-bench/http_router"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type chiImpl struct {
	dispatch *chi.Mux
}

func New() http_router.Router {
	return &chiImpl{dispatch: chi.NewRouter()}
}

func (m *chiImpl) GET(uri string, resolver http_router.ResolveHandler) {

	m.dispatch.Get(uri, func(w http.ResponseWriter, r *http.Request) {
		ctx := newContext(w, r)
		resolver(ctx)
	})
}

func (m *chiImpl) SERVE(port int) {
	portString := strconv.Itoa(port)
	fmt.Println("serving with CHI on ", port)
	log.Fatal(http.ListenAndServe(":"+portString, m.dispatch))
}
