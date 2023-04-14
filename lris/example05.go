package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	// register our routes.
	app.Get("/", indexHandler)
	app.Get("/contact", contactHandler)

	// Order of those calls does not matter,
	// `UseGlobal` and `DoneGlobal` are applied to existing routes
	// and future routes also.
	//
	// Remember: the `Use` and `Done` are applied to the current party's and its children,
	// so if we used the `app.Use/Done before the routes registration
	// it would work like UseGlobal/DoneGlobal in this case,
	// because the `app` is the root "Party".
	app.UseGlobal(before)
	app.DoneGlobal(after)

	app.Run(iris.Addr(":8080"))
}

func before(ctx iris.Context) {
	println("before")
	ctx.Next()
	// [...]
}

func after(ctx iris.Context) {
	println("after")
	ctx.Next()
	// [...]
}

func indexHandler(ctx iris.Context) {
	// write something to the client as a response.
	ctx.HTML("<h1>Index</h1>")

	ctx.Next() // execute the "after" handler registered via `Done`.
}

func contactHandler(ctx iris.Context) {
	// write something to the client as a response.
	ctx.HTML("<h1>Contact</h1>")

	ctx.Next() // execute the "after" handler registered via `Done`.
}