package simplyst

import (
	"fmt"
	"gopkg.in/mgo.v2"
)



type mongoDB struct{

	conn *mgo.Session
	c    *mgo.Collection
}

var _ UserDatabase = &mongoDB{}

func newMongoDB(addr string, cred *mgo.Credential) (UserDatabase, error){

	conn, err := mgo.Dial(addr)
	if err !=nil {
		return nil,fmt.Errorf("mongo: could not dial:%v", err)
	}

	if cred != nil{
		if err := conn.Login(cred); err != nil {
			return nil, err
		}
	}
	return &mongoDB{
		conn: conn,
		c: conn.DB("productshelf").C("products"),
	},nil	
}

func(db *mongoDB) Close(){
	db.conn.Close()
}