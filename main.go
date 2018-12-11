package main

import (
	"flag"
	"fmt"
	"net/http"
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
	m.Use(macaron.Recovery())
	//m.Use(session.Sessioner())

	//m.Use(cors.AllowAll)
	m.Use(setCorsAllow)

	//m.Use(csrf.Csrfer(csrf.Options{SetHeader: true, SetCookie: false}))

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
	//m.NotFound(Ping)
	m.Run(7878)
}

func setCorsAllow(ctx *macaron.Context) {
	// 设置为可跨域访问
	ctx.Header().Add("Access-Control-Allow-Origin", "*")
	ctx.Header().Add("Access-Control-Allow-Methods", "OPTIONS,POST,GET,OPTIONS,DELETE,PUT")
	ctx.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	ctx.Header().Add("Access-Control-Allow-Credentials", "true")
	//ctx.JSON(200, nil)

	if ctx.Req.Method == "OPTIONS" {
		ctx.JSON(http.StatusOK, "Options Request!")
	}
	ctx.Next()

}

func Ping(ctx *macaron.Context) string {
	return "hello %s" + "world"
}
