package handler

import (
	"github.com/go-macaron/binding"
	"personal/moudle"

	"gopkg.in/macaron.v1"
)

func Router(m *macaron.Macaron) {
	m.Get("/article", GetArticle)
	m.Post("/article", binding.Bind(moudle.Article{}), AddArticle)
}
