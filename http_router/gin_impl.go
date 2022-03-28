package http_router

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ginImpl struct {
	dispatch *gin.Engine
}

func NewGin() Router {
	gin.SetMode(gin.ReleaseMode)
	return &ginImpl{dispatch: gin.Default()}
}

func (g *ginImpl) GET(uri string, resolver ResolveHandler) {
	g.dispatch.GET(uri, func(c *gin.Context) {
		ctx := NewGinContext(c)
		resolver(ctx)
	})
}

func (g *ginImpl) SERVE(port int) {
	portString := strconv.Itoa(port)
	fmt.Println("serving with GIN on ", port)
	log.Fatal(g.dispatch.Run(":" + portString))
}

func NewGinContext(ctx *gin.Context) ContextRouter {
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
