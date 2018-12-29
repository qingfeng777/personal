package moudle

import (
	"log"
	"personal/common"

	"gopkg.in/mgo.v2"
)

const typeDB, typeArticle, typeUser, typeTag, typeCategory = "personal", "article", "user", "tag", "category"

var mdb *mgo.Session

func Init() {
	log.Print("init db")

	var err error
	mdb, err = mgo.Dial(common.ConfigRef.DB.Url)
	if err != nil {
		log.Fatal(err)
	}
	//defer mdb.Close()

	// Optional. Switch the session to a monotonic behavior.
	//mdb.SetMode(mgo.Monotonic, true)
}

type Pagination struct {
	Admin bool   `json:"admin"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Key   string `json:"key"`
}
