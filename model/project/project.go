package projectModel

import (
	"github.com/isaacd9/miguel/model/build"
	"gopkg.in/mgo.v2/bson"
)

type (
	CredentialSet struct {
		user string `bson:"name"`
		key  string `bson:"key"`
	}

	Project struct {
		ID     bson.ObjectId      `bson:"_id,omitempty"`
		Name   string             `bson:"name"`
		URL    string             `bson:"URL"`
		Auth   CredentialSet      `bson:"credentials"`
		Builds []buildModel.Build `bson:"builds"`
	}
)
