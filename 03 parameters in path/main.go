/*
Created on 2018/6/4 14:42

author: ChenJinLong

Content: 
*/
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	)

func main() {
	router := gin.Default()

	// 这个 handler 将 匹配 /user/john 但不会匹配 /user/ 或者 /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 然而, 这个将匹配 /user/john/ 和 /user/john/send
	// 如果没有其他 routers 匹配 /user/john, 它将重定向/user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run(":8080")
}
