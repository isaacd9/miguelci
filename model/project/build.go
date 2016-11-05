package projectModel

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type BuildState string

const (
	BuildCreated  BuildState = "CREATED"
	BuildWaiting             = "WAITING"
	BuildStarted             = "STARTED"
	BuildFinished            = "FINISHED"
	BuildFailed              = "FAILED"
)

type (
	Build struct {
		ID          bson.ObjectId `bson:"_id,omitempty"`
		Number      int           `bson:"Number"`
		State       BuildState    `bson:"state"`
		CreatedTime time.Time
		UpdatedTime time.Time
	}
)
