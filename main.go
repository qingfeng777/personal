package main

import (
	"personal/handler"

	//_ "golang.org/x/crypto"
	"github.com/go-macaron/i18n"

	_ "github.com/jtolds/gls"

	"gopkg.in/macaron.v1"

	"log"
	"personal/moudle"
)

func main() {

	moudle.Init()

	m := macaron.Classic()
	m.Use(i18n.I18n(i18n.Options{
		Langs: []string{"en-US", "zh-CN"},
		Names: []string{"English", "简体中文"},
	}))

	m.Get("/languages", func(locale i18n.Locale) string {
		return "current language is" + locale.Language()
	})

	m.Group("/ming", func() {
		m.Get("/ping", Ping)

		m.Group("/article", func() {
			handler.Router(m)
		}, nil, nil)
	}, nil, nil)

	log.Print("hello")
	m.Run("8000")
}

func Ping(ctx *macaron.Context) string {
	return ctx.Tr("hello %s", "world")
}
