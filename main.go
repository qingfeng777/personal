package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/go-macaron/i18n"
	"gopkg.in/macaron.v1"

	"personal/common"
	"personal/handler"
	"personal/moudle"
)

func main() {
	configPath := flag.String("config", "common/config.toml", "config file's path")
	flag.Parse()

	// 输出config路径
	fmt.Println("config path: [%v]", *configPath)
	common.InitConfig(*configPath)

	moudle.Init()
	runtime.GOMAXPROCS(common.ConfigRef.App.GoMaxPro)

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
	m.Run(7878)
}

func Ping(ctx *macaron.Context) string {
	return "hello %s" + "world"
}
