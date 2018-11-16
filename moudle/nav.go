package moudle

import "gopkg.in/mgo.v2/bson"

//around search
type Site struct {
	Search []string
	Sub    []string
}

type Node struct {
	Name string
	Url  string
}

type Category struct {
	Name string
	Node []Node
}

//around navigation
type Page struct {
	Name  string
	Block []Category
	Id    bson.ObjectId `json:"id" bson:"_id,omitempty"`
}

// two  api


