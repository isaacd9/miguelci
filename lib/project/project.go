package project

import (
	"errors"
	"time"

	"github.com/isaacd9/miguel/lib/database"
	"github.com/isaacd9/miguel/model/project"
	"gopkg.in/mgo.v2/bson"
)

func ListProjects() (project []*projectModel.Project, err error) {
	c := database.Manager.Database.C("projects")

	c.Find(nil).Limit(100).All(&project)
	if err != nil {
		return nil, err
	}

	if project == nil {
		return make([]*projectModel.Project, 0), nil
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

	p.CreatedTime = time.Now()
	p.UpdatedTime = time.Now()

	err = c.Insert(p)
	if err != nil {
		return err
	}

	return nil
}

func Delete(id string) (err error) {
	c := database.Manager.Database.C("projects")
	err = c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})

	if err != nil {
		return err
	}

	return
}

func FindProject(id string) (pp *projectModel.Project, err error) {
	c := database.Manager.Database.C("projects")
	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&pp)

	if err != nil {
		return nil, err
	}

	return pp, nil
}
