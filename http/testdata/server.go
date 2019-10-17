package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ikaiguang/go-utils/http"
	"os"
)

func main() {
	os.Setenv("AppHttpServerAddress", ":8999")
	os.Setenv("AppHttpServerMode", "release")
	os.Setenv("AppHttpSSLEnable", "true")
	os.Setenv("AppHttpSSLServerName", "uufff.com")
	os.Setenv("AppHttpSSLServerKey", "cert/server.key")
	os.Setenv("AppHttpSSLServerCrt", "cert/server.crt")
	os.Setenv("AppHttpSSLClientCrt", "cert/client.crt")

	// new server
	engine := gohttp.NewServer()

	// hand
	var handler = func(c *gin.Context) {
		c.JSON(200, gin.H{
			"url":     c.Request.URL.Path,
			"message": "pong",
		})
	}

	// route
	// curl -k https://127.0.0.1:8999/ping
	// curl -k http://127.0.0.1:8999/ping
	gohttp.RegisterRoutes(engine, []*gohttp.Route{
		gohttp.NewRoute("GET", "ping", handler),
	})

	// route group
	// curl -k https://127.0.0.1:8999/v1/ping
	// curl -k http://127.0.0.1:8999/v1/ping
	gohttp.RegisterRoutesWithGroupPath(engine, "v1", []*gohttp.Route{
		gohttp.NewRoute("GET", "ping", handler),
	})

	// start
	if err := gohttp.RunServer(engine); err != nil {
		panic(err)
	}
}
