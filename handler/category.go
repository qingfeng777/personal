package handler

import (
	"net/http"

	"github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"

	"personal/moudle"
	"personal/utils"
)

func cateGoryRouter(m *macaron.Macaron) {
	m.Get("/category", GetCategory)
	m.Post("/category", binding.Bind(moudle.Article{}), AddCategory)
	m.Delete("/category/:Id", binding.Bind(moudle.Article{}), DelCategory)
}

func GetCategory(ctx *macaron.Context) {
	art := new(moudle.Category)
	category, err := art.List()
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			Resp{
				Code: http.StatusInternalServerError,
				Msg:  utils.LogErrorF("get category is error ", err.Error()),
			},
		)
	}

	ctx.JSON(http.StatusOK, Resp{Code: http.StatusOK, Data: category})
}

func AddCategory(category moudle.Category, ctx *macaron.Context) {
	if err := category.Add(); err != nil {
		ctx.JSON(http.StatusInternalServerError, Resp{Code: http.StatusInternalServerError, Msg: utils.LogErrorF("add category is error ", err.Error())})
	}

	ctx.JSON(http.StatusOK, Resp{Code: http.StatusOK, Msg: "success"})
}

func DelCategory(Category moudle.Category, ctx *macaron.Context) {
	if err := Category.Delete(); err != nil {
		ctx.JSON(http.StatusInternalServerError, Resp{Code: http.StatusInternalServerError, Msg: utils.LogErrorF("delete category is error ", err.Error())})
	}

	ctx.JSON(http.StatusInternalServerError, Resp{Code: http.StatusOK, Msg: "success"})
}
