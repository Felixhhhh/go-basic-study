package main

//
//func main() {
//	// 1.创建路由
//	r := gin.Default()
//	//创建路由组
//	v1 :=r.Group("/v1")
//	// {} 是书写规范
//	{
//		v1.GET("/login", login)
//		v1.GET("/submit", submit)
//	}
//
//	v2 :=r.Group("/v2")
//	// {} 是书写规范
//	{
//		v2.GET("/login", login)
//		v2.GET("/submit", submit)
//	}
//	r.Run()
//}
//
//
//func login(c *gin.Context) {
//	name := c.DefaultQuery("name", "jack")
//	c.String(200, fmt.Sprintf("hello %s\n", name))
//}
//
//func submit(c *gin.Context) {
//	name := c.DefaultQuery("name", "lily")
//	c.String(200, fmt.Sprintf("hello %s\n", name))
//}
