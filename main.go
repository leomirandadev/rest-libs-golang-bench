package main

import (
	"net/http"
	router "rest-bench/http_router"
	echo "rest-bench/http_router/echo_impl"
	gin "rest-bench/http_router/gin_impl"
	mux "rest-bench/http_router/mux_impl"
	"rest-bench/model"
)

const portGin int = 8081
const portEcho int = 8082
const portMux int = 8083

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

	<-finished

}
