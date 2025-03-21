package router

import "github.com/gin-gonic/gin"

type Router struct {
	engine *gin.Engine
}

type Handler struct {
	Path    string
	Handler gin.HandlerFunc
}

func New() *Router {
	engine := gin.Default()

	return &Router{
		engine: engine,
	}
}

func (r *Router) RegisterQueries(handlers ...Handler) {
	for _, handler := range handlers {
		r.engine.GET(handler.Path, handler.Handler)
	}
}

func (r *Router) RegisterCommands(handlers ...Handler) {
	for _, handler := range handlers {
		r.engine.POST(handler.Path, handler.Handler)
	}
}

func (r *Router) Run(addr ...string) error {
	return r.engine.Run(addr...)
}
