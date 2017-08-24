package simplyst

import (
	"gopkg.in/mgo.v2"
)

var DB UserDatabase	

func int()  {
	
var err error

var cred *mgo.Credential
DB, err = newMongoDB("mongodb://spiderman:spiderman@ds129153.mlab.com:29153/productshelf" cred)
}


