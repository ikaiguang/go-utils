package gohttp

import "github.com/gin-gonic/gin"

// ClineIp ip
func ClineIp(ctx *gin.Context) string {
	return ctx.ClientIP()
}

// Value value
func Value(ctx *gin.Context, key string) (interface{}, bool) {
	return ctx.Get(key)
}
