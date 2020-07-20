package dao

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

var instance *GenericDao

var (
	URL_LOCAL = "mongodb://localhost"
)

const (
	DB  = "crud"
	TCP = "tcp"
)

var session *mgo.Session

type GenericDao struct{}

func GetSession() *mgo.Session {
	dialInfo, _ := mgo.ParseURL(URL_LOCAL)

	if session == nil {
		nsession, err := mgo.DialWithInfo(dialInfo)
		if err != nil {
			log.Fatal("#ERROR - MongoDB GetSession: ", err)
		}

		session = nsession
		return nsession.Copy()
	}

	return session.Copy()
}
