package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// 注册路由
func (r *router) addRouter(method string, pattern string, handlerFunc HandlerFunc) {
	log.Printf("Rount %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handlerFunc
}

// 处理http请求路由
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s \n", c.Path)
	}
}
