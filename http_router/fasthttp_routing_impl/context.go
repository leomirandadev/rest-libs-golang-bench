package fasthttp_routing

import (
	"encoding/json"
	"rest-bench/http_router"

	routing "github.com/qiangxue/fasthttp-routing"
)

func newContext(ctx *routing.Context) http_router.ContextRouter {
	return &echoContext{
		ctx: ctx,
	}
}

type echoContext struct {
	ctx *routing.Context
}

func (c *echoContext) JSON(code int, obj interface{}) error {
	c.ctx.SetStatusCode(code)
	objByte, _ := json.Marshal(obj)

	_, err := c.ctx.Write(objByte)
	return err
}
