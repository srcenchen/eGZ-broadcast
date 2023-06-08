package controller

import (
	v1 "broadcast_back_end/api/musicResource"
	"broadcast_back_end/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"path"
	"strings"
	"time"
)

type MusicResource struct{}

// GetMusicList 获取音乐列表
func (MusicResource) GetMusicList(ctx context.Context, req *v1.MusicListReq) (res *[]v1.MusicListRes, err error) {
	_ = req
	g.RequestFromCtx(ctx).Response.WriteJson(dao.GetMusicResourceList())
	return
}

// UploadMusic 上传音乐文件
func (MusicResource) UploadMusic(ctx context.Context, req *v1.UploadMusicReq) (res *v1.UploadMusicRes, err error) {
	_ = req
	_ = ctx
	file := req.File                               // 获取文件
	file.Filename = strings.ToLower(file.Filename) // 降为小写
	timeStamp := gconv.String(time.Now().Unix())   // 获取时间戳
	randomFileName := grand.Letters(16)            // 生成16位随机字符串
	file.Filename = timeStamp + randomFileName + path.Ext(file.Filename)
	fileName, err := file.Save("./resource/music")
	res = &v1.UploadMusicRes{
		FileName: fileName,
	}
	return
}

// AddMusic 添加音乐资源
func (MusicResource) AddMusic(ctx context.Context, req *v1.AddMusicReq) (res *v1.AddMusicRes, err error) {
	_ = req
	dao.AddMusicResource(req.Title, req.FileName)
	res = &v1.AddMusicRes{
		Message: "success",
	}
	return
}

// DeleteMusic 删除音乐资源
func (MusicResource) DeleteMusic(ctx context.Context, req *v1.DeleteMusicReq) (res *v1.DeleteMusicRes, err error) {
	_ = req
	dao.DeleteMusicResource(req.Id)
	res = &v1.DeleteMusicRes{
		Message: "success",
	}
	return
}
