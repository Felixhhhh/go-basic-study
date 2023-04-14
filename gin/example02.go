package main

//func main() {
//	// 1.创建路由
//	r := gin.Default()
//	// 2.绑定路由规则，执行的函数
//	// gin.Context，封装了request和response
//	//r.POST("/upload", func(c *gin.Context) {
//	//	file, err := c.FormFile("file")
//	//	if err != nil {
//	//		c.String(500, "上传图片出错")
//	//	}
//	//	c.SaveUploadedFile(file, file.Filename)
//	//	c.String(http.StatusOK, file.Filename)
//	//
//	//})
//
//	r.POST("/upload", func(c *gin.Context) {
//		_, headers, err := c.Request.FormFile("file")
//		if err != nil {
//			log.Printf("Error when try to get file: %v", err)
//		}
//		//headers.Size 获取文件大小
//		if headers.Size > 1024*1024*2 {
//			fmt.Println("文件太大了")
//			return
//		}
//		//headers.Header.Get("Content-Type")获取上传文件的类型
//		if headers.Header.Get("Content-Type") != "image/png" {
//			fmt.Println("只允许上传png图片")
//			return
//		}
//		c.SaveUploadedFile(headers, "./video/"+headers.Filename)
//		c.String(http.StatusOK, headers.Filename)
//	})
//
//	// 3.监听端口，默认在8080
//	// Run("里面不指定端口号默认为8080")
//	r.Run()
//}
