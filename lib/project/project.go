package project

import (
	"errors"

	"github.com/isaacd9/miguel/lib/database"
	"github.com/isaacd9/miguel/model/build"
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

func ListBuilds(id string) (builds *[]buildModel.Build, err error) {
	c := database.Manager.Database.C("projects")
	p := projectModel.Project{}
	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&p)

	if err != nil {
		return nil, err
	}

	return &p.Builds, err
}

func FindProject(id string) (pp *projectModel.Project, err error) {
	c := database.Manager.Database.C("projects")
	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&pp)

	if err != nil {
		return nil, err
	}

	return pp, nil
}

func StartBuild(id string) (err error) {
	selector := bson.M{"_id": bson.ObjectIdHex(id)}
	c := database.Manager.Database.C("projects")
	p := projectModel.Project{}
	err = c.Find(selector).One(&p)

	if err != nil {
		return err
	}

	buildNum := len(p.Builds) + 1
	newBuild := buildModel.Build{ID: bson.NewObjectId(), Number: buildNum, State: buildModel.Created}
	p.Builds = append(p.Builds, newBuild)

	err = c.Update(selector, &p)
	if err != nil {
		return err
	}

	return nil
}
