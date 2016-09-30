package database

import (
	"gopkg.in/mgo.v2"
	"log"
)

type SessionManager struct {
	session  *mgo.Session
	database *mgo.Database
}

var (
	Manager SessionManager
)

func Connect() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}

	Manager.session = session
	Manager.session.SetMode(mgo.Monotonic, true)

	Manager.database = Manager.session.DB("miguel")
}

func Disconnect() {
	Manager.session.Close()
}
