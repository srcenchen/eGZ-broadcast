package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func Router(s *ghttp.Server) {
	// 静态资源绑定
	s.SetIndexFolder(true)
	s.SetServerRoot("public")
	s.BindStatusHandler(404, func(r *ghttp.Request) {
		r.Response.ServeFile("public/index.html")
	})

	api := s.Group("/api")
	api.Middleware(ghttp.MiddlewareHandlerResponse)

	musicResource := api.Group("/music-resource") // 音乐资源组
	musicResource.Bind(new(MusicResource))

	music := api.Group("/music-ctl") // 音乐控制组
	music.Bind(new(MusicCtl))

	taskGroup := api.Group("/task-group") // 任务组
	taskGroup.Bind(new(TaskGroup))
}
