package buildModel

import (
	"gopkg.in/mgo.v2/bson"
)

type Build struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
}
