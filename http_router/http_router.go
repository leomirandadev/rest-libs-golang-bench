package http_router

type ResolveHandler = func(ContextRouter) error

type Router interface {
	GET(uri string, resolver ResolveHandler)

	SERVE(port int)
}

type ContextRouter interface {
	JSON(code int, obj interface{}) error
}
