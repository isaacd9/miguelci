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
	p.UpdatedTime = time.Now()

	err = c.Update(selector, &p)
	if err != nil {
		return nil, err
	}

	return &newBuild, nil
}

func setBuildState(buildId bson.ObjectId, newState projectModel.BuildState) (err error) {
	c := database.Manager.Database.C("projects")
	projectSelector := bson.M{"builds._id": buildId}
	err = c.Update(projectSelector,
		bson.M{"$set": bson.M{"builds.$.state": newState}})

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func queueBuild(build *projectModel.Build) (err error) {
	log.Print("Queueing Build")

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

	err = setBuildState(build.ID, projectModel.BuildWaiting)
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

func GetProjectFromBuild(build *projectModel.Build) (pp *projectModel.Project, err error) {
	c := database.Manager.Database.C("projects")
	p := projectModel.Project{}
	selector := bson.M{"builds._id": build.ID}

	err = c.Find(selector).One(&p)
	if err != nil {
		return nil, err
	}
	return &p, err
}
