package simplyst

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"gopkg.in/mgo.v2"
)

type mongoDB struct {
	conn *mgo.Session
	c    *mgo.Collection
}

var _ UserDatabase = &mongoDB{}

func newMongoDB(addr string, cred *mgo.Credential) (UserDatabase, error) {

	conn, err := mgo.Dial(addr)
	if err != nil {
		return nil, fmt.Errorf("mongo: could not dial:%v", err)
	}

	if cred != nil {
		if err := conn.Login(cred); err != nil {
			return nil, err
		}
	}
	return &mongoDB{
		conn: conn,
		c:    conn.DB("simplystdatabase").C("users"),
	}, nil
}

func (db *mongoDB) Close() {
	db.conn.Close()
}

var maxRand = big.NewInt(1<<63 - 1)

// randomID returns a positive number that fits within an int64.
func randomID() (int64, error) {
	// Get a random number within the range [0, 1<<63-1)
	n, err := rand.Int(rand.Reader, maxRand)
	if err != nil {
		return 0, err
	}
	// Don't assign 0.
	return n.Int64() + 1, nil
}

// AddUser saves a given book, assigning it a new ID.
func (db *mongoDB) Adduser(u *User) (id int64, err error) {
	id, err = randomID()
	if err != nil {
		return 0, fmt.Errorf("mongodb: could not assign an new ID: %v", err)
	}

	u.ID = id
	if err := db.c.Insert(u); err != nil {
		return 0, fmt.Errorf("mongodb: could not add user: %v", err)
	}
	return id, nil
}
