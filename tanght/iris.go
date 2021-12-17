package main

import "github.com/kataras/iris/v12"

func iris_main() {
	app := iris.New()
	app.Get("/tanght/", iris_test01)
	app.Run(iris.Addr(":8000"))
}

func iris_test01(ctx iris.Context) {
	ctx.Write([]byte("hahaha, tanght"))
}
