package project

import (
	"github.com/isaacd9/miguel/lib/database"
	"github.com/isaacd9/miguel/model/project"
	"log"
)

func ListProjects() (project []*projectModel.Project, err error) {
	var result []projectModel.Project
	c := database.Manager.Database.C("projects")

	iter := c.Find(nil).Limit(100).Iter()
	err = iter.All(&result)
	if err != nil {
		return nil, err
	}

	return
}

func NewProject(p *projectModel.Project) (err error) {
	c := database.Manager.Database.C("projects")

	err = c.Insert(p)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
