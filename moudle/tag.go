package moudle

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TagOperation interface {
	List() ([]Tag, error)
	Add() error
	Delete() error
}


type Tag struct {
	Name     string   `json:"name"`
	Slug     string   `json:"slug"`
	Desc     string   `json:"desc"`
	Extends  []Extend `json:"extends"`
	CreateAt int64    `json:"create_at"`
	Id       int      `json:"id"`
	State    int      `json:"state"`
}

func (Tag) List() ([]Tag, error) {
	var tag []Tag
	return tag, tagCollection().Find(bson.M{"state": 0}).All(&tag)
}

func (tag *Tag) Add() error {
	return tagCollection().Insert(tag)
}

func (tag *Tag) Delete() error {
	return tagCollection().Remove(bson.M{"_id":tag.Id})
}

func tagCollection() *mgo.Collection {
	return mdb.DB(typeDB).C(typeTag)
}
