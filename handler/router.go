package handler

import (
	"encoding/json"
	"fmt"

	"personal/moudle"
	"personal/pkg/binding"

	"gopkg.in/macaron.v1"
)

func Router(m *macaron.Macaron) {
	m.Get("/article/:limit/:page", GetArticle)
	m.Get("/article/search", binding.Bind(moudle.Pagination{}), SearchArticle) ///:key/:limit/:page
	m.Post("/article", binding.Bind(moudle.Article{}), AddArticle)
	m.Delete("/article/:Id", binding.Bind(moudle.Article{}), DelArticle)
	m.Get("/article/:Id", binding.Bind(moudle.Article{}), GetArticleById)
	m.Put("/article", binding.Bind(moudle.Article{}), UpdateArticle)
	cateGoryRouter(m)
}

type Resp struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	Total int         `json:"total"`
}

func (resp Resp) String() string {
	byte, _ := json.Marshal(resp)
	return fmt.Sprintf(string(byte))
}
