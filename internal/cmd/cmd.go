package cmd

import (
	"broadcast_back_end/internal/controller"
	"broadcast_back_end/internal/model/entity"
	"context"
	"os"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name: "main",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			initDir()         // 初始化文件夹
			entity.InitData() // 初始化数据
			s := g.Server()
			controller.Router(s)
			s.Run()
			return nil
		},
	}
)

// createDir /** 创建文件夹
func createDir(path string) {
	// 检测是否存在 path 文件夹 如果不存在则创建
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, os.ModePerm)
	}
}

// initDir /** 初始化文件夹
func initDir() {
	// 创建文件夹
	createDir("./resource")
	createDir("./resource/database")
	createDir("./resource/music")
}
