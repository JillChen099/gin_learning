/*
Created on 2018/6/7 10:19

author: ChenJinLong

Content: 分组路由
*/
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/login",loginEndPoint)
		v1.POST("/submit",submitEndPoint)
		v1.POST("/read",readEndpoint)

	}
	v2 := router.Group("/v2")
	{
		v2.POST("/login",loginEndPoint)
		v2.POST("/submit",submitEndPoint)
		v2.POST("/read",readEndpoint)
	}
	router.Run(":8080")



}

