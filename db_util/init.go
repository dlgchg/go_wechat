package db_util

import "gopkg.in/mgo.v2"

func ConnectMongoDB() *mgo.Session {
	dial, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
		return nil
	}

	defer dial.Close()

	dial.SetMode(mgo.Monotonic, true)

	return dial
}
