package gohttp

import "github.com/gin-gonic/gin"

// Route : http route
type Route struct {
	HttpMethod   string            // HTTP method
	RelativePath string            // relative path
	HandlerFunc  []gin.HandlerFunc // handler func
}

// NewRoute new route
func NewRoute(httpMethod, RouteRelativePath string, handler gin.HandlerFunc, middleware ...gin.HandlerFunc) *Route {
	return &Route{
		HttpMethod:   httpMethod,
		RelativePath: RouteRelativePath,
		HandlerFunc:  append(middleware, handler),
	}
}

// NewRouteGroup route group
func NewRouteGroup(engine *gin.Engine, groupRelativePath string) *gin.RouterGroup {
	return engine.Group(groupRelativePath)
}

// RegisterRoutes : register routes
func RegisterRoutes(engine *gin.Engine, routes []*Route) {
	for i := range routes {
		engine.Handle(routes[i].HttpMethod, routes[i].RelativePath, routes[i].HandlerFunc...)
	}
}

// RegisterGroupRoutes register routes
func RegisterGroupRoutes(group *gin.RouterGroup, routes []*Route) {
	for i := range routes {
		group.Handle(routes[i].HttpMethod, routes[i].RelativePath, routes[i].HandlerFunc...)
	}
}

// RegisterRoutesWithGroupPath : register routes
func RegisterRoutesWithGroupPath(engine *gin.Engine, groupRelativePath string, routes []*Route) {
	group := NewRouteGroup(engine, groupRelativePath)

	for i := range routes {
		group.Handle(routes[i].HttpMethod, routes[i].RelativePath, routes[i].HandlerFunc...)
	}
}
