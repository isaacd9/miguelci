package routes

import (
	"log"

	"github.com/isaacd9/miguel/lib/project"
	"github.com/isaacd9/miguel/model/error"
	"github.com/kataras/iris"
)

func ListBuilds(ctx *iris.Context) {
	id := ctx.Param("id")
	buildList, err := project.ListBuilds(id)

	if err != nil {
		log.Print("Error fetching builds: " + err.Error())
		ctx.JSON(401, errorModel.Error{
			Message: err.Error(),
		})
	}

	ctx.JSON(iris.StatusOK, buildList)
}

func StartBuild(ctx *iris.Context) {
	id := ctx.Param("id")
	err := project.StartBuild(id)

	if err != nil {
		log.Print("Error adding build: " + err.Error())
		ctx.JSON(500, errorModel.Error{
			Message: err.Error(),
		})
	}

	ctx.Data(iris.StatusCreated, nil)
}
