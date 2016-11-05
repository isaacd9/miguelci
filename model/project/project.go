package projectModel

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	CredentialSet struct {
		User string `bson:"name"`
		Key  string `bson:"key"`
	}

	Project struct {
		ID          bson.ObjectId `bson:"_id,omitempty"`
		Name        string        `bson:"name"`
		URL         string        `bson:"URL"`
		Auth        CredentialSet `bson:"credentials"`
		Builds      []Build       `bson:"builds"`
		CreatedTime time.Time
		UpdatedTime time.Time
	}
)
