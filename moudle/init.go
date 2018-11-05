package moudle

import (
	"log"

	"gopkg.in/mgo.v2"
)

const typeDB, typeArticle, typeManager, typeTag =  "personal", "article", "manager", "tag"

var mdb *mgo.Session

func Init() {
	log.Print("init db")

	var err error
	mdb, err = mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	// Optional. Switch the session to a monotonic behavior.
	//mdb.SetMode(mgo.Monotonic, true)
}
