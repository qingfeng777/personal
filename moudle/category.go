package moudle

import (
"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
)

type CategoryOperation interface {
	List() ([]Category, error)
	Add() error
	Delete() error
}

type Category struct {
	Name     string   `json:"name"`
	Slug     string   `json:"slug"`
	Desc     string   `json:"desc"`
	Extends  []Extend `json:"extends"`
	CreateAt int64    `json:"createAt"`
	Id       int      `json:"id"`
	State    int      `json:"state"`
}

func (Category) List() ([]Category, error) {
	var category []Category
	return category, categoryCollection().Find(bson.M{"state": 0}).All(&category)
}

func (category *Category) Add() error {
	return categoryCollection().Insert(category)
}

func (category *Category) Delete() error {
	return categoryCollection().Remove(bson.M{"_id":category.Id})
}

func categoryCollection() *mgo.Collection {
	return mdb.DB(typeDB).C(typeCategory)
}
