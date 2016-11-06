package main

import (
	"encoding/json"
	"github.com/isaacd9/miguel/lib/database"
	"github.com/isaacd9/miguel/lib/project"
	"github.com/isaacd9/miguel/lib/queue"
	"github.com/isaacd9/miguel/model/project"
	"log"
	"time"
)

func parseBuild(build string) (b *projectModel.Build, err error) {
	bb := projectModel.Build{}
	err = json.Unmarshal([]byte(build), &bb)
	if err != nil {
		return nil, err
	}

	return &bb, nil
}

func processBuild(build string) (err error) {
	log.Print("Processing")

	ret, err := parseBuild(build)
	if err != nil {
		log.Print(err)
		panic(err)
	}

	p, err := project.GetProjectFromBuild(ret)
	if err != nil {
		log.Print(err)
		panic(err)
	}

	log.Print(p)

	return nil
}

func main() {
	database.Connect()
	queue.Connect()

	q := queue.Manager.Client

	for {
		t, err := time.ParseDuration("100s")
		if err != nil {
			panic(err)
		}
		ret := q.BRPop(t, "builds")
		st, err := ret.Result()
		if st != nil {
			if err != nil {
				panic(err)
			}

			go processBuild(st[1])
		}
	}
}
