package main

import (
	"fmt"
	"github.com/kataras/iris"
)

func main() {
	app := newApp()

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

}

func newApp() *iris.Application {
	app := iris.New()

	//	app.Get("/",func(ctx iris.Context) {
	//		ctx.Text("Hello World")
	//	})

	app.Get("{name:string}", func(ctx iris.Context) {
		ctx.Text(fmt.Sprintf("Hello %s", ctx.Params().Get("name")))
	})

	return app
}
