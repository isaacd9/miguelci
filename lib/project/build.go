package project

import (
	"encoding/json"
	"log"

	"github.com/isaacd9/miguel/lib/database"
	"github.com/isaacd9/miguel/lib/queue"
	"github.com/isaacd9/miguel/model/project"
	"gopkg.in/mgo.v2/bson"
	"time"
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

func createBuild(id string) (build *projectModel.Build, err error) {
	selector := bson.M{"_id": bson.ObjectIdHex(id)}
	c := database.Manager.Database.C("projects")
	p := projectModel.Project{}
	err = c.Find(selector).One(&p)

	if err != nil {
		return nil, err
	}

	buildNum := len(p.Builds) + 1
	newBuild := projectModel.Build{ID: bson.NewObjectId(), Number: buildNum, State: projectModel.BuildCreated, CreatedTime: time.Now(), UpdatedTime: time.Now()}
	p.Builds = append(p.Builds, newBuild)

	err = c.Update(selector, &p)
	if err != nil {
		return nil, err
	}

	return &newBuild, nil
}

func queueBuild(build *projectModel.Build) (err error) {
	log.Print("Queueing Build")
	//	buildSelector := bson.M{"builds": bson.M{"$slice": -1}}
	//
	//	err = c.Find(projectSelector).Select(buildSelector).One(&p)
	//	b := p.Builds[0]

	if build.State != projectModel.BuildCreated {
		return
	}

	q := queue.Manager.Client

	build.State = projectModel.BuildWaiting
	res, err := json.Marshal(build)
	if err != nil {
		log.Print(err)
		return err
	}

	c := database.Manager.Database.C("projects")
	//p := projectModel.Project{}

	projectSelector := bson.M{"builds._id": build.ID}
	err = c.Update(projectSelector, bson.M{"$set": bson.M{"builds.$.state": projectModel.BuildWaiting}})
	if err != nil {
		log.Print(err)
		return err
	}

	q.LPush("builds", string(res))
	return nil
}

func NewBuild(projectId string) (err error) {
	newBuild, err := createBuild(projectId)
	if err != nil {
		return err
	}

	go queueBuild(newBuild)

	return nil
}
