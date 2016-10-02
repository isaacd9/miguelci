package database

import (
	"gopkg.in/mgo.v2"
	"log"
)

type SessionManager struct {
	Session  *mgo.Session
	Database *mgo.Database
}

var (
	Manager SessionManager
)

func Connect() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}

	Manager.Session = session
	Manager.Session.SetMode(mgo.Monotonic, true)

	Manager.Database = Manager.Session.DB("miguel")
}

func Disconnect() {
	Manager.Session.Close()
}
