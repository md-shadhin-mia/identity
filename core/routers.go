package core

import "github.com/gin-gonic/gin"

type Engine struct {
	*gin.Engine
}

func NewEngine() *Engine {
	return &Engine{Engine: gin.Default()}
}

type IRoutes interface {
	POST(relativePath string, handlers ...HandlerFunc) gin.IRoutes
	GET(relativePath string, handlers ...HandlerFunc) gin.IRoutes
	DELETE(relativePath string, handlers ...HandlerFunc) gin.IRoutes
	PATCH(relativePath string, handlers ...HandlerFunc) gin.IRoutes
	HEAD(relativePath string, handlers ...HandlerFunc) gin.IRoutes
}

func (e *Engine) POST(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	handlersGin := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handlersGin[i] = func(ctx *gin.Context) {
			handler(NewContext(ctx))
		}
	}
	return e.Engine.POST(relativePath, handlersGin...)
}

func (e *Engine) GET(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	handlersGin := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handlersGin[i] = func(ctx *gin.Context) {
			handler(NewContext(ctx))
		}
	}
	return e.Engine.GET(relativePath, handlersGin...)
}

func (e *Engine) DELETE(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	handlersGin := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handlersGin[i] = func(ctx *gin.Context) {
			handler(NewContext(ctx))
		}
	}
	return e.Engine.DELETE(relativePath, handlersGin...)
}

func (e *Engine) PATCH(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	handlersGin := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handlersGin[i] = func(ctx *gin.Context) {
			handler(NewContext(ctx))
		}
	}
	return e.Engine.PATCH(relativePath, handlersGin...)
}

func (e *Engine) PUT(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	handlersGin := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handlersGin[i] = func(ctx *gin.Context) {
			handler(NewContext(ctx))
		}
	}
	return e.Engine.PUT(relativePath, handlersGin...)
}

func (e *Engine) HEAD(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	handlersGin := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handlersGin[i] = func(ctx *gin.Context) {
			handler(NewContext(ctx))
		}
	}
	return e.Engine.HEAD(relativePath, handlersGin...)
}

func (e *Engine) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	handlersGin := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handlersGin[i] = func(ctx *gin.Context) {
			handler(NewContext(ctx))
		}
	}
	return &RouterGroup{e.Engine.Group(relativePath, handlersGin...)}
}

type RouterGroup struct {
	*gin.RouterGroup
}

func (r *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	handlersGin := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handlersGin[i] = func(ctx *gin.Context) {
			handler(NewContext(ctx))
		}
	}
	return &RouterGroup{r.RouterGroup.Group(relativePath, handlersGin...)}
}

func (e *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	handlersGin := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handlersGin[i] = func(ctx *gin.Context) {
			handler(NewContext(ctx))
		}
	}
	return e.RouterGroup.POST(relativePath, handlersGin...)
}

func (e *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	handlersGin := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handlersGin[i] = func(ctx *gin.Context) {
			handler(NewContext(ctx))
		}
	}
	return e.RouterGroup.GET(relativePath, handlersGin...)
}

func (e *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	handlersGin := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handlersGin[i] = func(ctx *gin.Context) {
			handler(NewContext(ctx))
		}
	}
	return e.RouterGroup.DELETE(relativePath, handlersGin...)
}

func (e *RouterGroup) PATCH(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	handlersGin := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handlersGin[i] = func(ctx *gin.Context) {
			handler(NewContext(ctx))
		}
	}
	return e.RouterGroup.PATCH(relativePath, handlersGin...)
}

func (e *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	handlersGin := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handlersGin[i] = func(ctx *gin.Context) {
			handler(NewContext(ctx))
		}
	}
	return e.RouterGroup.PUT(relativePath, handlersGin...)
}

func (e *RouterGroup) HEAD(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	handlersGin := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handlersGin[i] = func(ctx *gin.Context) {
			handler(NewContext(ctx))
		}
	}
	return e.RouterGroup.HEAD(relativePath, handlersGin...)
}
