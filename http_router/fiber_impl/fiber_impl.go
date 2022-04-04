package fiber_impl

import (
	"fmt"
	"log"
	"rest-bench/http_router"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type fiberImpl struct {
	dispatch *fiber.App
}

func New() http_router.Router {
	return &fiberImpl{dispatch: fiber.New()}
}

func (e *fiberImpl) GET(uri string, resolver http_router.ResolveHandler) {
	e.dispatch.Get(uri, func(c *fiber.Ctx) error {
		ctx := newContext(c)
		return resolver(ctx)
	})
}

func (e *fiberImpl) SERVE(port int) {
	portString := strconv.Itoa(port)
	fmt.Println("serving with FIBER on ", port)
	log.Fatal(e.dispatch.Listen(":" + portString))
}
