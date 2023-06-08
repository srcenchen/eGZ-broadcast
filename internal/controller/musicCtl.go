package controller

import (
	"broadcast_back_end/api/musicCtl"
	"broadcast_back_end/internal/dao"
	"broadcast_back_end/internal/service/mediaService"
	"broadcast_back_end/internal/service/taskGroupService"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type MusicCtl struct{}

// PlayMusicByID 通过ID播放音乐
func (MusicCtl) PlayMusicByID(cxt context.Context, req *musicCtl.PlayMusicReq) (res *musicCtl.PlayMusicRes, err error) {
	_ = cxt
	// 检测任务组是否正在运行
	if taskGroupService.RunningTaskGroup != "" {
		// 停止任务组
		taskGroupService.StopTaskGroup()
	}
	// 通过音乐ID 获取音乐文件名和标题
	musicInfo := dao.GetMusicResourceByID(req.Id)
	// 解析文件格式
	mediaService.Play(musicInfo.MusicFile, false, false)

	mediaService.MusicName = musicInfo.Title
	res = &musicCtl.PlayMusicRes{
		Message: "success",
	}
	return
}

// PauseMusic 暂停音乐
func (MusicCtl) PauseMusic(cxt context.Context, req *musicCtl.PauseMusicReq) (res *musicCtl.PauseMusicRes, err error) {
	_ = cxt
	_ = req
	mediaService.Pause()
	res = &musicCtl.PauseMusicRes{
		Message: "success",
	}
	return
}

// ResumeMusic 继续播放音乐
func (MusicCtl) ResumeMusic(cxt context.Context, req *musicCtl.ResumeMusicReq) (res *musicCtl.ResumeMusicRes, err error) {
	_ = cxt
	_ = req
	mediaService.Resume()
	res = &musicCtl.ResumeMusicRes{
		Message: "success",
	}
	return
}

// StopMusic 停止播放音乐
func (MusicCtl) StopMusic(cxt context.Context, req *musicCtl.StopMusicReq) (res *musicCtl.StopMusicRes, err error) {
	_ = cxt
	_ = req
	mediaService.Stop()
	res = &musicCtl.StopMusicRes{
		Message: "success",
	}
	return
}

// EnableLoop 开启循环播放
func (MusicCtl) EnableLoop(cxt context.Context, req *musicCtl.EnableLoopReq) (res *musicCtl.EnableLoopRes, err error) {
	_ = cxt
	_ = req
	mediaService.EnableLoop()
	res = &musicCtl.EnableLoopRes{
		Message: "success",
	}
	return
}

// DisableLoop 关闭循环播放
func (MusicCtl) DisableLoop(cxt context.Context, req *musicCtl.DisableLoopReq) (res *musicCtl.DisableLoopRes, err error) {
	_ = cxt
	_ = req
	mediaService.DisableLoop()
	res = &musicCtl.DisableLoopRes{
		Message: "success",
	}
	return
}

// MusicInfo 获取当前播放音乐信息
func (MusicCtl) MusicInfo(cxt context.Context, req *musicCtl.MusicInfoReq) (res *musicCtl.MusicInfoRes, err error) {
	_ = req
	ws, _ := g.RequestFromCtx(cxt).WebSocket()
	for {
		position := 0
		length := 0
		if mediaService.Streamer != nil {
			position = mediaService.Streamer.Position() / mediaService.Format.SampleRate.N(time.Second)
			length = mediaService.Streamer.Len() / mediaService.Format.SampleRate.N(time.Second)
		}
		_ = ws.WriteJSON(g.Map{
			"position":  position,
			"length":    length,
			"loop":      mediaService.Loop,
			"pause":     mediaService.Ctrl.Paused,
			"musicName": mediaService.MusicName,
		})
		// 每隔700ms发送一次
		time.Sleep(time.Millisecond * 700)
	}
}
