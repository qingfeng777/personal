package moudle

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func GetArticle()([]Article, error) {
	var result []Article
	return result, ArticleCollection().Find(bson.M{"state": 0}).All(&result)
}

func GetArticleById() {

}

func UpdateArticle() {

}

func AddArticle(article Article)error {
	return ArticleCollection().Insert(&article)
}

func DelArticle() {

}


func ArticleCollection() *mgo.Collection{
	return mdb.DB(typeDB).C(typeArticle)
}
