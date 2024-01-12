package api

import (
	"github.com/iproj/file-rotatelogs"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/kataras/iris/v12/middleware/cors"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/middleware/requestid"
)

func NewApp(name string, debug bool) *iris.Application {
	ac := makeAccessLog(name)
	app := iris.New()
	app.ConfigureHost(func(su *iris.Supervisor) {
		su.RegisterOnShutdown(func() {
			ac.Close()
		})
	})
	app.UseRouter(ac.Handler)

	ac.AddOutput(app.Logger().Printer)
	if debug {
		app.Logger().SetLevel("debug")
	}

	app.UseRouter(requestid.New())
	app.UseRouter(recover.New())
	app.UseRouter(cors.New().ExtractOriginFunc(cors.DefaultOriginExtractor).
		ReferrerPolicy(cors.NoReferrerWhenDowngrade).
		AllowOriginFunc(cors.AllowAnyOrigin).Handler())

	return app
}

func makeAccessLog(name string) *accesslog.AccessLog {
	pathToAccessLog := "./" + name + ".%m%d.log"
	w, err := rotatelogs.New(
		pathToAccessLog,
		rotatelogs.WithMaxAge(-1),
		rotatelogs.WithRotationCount(15),
	)
	if err != nil {
		panic(err)
	}
	ac := accesslog.New(w)

	// for debug
	//ac.SetFormatter(&accesslog.Template{
	//	Text: "{{.Now.Format .TimeFormat}}|{{.Latency}}|{{.Code}}|{{.Method}}|{{.Path}}|{{.IP}}\n",
	//})

	//ac := accesslog.New(bufio.NewWriter(w))
	return ac
}
