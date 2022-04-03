package echo_impl

import (
	"rest-bench/http_router"

	"github.com/labstack/echo/v4"
)

func newContext(ctx echo.Context) http_router.ContextRouter {
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
