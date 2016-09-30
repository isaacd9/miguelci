package main

import (
	"github.com/isaacd9/miguel/lib/database"
	"github.com/kataras/iris"
)

func hi(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, iris.Map{"isaac": "Diamond"})
}

func main() {
	app := iris.New()

	database.Connect()

	app.Get("/", hi)
	app.Listen(":8080")
}