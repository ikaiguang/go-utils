# go-http

base on github.com/gin-gonic/gin

## config

see ./config.go

## dev environment

go version

> go version go1.12 darwin/amd64 & go module enable

## test

```bash

cd ./testdata

go run server.go
# INFO[0000] server addr :8999, isHTTPS : true

go run client.go
# {"message":"pong","url":"/v1/ping"} 


# ssl enable
curl -k https://127.0.0.1:8099/ping
# {"message":"pong","url":"/ping"}
curl -k https://127.0.0.1:8099/v1/ping
# {"message":"pong","url":"/v1/ping"}


# ssl unable
curl -k http://127.0.0.1:8099/ping
# {"message":"pong","url":"/ping"}
curl -k http://127.0.0.1:8099/v1/ping
# {"message":"pong","url":"/v1/ping"}

```
