/*
Created on 2018/6/4 15:27

author: ChenJinLong

Content: 
*/
package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	// 为多部分表单设置较低的内存限制 (默认32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// 上传文件到目标地址
		// c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8080")
}