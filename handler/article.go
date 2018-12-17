package handler

import (
	"gopkg.in/macaron.v1"
	"net/http"
	"personal/server"

	"personal/moudle"
	"personal/utils"
)

func GetArticle(ctx *macaron.Context) {
	art := new(moudle.Article)
	article, err := art.List()
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			Resp{
				Code: http.StatusInternalServerError,
				Msg:  utils.LogErrorF("get article is error ", err.Error()),
			},
		)
	}

	ctx.JSON(http.StatusOK, Resp{Code: http.StatusOK, Data: article})
}

func AddArticle(article moudle.Article, ctx *macaron.Context) {
	if err :=server.CheckAndAdd(article.Category); err != nil{
		ctx.JSON(http.StatusInternalServerError, Resp{Code: http.StatusInternalServerError, Msg: utils.LogErrorF("check category err: ", err.Error())})
	}

	if err := article.Add(); err != nil {
		ctx.JSON(http.StatusInternalServerError, Resp{Code: http.StatusInternalServerError, Msg: utils.LogErrorF("add article is error ", err.Error())})
	}

	ctx.JSON(http.StatusOK, Resp{Code: http.StatusOK, Msg: "success"})
}

func DelArticle(article moudle.Article, ctx *macaron.Context) {
	if err := article.Delete(); err != nil {
		ctx.JSON(http.StatusInternalServerError, Resp{Code: http.StatusInternalServerError, Msg: utils.LogErrorF("delete article is error ", err.Error())})
	}

	ctx.JSON(http.StatusInternalServerError, Resp{Code: http.StatusOK, Msg: "success"})
}

func UpdateArticle(article moudle.Article, ctx *macaron.Context) {
	if err := article.Update(); err != nil {
		ctx.JSON(http.StatusInternalServerError, Resp{Code: http.StatusInternalServerError, Msg: utils.LogErrorF("add article is error ", err.Error())})
	}

	ctx.JSON(http.StatusInternalServerError, Resp{Code: http.StatusOK, Msg: "success"})
}

func GetArticleById(article moudle.Article, ctx *macaron.Context) {
	art, err := article.ById()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Resp{Code: http.StatusInternalServerError, Msg: utils.LogErrorF("get article is error ", err.Error())})
	}

	ctx.JSON(http.StatusInternalServerError, Resp{Code: http.StatusOK, Data:art, Msg: "success"})
}