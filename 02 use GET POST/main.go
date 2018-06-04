/*
Created on 2018/6/4 14:41

author: ChenJinLong

Content: 
*/
package main

import "github.com/gin-gonic/gin"

func main() {
	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	// 创建默认中间件的gin路由:
	// logger 和recovery (crash-free) 中间件
	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)

	// 如果没有事先定义服务运行端口,默认运行在:8080
	router.Run()
	// router.Run(":3000")自定义运行端口
}
