package gin_impl

import (
	"rest-bench/http_router"

	"github.com/gin-gonic/gin"
)

func newContext(ctx *gin.Context) http_router.ContextRouter {
	return &ginContext{
		ctx: ctx,
	}
}

type ginContext struct {
	ctx *gin.Context
}

func (c *ginContext) JSON(code int, obj interface{}) error {
	c.ctx.JSON(code, obj)
	return nil
}
