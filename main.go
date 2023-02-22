package main

import (
	"douyin/dao"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	dao.Init()

	// use hertz to build http
	h := server.Default()

	registerRoute(h)

	h.Spin()

}
