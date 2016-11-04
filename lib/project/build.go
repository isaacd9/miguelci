package project

import (
	"github.com/isaacd9/miguel/lib/database"
	"github.com/isaacd9/miguel/model/project"
	"gopkg.in/mgo.v2/bson"
)

func ListBuilds(id string) (builds *[]projectModel.Build, err error) {
	c := database.Manager.Database.C("projects")
	p := projectModel.Project{}
	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&p)

	if err != nil {
		return nil, err
	}

	return &p.Builds, err
}

func NewBuild(id string) (err error) {
	selector := bson.M{"_id": bson.ObjectIdHex(id)}
	c := database.Manager.Database.C("projects")
	p := projectModel.Project{}
	err = c.Find(selector).One(&p)

	if err != nil {
		return err
	}

	buildNum := len(p.Builds) + 1
	newBuild := projectModel.Build{ID: bson.NewObjectId(), Number: buildNum, State: projectModel.BuildCreated}
	p.Builds = append(p.Builds, newBuild)

	err = c.Update(selector, &p)
	if err != nil {
		return err
	}

	return nil
}
