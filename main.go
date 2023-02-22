package main

import (
	"douyin/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	dao.Init()

	r := gin.Default()

	registerRoute(r)

	r.Run(":8082")

}
