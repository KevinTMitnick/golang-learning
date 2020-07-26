package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main01() {
	// 1. 创建路由
	r := gin.Default()

	// 2. 绑定路由
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "<h1>Hello World</h1>")
	//})

	r.GET("/welcome", func(c *gin.Context) {

		//name := c.Param("name")

		name := c.DefaultQuery("name", "KT")

		c.String(http.StatusOK, fmt.Sprintf("Hello %s", name))
	})


	// 3. 监听端口，默认在8080
	r.Run(":8080")

}

// login.html
func main02() {
	r := gin.Default()

	r.POST("form", func(c *gin.Context) {

		type1 := c.DefaultPostForm("type", "alert")
		username := c.PostForm("username")
		password := c.PostForm("password")

		// 多选框
		hobbys := c.PostFormArray("Hobby")
		c.String(http.StatusOK,
			fmt.Sprintf("Type is %s, username is %s, password is %s, Hobbys is %v",
				type1, username,password,hobbys))
	})

	r.Run(":8080")
}

// 接收上传文件： formRecever.html
func main() {

	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		// 表单取文件
		file, _ := c.FormFile("file")
		log.Print(file.Filename)

		// 传到项目的根目录
		c.SaveUploadedFile(file, file.Filename)

		c.String(200, fmt.Sprintf("%s upload!", file.Filename))
	})

	r.Run(":8080")
}
