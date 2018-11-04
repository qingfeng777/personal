package moudle

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Init() {
	log.Print("hello")

	db, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Optional. Switch the session to a monotonic behavior.
	db.SetMode(mgo.Monotonic, true)

	c := db.DB("personal").C("article")
	err = c.Insert(&Article{Title: "Ale", Content: "hello mingbai"},
		&Article{Title: "Cla", Content: "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Article{}
	err = c.Find(bson.M{"title": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Content)
}
