package gee

import (
	"net/http"
)

// 定义处理接口
type HandlerFunc func(c *Context)

// 定义路由映射引擎
type Engine struct {
	router *router
}

// 创建路由映射引擎对象
func New() *Engine {
	return &Engine{router: newRouter()}
}

// 注册GET方法路由
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.router.addRoute("GET", pattern, handler)
}

// 注册POST方法路由
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.router.addRoute("POST", pattern, handler)
}

// 处理Http请求转发
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 封装上下文
	c := NewContext(w, req)
	engine.router.handle(c)
}

// 启动服务
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
