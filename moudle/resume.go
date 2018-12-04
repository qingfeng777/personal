package moudle

import "gopkg.in/mgo.v2/bson"

type Resume struct {
	Personal          Personal
	ProjectExperience ProjectExperience
	Education         Education
	Skill             Skill
	Testimonials      Testimonials
	Project           Project
	GetTouch          GetTouch
}

type Personal struct {
	UserId       bson.ObjectId `json:"id" bson:"user_id,omitempty"`
	Image        string        `json:"image"`
	Name         string        `json:"name"`
	Position     string        `json:"position"`
	Addr         string        `json:"addr"`
	Introduction string        `json:"introduction"`
	Email        string        `json:"email"`
	Phone        string        `json:"phone"`
	PersonalPage string        `json:"personalPage"`
}

type ProjectExperience struct {
	UserId      bson.ObjectId `json:"id" bson:"user_id,omitempty"`
	Position    string
	CompanyName string
	CompanyAddr string
	Title       string
	Desc        string
	Skill       []string
	StartAt     int64 //  Accuracy months
	EndAt       int64
}

type Education struct {
	UserId      bson.ObjectId `json:"id" bson:"user_id,omitempty"`
	Index       uint8 //show index ,begin from 0
	CollageName string
	StartAt     int64
	EndAt       int64
	Profession  string
	Level       string //高中、大专、本科、硕士、博士
	Desc        string
}

type TopSkill struct {
	Name   string
	Expert float32 //year
	Degree uint8   //1~100
	Desc   string
}

type Skill struct {
	UserId     bson.ObjectId `json:"id" bson:"user_id,omitempty"`
	TopSkill   []TopSkill
	OtherSkill []string
}

type Testimonials struct {
	UserId      bson.ObjectId `json:"id" bson:"user_id,omitempty"`
	Desc        string
	Company     string
	LeaderTitle string
	Time        int64
}

type Project struct {
	UserId bson.ObjectId `json:"id" bson:"user_id,omitempty"`
	Name   string
	Skill  []string //two or three
	Url    string
	Img    string
	Scope  string // java or go .  back or front 
}
type GetTouch struct {
	UserId       bson.ObjectId `json:"id" bson:"user_id,omitempty"`
	Introduction string
	HelpScope    []string
	Email        string
	Phone        string
}
