package main

import (
	"github.com/go-macaron/i18n"
	"gopkg.in/macaron.v1"

	"personal/handler"
	"personal/moudle"
)

func main() {

	moudle.Init()

	m := macaron.Classic()
	m.Use(macaron.Renderer())

	m.Get("/languages", func(locale i18n.Locale) string {
		return "current language is" + locale.Language()
	})

	m.Group("/ming", func() {
		m.Get("/ping", Ping)

		handler.Router(m)
	})
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}

func Ping(ctx *macaron.Context) string {
	return "hello %s" + "world"
}
