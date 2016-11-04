package main

import (
	"github.com/iris-contrib/middleware/logger"
	"github.com/isaacd9/miguel/lib/database"
	"github.com/isaacd9/miguel/routes"
	"github.com/kataras/iris"
)

func hi(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, iris.Map{"isaac": "Diamond"})
}

func main() {
	app := iris.New()
	app.Use(logger.New())
	database.Connect()

	app.Get("/", hi)
	app.Get("/v1/projects", routes.ListProjects)
	app.Post("/v1/projects", routes.AddProject)

	app.Delete("/v1/projects/:id", routes.RemoveProject)
	app.Get("/v1/projects/:id", routes.GetProjectInfo)

	app.Get("/v1/projects/:id/build", routes.ListBuilds)
	app.Post("/v1/projects/:id/build", routes.StartBuild)

	app.Listen(":8080")
}
