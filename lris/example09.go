package main

//
//func main() {
//	app := iris.New()
//
//	// middleware for all versions.
//	//myMiddleware := func(ctx iris.Context) {
//	//	// [...]
//	//	ctx.Next()
//	//}
//	//
//	//myCustomNotVersionFound := func(ctx iris.Context) {
//	//	ctx.StatusCode(404)
//	//	ctx.Writef("%s version not found", versioning.GetVersion(ctx))
//	//}
//
//	userAPI := app.Party("/api/user")
//	userAPI.Get("/", func(ctx iris.Context) {
//		ctx.WriteString("test")
//	})
//	//userAPI.Get("/", myMiddleware, versioning.NewMatcher(versioning.Map{
//	//	"1.0":               sendHandler(v10Response),
//	//	">= 2, < 3":         sendHandler(v2Response),
//	//	versioning.NotFound: myCustomNotVersionFound,
//	//}))
//	app.Run(iris.Addr(":8080"))
//}
