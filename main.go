package main

import (
	"fmt"

	//_ "golang.org/x/crypto"

	_ "github.com/jtolds/gls"

	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"personal/moudle"
)

func main() {

	log.Print("hello")

	db, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Optional. Switch the session to a monotonic behavior.
	db.SetMode(mgo.Monotonic, true)

	c := db.DB("personal").C("article")
	err = c.Insert(&moudle.Article{Title: "Ale", Content: "hello mingbai"},
		&moudle.Article{Title: "Cla", Content: "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := moudle.Article{}
	err = c.Find(bson.M{"title": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Content)

	/*moudle.Init()

	m := macaron.Classic()
	m.Use(i18n.I18n(i18n.Options{
		Langs: []string{"en-US", "zh-CN"},
		Names: []string{"English", "简体中文"},
	}))

	m.Get("/languages", func(locale i18n.Locale) string {
		return "current language is" + locale.Lang
	})

	m.Group("/ming", func() {
		m.Get("/ping", Ping)

		m.Group("/article", func() {
			handler.Router(m)
		}, nil, nil)
	}, nil, nil)

	log.Print("hello")*/
	//m.Run("8000")
}

func Ping(ctx *macaron.Context) string {
	return ctx.Tr("hello %s", "world")
}
