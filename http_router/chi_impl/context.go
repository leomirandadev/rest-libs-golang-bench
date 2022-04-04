package chi_impl

import (
	"encoding/json"
	"net/http"
	"rest-bench/http_router"
)

func newContext(w http.ResponseWriter, r *http.Request) http_router.ContextRouter {
	return &muxContext{w: w, r: r}
}

type muxContext struct {
	w http.ResponseWriter
	r *http.Request
}

func (c *muxContext) JSON(code int, obj interface{}) error {
	c.w.Header().Set("Content-Type", "application/json")
	c.w.WriteHeader(code)
	return json.NewEncoder(c.w).Encode(obj)
}
