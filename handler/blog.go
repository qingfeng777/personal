package handler

import (
	"gopkg.in/macaron.v1"
	"log"
	"personal/moudle"
	"time"
)

func GetArticle(ctx *macaron.Context)Resp {
	article, err := moudle.GetArticle()
	if err !=nil {
		log.Println("")
		return Resp{Code:500, Msg:"something is err"}
	}

	return Resp{Code:200, Data:article}
}

func AddArticle(article moudle.Article) Resp {
	article.CreateAt = time.Now().Unix()
	if err:=moudle.AddArticle(article);err !=nil{
		log.Print("add article err :", err.Error())
		return Resp{}
	}

	return Resp{Code:200, Data:article}
}
