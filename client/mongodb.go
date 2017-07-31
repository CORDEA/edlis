package client

import (
	"gopkg.in/mgo.v2"
	"log"
)

const (
	dbName    string = "Edlis"
	slide     string = "slide"
	slideData string = "slideData"
	user      string = "user"
	comment   string = "comment"
)

type MongoDb struct {
	Session *mgo.Session
}

func (m *MongoDb) Open() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatalln(err)
	}
	session.SetMode(mgo.Monotonic, true)

	log.Println("open server.")
	m.Session = session
}

func (m *MongoDb) collection(name string) *mgo.Collection {
	return m.Session.DB(dbName).C(name)
}

func (m *MongoDb) User() *mgo.Collection {
	return m.collection(user)
}

func (m *MongoDb) Comment() *mgo.Collection {
	return m.collection(comment)
}

func (m *MongoDb) Slide() *mgo.Collection {
	return m.collection(slide)
}

func (m *MongoDb) SlideData() *mgo.Collection {
	return m.collection(slideData)
}
