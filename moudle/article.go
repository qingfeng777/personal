package moudle

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
)


type Counter struct {
	View    int `json:"view"`
	Like    int `json:"like"`
	Comment int `json:"comment"`
}

type Extend struct {
	name string
	val  string
}

type Article struct {
	Counter  Counter  `json:"counter"`
	KeyWord  []string `json:"key_word"`
	State    string   `json:"state"`
	Public   int      `json:"public"` // 0 public ,1 private
	Origin   int      `json:"origin"` // 0 or 1
	Pwd      string   `json:"pwd"`
	Tag      []string `json:"tag"`
	Category string   `json:"category"`
	Title    string   `json:"title"`
	Desc     string   `json:"desc"`
	Content  string   `json:"content"`
	Thumb    string   `json:"thumb"`
	Extends  []Extend `json:"extends"`
	CreateAt int64    `json:"create_at"`
	UpdateAt int64    `json:"update_at"`
	Id       int      `json:"id"`
}

func ListArticle()[]Article {

	c := mdb.DB("personal").C("article")
	err := c.Insert(&Article{Title: "Ale", Content: "hello mingbai"},
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
	return nil
}

func GetArticle() {

}

func UpdateArticle() {

}

func AddArticle(article Article)error {
	c := mdb.DB(typeDB).C(typeArticle)
	return c.Insert(&article)
}

func DelArticle() {

}
