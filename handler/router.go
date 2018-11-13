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
