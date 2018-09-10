package mongo

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
)

type MongoHandler struct {
	Db      *mgo.Database
	Session *mgo.Session
}

func NewMongoHandler(appc AppConfig, dbc DbConfig) *MongoHandler {
	session, err := mgo.Dial(fmt.Sprintf("mongodb://%s:%d", dbc.Host, dbc.Port))
	if err != nil {
		panic(fmt.Sprintf("Initialize mongodb error:%v", err))
	}
	database := session.DB(dbc.Database)
	if err = session.Ping(); err != nil {
		panic(fmt.Sprintf("MongoDB execute ping error:%v", err))
	}
	log.Println("MongoDB initialize success.")
	mgo.SetDebug(appc.Debug)
	mongoHandler := &MongoHandler{
		Db:      database,
		Session: session,
	}
	return mongoHandler

}

func (mh *MongoHandler) FindOne(collection string, query interface{}, res interface{}) error {
	return mh.Db.C(collection).Find(query).One(res)
}

func (mh *MongoHandler) FindAll(collection string, query interface{}, res interface{}) error {
	return mh.Db.C(collection).Find(query).All(res)
}

func (mh *MongoHandler) Upsert(collection string, query interface{}, upsert interface{}) error {
	_, err := mh.Db.C(collection).Upsert(query, upsert)
	return err
}

func (mh *MongoHandler) Insert(collection string, object interface{}) error {
	return mh.Db.C(collection).Insert(object)
}

func (mh *MongoHandler) Delete(collection string, query interface{}) error {
	return mh.Db.C(collection).Remove(query)
}
