package main

import (
	"net/http"
	router "rest-bench/http_router"
	chi "rest-bench/http_router/chi_impl"
	echo "rest-bench/http_router/echo_impl"
	fasthttp "rest-bench/http_router/fasthttp_routing_impl"
	fiber "rest-bench/http_router/fiber_impl"
	gin "rest-bench/http_router/gin_impl"
	mux "rest-bench/http_router/mux_impl"
	"rest-bench/model"
)

const portGin int = 8081
const portEcho int = 8082
const portMux int = 8083
const portFiber int = 8084
const portFastHttpRouting int = 8085
const portChi int = 8086

func main() {
	finished := make(chan bool)

	// ---------- INIT GIN ----------
	ginRouter := gin.New()
	ginRouter.GET("/", func(c router.ContextRouter) error {
		return c.JSON(http.StatusOK, model.Router{
			Name: "gin",
		})
	})
	go ginRouter.SERVE(portGin)

	// // ---------- INIT ECHO ----------
	echoRouter := echo.New()
	echoRouter.GET("/", func(c router.ContextRouter) error {
		return c.JSON(http.StatusOK, model.Router{
			Name: "echo",
		})
	})
	go echoRouter.SERVE(portEcho)

	// // ---------- INIT MUX ----------
	muxRouter := mux.New()
	muxRouter.GET("/", func(c router.ContextRouter) error {
		return c.JSON(http.StatusOK, model.Router{
			Name: "mux",
		})
	})
	go muxRouter.SERVE(portMux)

	// // ---------- INIT CHI ----------
	chiRouter := chi.New()
	chiRouter.GET("/", func(c router.ContextRouter) error {
		return c.JSON(http.StatusOK, model.Router{
			Name: "chi",
		})
	})
	go chiRouter.SERVE(portChi)

	// // ---------- INIT FIBER ----------
	fiberRouter := fiber.New()
	fiberRouter.GET("/", func(c router.ContextRouter) error {
		return c.JSON(http.StatusOK, model.Router{
			Name: "fiber",
		})
	})
	go fiberRouter.SERVE(portFiber)

	// // ---------- INIT FASTHTTP ROUTING ----------
	fasthttpRouter := fasthttp.New()
	fasthttpRouter.GET("/", func(c router.ContextRouter) error {
		return c.JSON(http.StatusOK, model.Router{
			Name: "fasthttp",
		})
	})
	go fasthttpRouter.SERVE(portFastHttpRouting)

	<-finished

}
