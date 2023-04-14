package main

import "github.com/kataras/iris/v12"

func main() {
	//app := iris.Default()
	//
	//app.Post("/form_post", func(ctx iris.Context) {
	//	message := ctx.FormValue("message")
	//	nick := ctx.FormValueDefault("nick", "anonymous")
	//
	//	ctx.JSON(iris.Map{
	//		"status":  "posted",
	//		"message": message,
	//		"nick":    nick,
	//	})
	//})
	//
	//app.Run(iris.Addr(":8080"))

	app := iris.Default()

	app.Post("/post", func(ctx iris.Context) {
		id := ctx.URLParam("id")
		page := ctx.URLParamDefault("page", "0")
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")
		// or `ctx.PostValue` for POST, PUT & PATCH-only HTTP Methods.

		app.Logger().Infof("id: %s; page: %s; name: %s; message: %s",
			id, page, name, message)
	})

	app.Run(iris.Addr(":8080"))
}