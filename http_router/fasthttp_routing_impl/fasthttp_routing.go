package fasthttp_routing

import (
	"fmt"
	"log"
	"rest-bench/http_router"
	"strconv"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

type fiberImpl struct {
	dispatch *routing.Router
}

func New() http_router.Router {
	return &fiberImpl{dispatch: routing.New()}
}

func (e *fiberImpl) GET(uri string, resolver http_router.ResolveHandler) {
	e.dispatch.Get(uri, func(c *routing.Context) error {
		ctx := newContext(c)
		return resolver(ctx)
	})
}

func (e *fiberImpl) SERVE(port int) {
	portString := strconv.Itoa(port)
	fmt.Println("serving with fiber on ", port)
	log.Fatal(fasthttp.ListenAndServe(":"+portString, e.dispatch.HandleRequest))
}
