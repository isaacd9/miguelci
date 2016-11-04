package buildModel

import (
	"gopkg.in/mgo.v2/bson"
)

type BuildState string

const (
	Created  BuildState = "CREATED"
	Waiting             = "WAITING"
	Started             = "STARTED"
	Finished            = "FINISHED"
	Failed              = "FAILED"
)

type (
	Build struct {
		ID     bson.ObjectId `bson:"_id,omitempty"`
		Number int           `bson:"Number"`
		State  BuildState    `bson:"state"`
	}
)
