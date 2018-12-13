package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-macaron/binding"
	"personal/moudle"

	"gopkg.in/macaron.v1"
)

func Router(m *macaron.Macaron) {
	m.Get("/article", GetArticle)
	m.Post("/article", binding.Bind(moudle.Article{}), AddArticle)
	m.Delete("/article/:Id", binding.Bind(moudle.Article{}), DelArticle)
	m.Get("/article/:Id",binding.Bind(moudle.Article{}), GetArticleById)
	m.Put("/article", binding.Bind(moudle.Article{}), UpdateArticle)
	cateGoryRouter(m)
}

type Resp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func (resp Resp) String() string {
	byte, _ := json.Marshal(resp)
	return fmt.Sprintf(string(byte))
}
