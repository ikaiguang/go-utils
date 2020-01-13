# go-utils

golang utils

## 介绍

一些简单的封装，方便调用

## config

see ./xxx/config.go

## dev environment

go version : 1.12.0 +

## test

go test -v .
go test -v ./xxx/

## todo

待做事项，做完打勾（中括号加x）

例如 ： [x] 用户注册 2019-08-31

- [x] 数据库 初始化
- [x] grpc 初始化
- [x] gin 初始化
- [x] etcd 初始化
- [x] redis 初始化
- [x] session 初始化
- [x] gin & grpc 中间件 : jwt
- [x] 日志
- [ ] 使用 json & yaml 配置 全局配置
- [ ] grpc 的 balancer, 参考 etcd 
- [ ] grpc 客户端 连接池, 参考 db.pool redis.pool