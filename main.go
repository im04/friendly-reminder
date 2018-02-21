package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"friendly-reminder/route"
)

func main() {
	defer getPanic()
	initIris();
}

func initIris() {
	app := iris.New()
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))
	app.Use(logger.New())
	route.Route(app)
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"errNo" : 404,
			"msg"   : "Not Found",
			"data"  : iris.Map{},
		})
	})
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"errNo" : 1,
			"msg"   : "error",
			"data"  : iris.Map{},
		})
	})
	addr := iris.Addr(":10002")
	app.Run(addr)
}

func getPanic() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}