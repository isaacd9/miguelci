package routes

import (
	"github.com/isaacd9/miguel/lib/project"
	"github.com/isaacd9/miguel/model/error"
	"github.com/isaacd9/miguel/model/project"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2/bson"
	"log"
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

func AddProject(ctx *iris.Context) {
	p := parseProject(ctx)

	err := project.New(p)
	if err != nil {
		log.Print("Error adding project: " + err.Error())
		ctx.JSON(401, errorModel.Error{
			Message: err.Error(),
		})
	}
	ctx.JSON(iris.StatusOK, map[string]string{
		"message": "success",
	})
}

func RemoveProject(ctx *iris.Context) {
	id := bson.ObjectId(ctx.Param("id"))
	p := projectModel.Project{ID: id}

	err := project.Delete(&p)
	if err != nil {
		log.Print("Error deleting project: " + err.Error())
		ctx.JSON(401, errorModel.Error{
			Message: err.Error(),
		})
	}
	ctx.JSON(iris.StatusOK, map[string]string{
		"message": "success",
	})
}
