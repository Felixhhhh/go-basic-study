package main

//
//func main() {
//	// 1.创建路由
//	r := gin.Default()
//	// 限制表单上传大小 8MB，默认为32MB
//	//r.MaxMultipartMemory = 8 << 20
//	r.POST("/upload", func(c *gin.Context) {
//		form, err := c.MultipartForm()
//		if err != nil {
//			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
//		}
//		// 获取所有图片
//		files := form.File["files"]
//		//遍历所有图片
//		for _, file := range files {
//			//逐个存储
//			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
//				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
//				return
//			}
//		}
//		c.String(http.StatusOK,fmt.Sprintf("upload ok %d files", len(files)))
//	})
//	r.Run()
//}
