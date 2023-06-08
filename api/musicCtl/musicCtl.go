package musicCtl

import "github.com/gogf/gf/v2/frame/g"

// PlayMusicReq 通过id播放音乐请求
type PlayMusicReq struct {
	g.Meta `method:"PUT" summary:"通过id播放音乐" tags:"音乐控制"`
	Id     string `json:"id" v:"required#id不能为空"`
}

// PlayMusicRes 通过id播放音乐响应
type PlayMusicRes struct {
	Message string `json:"message" example:"success"`
}

// PauseMusicReq 暂停音乐请求
type PauseMusicReq struct {
	g.Meta `method:"PUT" summary:"暂停音乐" tags:"音乐控制"`
}

// PauseMusicRes 暂停音乐响应
type PauseMusicRes struct {
	Message string `json:"message" example:"success"`
}

// ResumeMusicReq 继续播放音乐请求
type ResumeMusicReq struct {
	g.Meta `method:"PUT" summary:"继续播放音乐" tags:"音乐控制"`
}

// ResumeMusicRes 继续播放音乐响应
type ResumeMusicRes struct {
	Message string `json:"message" example:"success"`
}

// StopMusicReq 停止音乐请求
type StopMusicReq struct {
	g.Meta `method:"PUT" summary:"停止音乐" tags:"音乐控制"`
}

// StopMusicRes 停止音乐响应
type StopMusicRes struct {
	Message string `json:"message" example:"success"`
}

// EnableLoopReq 开启循环播放请求
type EnableLoopReq struct {
	g.Meta `method:"PUT" summary:"开启循环播放" tags:"音乐控制"`
}

// EnableLoopRes 开启循环播放响应
type EnableLoopRes struct {
	Message string `json:"message" example:"success"`
}

// DisableLoopReq 关闭循环播放请求
type DisableLoopReq struct {
	g.Meta `method:"PUT" summary:"关闭循环播放" tags:"音乐控制"`
}

// DisableLoopRes 关闭循环播放响应
type DisableLoopRes struct {
	Message string `json:"message" example:"success"`
}

// MusicInfoReq 获取音乐信息请求
type MusicInfoReq struct {
	g.Meta `method:"ALL" summary:"获取音乐信息 ws" tags:"音乐状态"`
}

// MusicInfoRes 获取音乐信息响应
type MusicInfoRes struct{}
