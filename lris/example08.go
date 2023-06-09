package main

//
//// 2.
//// Create your own custom Context, put any fields you'll need.
//type MyContext struct {
//	// Embed the `iris.Context` -  嵌入 iris.Context
//	// It's totally optional but you will need this if you
//	// don't want to override all the context's methods!
//	// 这是可选的，但是你不想重写所有的 context的方法，需要这么做。
//	iris.Context
//}
//
//// Optionally: validate MyContext implements iris.Context on compile-time.
//// 可选的，在编译的时候验证是否实现 iris.Context
//var _ iris.Context = &MyContext{}
//
//// 3.
//func (ctx *MyContext) Do(handlers context.Handlers) {
//	context.Do(ctx, handlers)
//}
//
//// 3.
//func (ctx *MyContext) Next() {
//	context.Next(ctx)
//}
//
//// [Override any context's method you want here...]
//// Like the HTML below:
//
//func (ctx *MyContext) HTML(htmlContents string) (int, error) {
//	ctx.Application().Logger().Infof("Executing .HTML function from MyContext")
//
//	ctx.ContentType("text/html")
//	return ctx.WriteString(htmlContents)
//}
//
//func main() {
//	app := iris.New()
//
//	// 4.
//	app.ContextPool.Attach(func() iris.Context {
//		return &MyContext{
//			// If you use the embedded Context,
//			// call the `context.NewContext` to create one:
//			Context: context.NewContext(app),
//		}
//	})
//
//	// Register a view engine on .html files inside the ./view/** directory.
//	app.RegisterView(iris.HTML("./view", ".html"))
//
//	// Register your route, as you normally do
//	app.Handle("GET", "/", recordWhichContextForExample,
//		func(ctx iris.Context) {
//			// use the context's overridden HTML method.
//			ctx.HTML("<h1> Hello from my custom context's HTML! </h1>")
//		})
//
//	// This will be executed by the
//	// MyContext.Context embedded default context
//	// when MyContext is not directly define the View function by itself.
//	app.Handle("GET", "/hi/{firstname:alphabetical}", recordWhichContextForExample,
//		func(ctx iris.Context) {
//			firstname := ctx.Values().GetString("firstname")
//
//			ctx.ViewData("firstname", firstname)
//			ctx.Gzip(true)
//
//			ctx.View("hi.html")
//		})
//
//	app.Run(iris.Addr(":8080"))
//}
//
//// Should always print "($PATH) Handler is executing from 'MyContext'"
//func recordWhichContextForExample(ctx iris.Context) {
//	ctx.Application().Logger().Infof("(%s) Handler is executing from: '%s'",
//		ctx.Path(), reflect.TypeOf(ctx).Elem().Name())
//
//	ctx.Next()
//}
