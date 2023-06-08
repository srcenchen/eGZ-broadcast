package dao

import (
	"broadcast_back_end/internal/model/entity"
	"os"
)

// GetMusicResourceList 获取音乐资源列表
func GetMusicResourceList() []entity.MusicResourceTable {
	var musicList []entity.MusicResourceTable
	entity.GetDatabase().Find(&musicList)
	return musicList
}

// AddMusicResource 添加音乐资源
func AddMusicResource(title string, fileName string) {
	musicResource := entity.MusicResourceTable{
		Title:     title,
		MusicFile: fileName,
	}
	entity.GetDatabase().Create(&musicResource)
}

// DeleteMusicResource 删除音乐资源
func DeleteMusicResource(id string) {
	// 删除任务策略中的音乐
	entity.GetDatabase().Delete(&entity.TaskTable{}, "music_id = ?", id)
	// 获取音乐文件名
	musicInfo := GetMusicResourceByID(id)
	// 删除音乐文件
	_ = os.Remove("./resource/music/" + musicInfo.MusicFile)
	entity.GetDatabase().Delete(&entity.MusicResourceTable{}, id)
}

// GetMusicResourceByID 通过ID获取音乐资源
func GetMusicResourceByID(id string) entity.MusicResourceTable {
	var musicResource entity.MusicResourceTable
	entity.GetDatabase().First(&musicResource, id)
	return musicResource
}
