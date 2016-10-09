package project

import (
	"errors"
	"github.com/isaacd9/miguel/lib/database"
	"github.com/isaacd9/miguel/model/project"
	"log"
)

func ListProjects() (project []*projectModel.Project, err error) {
	c := database.Manager.Database.C("projects")

	c.Find(nil).Limit(100).All(&project)
	if err != nil {
		return nil, err
	}

	return
}

func New(p *projectModel.Project) (err error) {
	c := database.Manager.Database.C("projects")

	if p.Name == "" {
		return errors.New("Project name is not set")
	}

	if p.URL == "" {
		return errors.New("Project URL is not set")
	}

	err = c.Insert(p)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func Delete(p *projectModel.Project) (err error) {
	c := database.Manager.Database.C("projects")

	err = c.Remove(p)

	if err != nil {
		return err
	}

	return
}
