package http_router

import (
	"fmt"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

type echoImpl struct {
	dispatch *echo.Echo
}

func NewEcho() Router {
	return &echoImpl{dispatch: echo.New()}
}

func (e *echoImpl) GET(uri string, resolver ResolveHandler) {
	e.dispatch.GET(uri, func(c echo.Context) error {
		ctx := NewEchoContext(c)
		return resolver(ctx)
	})
}

func (e *echoImpl) SERVE(port int) {
	portString := strconv.Itoa(port)
	fmt.Println("serving with ECHO on ", port)
	log.Fatal(e.dispatch.Start(":" + portString))
}

func NewEchoContext(ctx echo.Context) ContextRouter {
	return &echoContext{
		ctx: ctx,
	}
}

type echoContext struct {
	ctx echo.Context
}

func (c *echoContext) JSON(code int, obj interface{}) error {
	return c.ctx.JSON(code, obj)
}
