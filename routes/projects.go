package routes

import (
	"log"

	"github.com/isaacd9/miguel/lib/project"
	"github.com/isaacd9/miguel/model/error"
	"github.com/isaacd9/miguel/model/project"
	"github.com/kataras/iris"
)

func parseProject(ctx *iris.Context) (p *projectModel.Project) {
	err := ctx.ReadJSON(&p)
	if err != nil {
		log.Print("Error adding project: " + err.Error())
		ctx.JSON(401, errorModel.Error{
			Message: err.Error(),
		})
	}

	return
}

func ListProjects(ctx *iris.Context) {
	projectList, err := project.ListProjects()
	if err != nil {
		log.Print(err)
	}

	ctx.JSON(iris.StatusOK, projectList)
	return
}

func GetProjectInfo(ctx *iris.Context) {
	id := ctx.Param("id")
	pp, err := project.FindProject(id)

	if err != nil {
		log.Print("Error finding project " + id + ": " + err.Error())
		ctx.JSON(404, errorModel.Error{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(iris.StatusOK, pp)
}

func AddProject(ctx *iris.Context) {
	p := parseProject(ctx)

	err := project.New(p)
	if err != nil {
		log.Print("Error adding project: " + err.Error())
		ctx.JSON(401, errorModel.Error{
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(iris.StatusOK, map[string]string{
		"message": "success",
	})
}

func RemoveProject(ctx *iris.Context) {
	id := ctx.Param("id")
	err := project.Delete(id)

	if err != nil {
		log.Print("Error deleting project: " + err.Error())
		ctx.JSON(401, errorModel.Error{
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(iris.StatusOK, map[string]string{
		"message": "success",
	})
}
