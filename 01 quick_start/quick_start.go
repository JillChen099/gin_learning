/*
Created on 2018/6/4 13:40

author: ChenJinLong

Content: 
*/
package main

import "github.com/gin-gonic/gin"

func main() {
	gin.DisableConsoleColor()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
		"message": "pong",
			})
		})
	r.Run() // 监听和启动服务 0.0.0.0:8080
}