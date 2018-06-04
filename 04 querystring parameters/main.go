/*
Created on 2018/6/4 14:54

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

	// 查询字符串通过现有的底层请求对象进行解析
	// 请求根据url匹配进行响应:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	router.Run(":8080")
}
