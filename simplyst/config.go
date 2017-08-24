package simplyst

import (
	"log"

	"gopkg.in/mgo.v2"
)

// DB UserDatabase
var DB UserDatabase

func int() {

	var err error

	var cred *mgo.Credential
	DB, err = newMongoDB("mongodb://kashyaksupport:kashyaksupport123@ds151973.mlab.com:51973/simplystdatabase", cred)
	if err != nil {
		log.Fatal(err)
	}
}
