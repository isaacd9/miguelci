package projectModel

import (
	"labix/v2/mgo/bson"
)

type Build struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
}
