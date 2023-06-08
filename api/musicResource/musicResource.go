package musicResource

import (
	"broadcast_back_end/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// MusicListReq 获取音乐列表请求
type MusicListReq struct {
	g.Meta `method:"GET" summary:"获取音乐列表" tags:"音乐资源"`
}

// MusicListRes 获取音乐列表响应
type MusicListRes struct {
	entity.MusicResourceTable
}

// UploadMusicReq 上传音乐文件请求
type UploadMusicReq struct {
	g.Meta `method:"POST" summary:"上传音乐文件" tags:"音乐资源"`
	File   *ghttp.UploadFile `json:"file" v:"required#文件不能为空"`
}

// UploadMusicRes 上传音乐文件响应
type UploadMusicRes struct {
	FileName string `json:"fileName" example:"test.mp3"`
}

// AddMusicReq 添加音乐资源请求
type AddMusicReq struct {
	g.Meta   `method:"POST" summary:"添加音乐资源" tags:"音乐资源"`
	Title    string `json:"title" v:"required#标题不能为空"`
	FileName string `json:"fileName" v:"required#文件名不能为空"`
}

// AddMusicRes 添加音乐资源响应
type AddMusicRes struct {
	Message string `json:"message" example:"success"`
}

// DeleteMusicReq 删除音乐资源请求
type DeleteMusicReq struct {
	g.Meta `method:"DELETE" summary:"删除音乐资源" tags:"音乐资源"`
	Id     string `json:"id" v:"required#id不能为空"`
}

// DeleteMusicRes 删除音乐资源响应
type DeleteMusicRes struct {
	Message string `json:"message" example:"success"`
}
