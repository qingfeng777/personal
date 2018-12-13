package moudle

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)


type ArticleOperation interface {
	List() ([]Article, error)
	ById() (*Article, error)
	Update() error
	Add() error
	Delete() error
}


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
	Counter  Counter       `json:"counter"`
	KeyWord  string        `json:"keyWord"`
	State    int           `json:"state"`
	Public   int           `json:"public"` // 0 public ,1 private
	Origin   int           `json:"origin"` // 0 or 1
	Pwd      string        `json:"pwd"`
	Tag      []string      `json:"tag"`
	Category string        `json:"category"`
	Title    string        `json:"title" binding:"Required"`
	Desc     string        `json:"desc"`
	Content  string        `json:"content" binding:"Required"`
	Thumb    string        `json:"thumb"`
	Extends  []Extend      `json:"extends"`
	CreateAt int64         `json:"createAt"`
	UpdateAt int64         `json:"updateAt"`
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
}

func (*Article) List() ([]Article, error) {
	var result []Article
	return result, articleCollection().Find(bson.M{"state": 0}).All(&result)
}

func (article *Article) ById() (*Article, error) {
	return article, articleCollection().Find(bson.M{"id": article.Id}).One(article)
}

func (article *Article) Update() error {
	article.UpdateAt = time.Now().Unix()
	return articleCollection().Update(bson.M{"id": article.Id}, article)
}

func (article *Article) Add() error {
	article.CreateAt = time.Now().Unix()
	return articleCollection().Insert(article)
}

func (article *Article) Delete() error {
	article.State = 1
	return articleCollection().Update(bson.M{"id": article.Id}, article)
}

func articleCollection() *mgo.Collection {
	return mdb.DB(typeDB).C(typeArticle)
}
