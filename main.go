package main

import (
	"github.com/isaacd9/miguel/lib/database"
	"github.com/isaacd9/miguel/routes"
	"github.com/kataras/iris"
)

func hi(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, iris.Map{"isaac": "Diamond"})
}

func main() {
	app := iris.New()
	database.Connect()

	app.Get("/", hi)
	app.Get("/projects", routes.ListProjects)
	app.Post("/projects", routes.AddProject)
	app.Listen(":8080")
}
