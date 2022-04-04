package fiber_impl

import (
	"rest-bench/http_router"

	"github.com/gofiber/fiber/v2"
)

func newContext(ctx *fiber.Ctx) http_router.ContextRouter {
	return &echoContext{
		ctx: ctx,
	}
}

type echoContext struct {
	ctx *fiber.Ctx
}

func (c *echoContext) JSON(code int, obj interface{}) error {
	c.ctx.Status(code)
	return c.ctx.JSON(obj)
}
