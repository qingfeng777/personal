package handler

import (
	"net/http"

	"personal/pkg/macaron"
	"personal/moudle"
	"personal/server"
	"personal/utils"
)

func SearchArticle(ctx *macaron.Context, pagination moudle.Pagination) {
	art := new(moudle.Article)
	article, err := art.Search(pagination)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Resp{Code: http.StatusInternalServerError, Msg: utils.LogErrorF("search article  error ", err.Error())})
	}

	total, err := art.Count(pagination)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Resp{Code: http.StatusInternalServerError, Msg: utils.LogErrorF("count article  error ", err.Error())})
	}

	ctx.JSON(http.StatusOK, Resp{Code: http.StatusOK, Data: article, Total:total})
}

func GetArticle(ctx *macaron.Context) {
	art := new(moudle.Article)
	article, err := art.List(ctx.ParamsInt(":limit"), ctx.ParamsInt(":page"))
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
	if err := server.CheckAndAdd(article.Category); err != nil {
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

	ctx.JSON(http.StatusOK, Resp{Code: http.StatusOK, Msg: "success"})
}

func UpdateArticle(article moudle.Article, ctx *macaron.Context) {
	if err := server.CheckAndAdd(article.Category); err != nil {
		ctx.JSON(http.StatusInternalServerError, Resp{Code: http.StatusInternalServerError, Msg: utils.LogErrorF("check category err: ", err.Error())})
	}

	if err := article.Update(); err != nil {
		ctx.JSON(http.StatusInternalServerError, Resp{Code: http.StatusInternalServerError, Msg: utils.LogErrorF("add article is error ", err.Error())})
	}

	ctx.JSON(http.StatusOK, Resp{Code: http.StatusOK, Msg: "success"})
}

func GetArticleById(article moudle.Article, ctx *macaron.Context) {
	art, err := article.ById()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Resp{Code: http.StatusInternalServerError, Msg: utils.LogErrorF("get article is error ", err.Error())})
	}

	ctx.JSON(http.StatusOK, Resp{Code: http.StatusOK, Data: art, Msg: "success"})
}
