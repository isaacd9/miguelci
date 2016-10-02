package routes

import (
	"github.com/isaacd9/miguel/lib/project"
	"github.com/isaacd9/miguel/model/project"
	"github.com/kataras/iris"
	"log"
)

func ListProjects(ctx *iris.Context) {
	projectList, err := project.ListProjects()
	if err != nil {
		log.Print(err)
	}

	ctx.JSON(iris.StatusOK, projectList)
	return
}

func AddProject(ctx *iris.Context) {
	p := projectModel.Project{}
	err := ctx.ReadJSON(&p)
	if err != nil {
		log.Print("Error when reading form: " + err.Error())
		ctx.JSON(401, iris.Map{
			"Message": err.Error(),
		})
	}

	project.NewProject(&p)
}
