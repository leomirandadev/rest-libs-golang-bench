package main

import (
	"rest-bench/http_router"
	"rest-bench/model"
)

func main() {
	finished := make(chan bool)

	ginRouter := http_router.NewMux()
	ginRouter.GET("/", func(c http_router.ContextRouter) error {
		return c.JSON(200, model.Router{
			Name: "mux",
		})
	})
	go ginRouter.SERVE(8081)

	echoRouter := http_router.NewEcho()
	echoRouter.GET("/", func(c http_router.ContextRouter) error {
		return c.JSON(200, model.Router{
			Name: "echo",
		})
	})
	go echoRouter.SERVE(8082)

	muxRouter := http_router.NewGin()
	muxRouter.GET("/", func(c http_router.ContextRouter) error {
		return c.JSON(200, model.Router{
			Name: "gin",
		})
	})
	go muxRouter.SERVE(8083)

	<-finished

}
