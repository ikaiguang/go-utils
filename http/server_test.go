package gohttp

import (
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	os.Setenv(EnvKeyHttpServerAddr, ":8999")
	os.Setenv(EnvKeyHttpServerMode, "release")
	os.Setenv(EnvKeyHttpSSLEnable, "true")
	os.Setenv(EnvKeyHttpSSLServerName, "uufff.com")
	os.Setenv(EnvKeyHttpSSLServerKey, "testdata/cert/server.key")
	os.Setenv(EnvKeyHttpSSLServerCrt, "testdata/cert/server.crt")
	os.Setenv(EnvKeyHttpSSLClientCrt, "testdata/cert/client.crt")
	os.Setenv(EnvKeyHttpRegAddress, "127.0.0.1:8999")
	os.Setenv(EnvKeyHttpRegDomain, "http://uufff.com")
	os.Setenv(EnvKeyHttpOssDomain, "http://uufff.com/attachment")

	// new server
	engine := NewServer()

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
	RegisterRoutes(engine, []*Route{
		NewRoute("GET", "ping", handler),
	})

	// run server
	go func() {
		if err := RunServer(engine); err != nil {
			t.Error(err)
			t.FailNow()
		}
	}()

	// sleep
	//time.Sleep(time.Millisecond * 100)
	time.Sleep(time.Second)

	// req
	reqUrl := "https://127.0.0.1:8999/ping"
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := httpClient.Get(reqUrl)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer resp.Body.Close()

	// bad request
	if resp.StatusCode != http.StatusOK {
		t.Errorf("fail : %s", reqUrl)
		t.FailNow()
	}

	// read body
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("fail : %s", reqUrl)
		t.FailNow()
	}
	t.Logf("success : %s \n", b)
	fmt.Printf("success : %s \n", b)

	os.Exit(0)
}
