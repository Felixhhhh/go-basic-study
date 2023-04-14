package main

//
//func main(){
//	app := iris.New()
//	app.OnErrorCode(iris.StatusNotFound, notFound)
//	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)
//	// to register a handler for all "error"
//	// status codes(kataras/iris/context.StatusCodeNotSuccessful)
//	// defaults to < 200 || >= 400:
//	// app.OnAnyErrorCode(handler)
//	app.RegisterView(iris.HTML("./views", ".html"))
//	app.Get("/", index)
//	app.Run(iris.Addr(":8080"))
//}
//
//func notFound(ctx iris.Context) {
//	// when 404 then render the template
//	// $views_dir/errors/404.html
//	ctx.View("errors/404.html")
//}
//
//func internalServerError(ctx iris.Context) {
//	ctx.WriteString("Oups something went wrong, try again")
//}
//
//func index(ctx iris.Context) {
//	ctx.View("index.html")
//}