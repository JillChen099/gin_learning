/*
Created on 2018/6/4 15:28

author: ChenJinLong

Content: 
*/
package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"log"
)

func main() {
	router := gin.Default()
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// 多文件表单
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// 上传文件到目标地址
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.Run(":8080")
}