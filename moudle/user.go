package moudle

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserOperation interface {
	Add() error
	Get() (*User, error)
}

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string
	Pwd      string
	Email    string
	Phone    string
	CreateAt int64
	UpdateAt int64
	ActiveAt int64
	Male     int8
	state    int8 // 0 or 1,  1 disable
}

func (usr *User) Get() (*User, error) {
	return usr, userCollection().Find(bson.M{"Name": usr.Name}).One(&usr)
}

func (usr *User) Add() error {
	return userCollection().Insert(usr)
}

func userCollection() *mgo.Collection {
	return mdb.DB(typeDB).C(typeArticle)
}
