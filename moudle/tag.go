package moudle

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Tag struct {
	Name     string   `json:"name"`
	Slug     string   `json:"slug"`
	Desc     string   `json:"desc"`
	Extends  []Extend `json:"extends"`
	CreateAt int64    `json:"create_at"`
	id       int      `json:"id"`
	state    int      `json:"state"`
}

type TagClient interface {
	List() ([]Tag, error)
	Add() error
}

func (Tag) List() ([]Tag, error) {
	var tag []Tag
	return tag, TagCollection().Find(bson.M{"state": 0}).All(&tag)
}

func (tag *Tag) Add() error {
	return TagCollection().Insert(tag)
}

func TagCollection() *mgo.Collection {
	return mdb.DB(typeDB).C(typeTag)
}
