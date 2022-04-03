package echo_impl

import (
	"fmt"
	"log"
	"rest-bench/http_router"
	"strconv"

	"github.com/labstack/echo/v4"
)

type echoImpl struct {
	dispatch *echo.Echo
}

func New() http_router.Router {
	return &echoImpl{dispatch: echo.New()}
}

func (e *echoImpl) GET(uri string, resolver http_router.ResolveHandler) {
	e.dispatch.GET(uri, func(c echo.Context) error {
		ctx := newContext(c)
		return resolver(ctx)
	})
}

func (e *echoImpl) SERVE(port int) {
	portString := strconv.Itoa(port)
	fmt.Println("serving with ECHO on ", port)
	log.Fatal(e.dispatch.Start(":" + portString))
}
