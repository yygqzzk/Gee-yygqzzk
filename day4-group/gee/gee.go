package gee

import (
	"log"
	"net/http"
)

type RouterGroup struct {
	prefix      string        // 路由组的前缀
	middlewares []HandlerFunc // 中间件列表
	engine      *Engine       // 关联的引擎实例
}

// 定义处理函数接口
type HandlerFunc func(c *Context)

// 定义路由映射引擎
type Engine struct {
	*RouterGroup                // 嵌入 RouterGroup，使 Engine 具有路由组的功能
	router       *router        // 路由表
	groups       []*RouterGroup // 所有路由组
}

// 创建路由映射引擎对象
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix, // 组合前缀
		engine: engine,

		// 关联引擎
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp // 组合完整路由路径
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// 注册GET方法路由
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// 注册POST方法路由
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
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
