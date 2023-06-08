package dao

import (
	"broadcast_back_end/internal/model/entity"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetTaskGroupList 获取任务组列表
func GetTaskGroupList() []entity.TaskGroupTable {
	var taskGroupList []entity.TaskGroupTable
	entity.GetDatabase().Find(&taskGroupList)
	return taskGroupList
}

// AddTaskGroup 添加任务组
func AddTaskGroup(name string) {
	taskGroup := entity.TaskGroupTable{
		Name: name,
	}
	entity.GetDatabase().Create(&taskGroup)
}

// DeleteTaskGroup 删除任务组
func DeleteTaskGroup(id string) {
	entity.GetDatabase().Delete(&entity.TaskGroupTable{}, id)
	entity.GetDatabase().Delete(&entity.TaskTable{}, "group_id = ?", id)
}

// GetTaskList 获取任务列表
func GetTaskList(groupID string) []entity.TaskTable {
	var taskList []entity.TaskTable
	entity.GetDatabase().Where("group_id = ?", groupID).Find(&taskList)
	return taskList
}

// GetTaskGroup 获取任务组
func GetTaskGroup(groupID string) entity.TaskGroupTable {
	var taskGroup entity.TaskGroupTable
	entity.GetDatabase().Find(&entity.TaskGroupTable{}, groupID).First(&taskGroup)
	return taskGroup
}

// GetAllTaskList 获取所有任务列表
func GetAllTaskList() []entity.TaskTable {
	var taskList []entity.TaskTable
	entity.GetDatabase().Find(&taskList)
	return taskList
}

// AddTask 添加任务
func AddTask(groupID string, musicID string, loop bool) {
	// 根据musicID获取音乐信息
	var musicName = ""
	if musicID == "-1" {
		musicName = "待命"
	} else {
		var musicInfo entity.MusicResourceTable
		entity.GetDatabase().Find(&entity.MusicResourceTable{}, musicID).First(&musicInfo)
		musicName = musicInfo.Title
	}
	task := entity.TaskTable{
		GroupID:   gconv.Int(groupID),
		MusicID:   gconv.Int(musicID),
		MusicName: musicName,
		Loop:      loop,
	}
	entity.GetDatabase().Create(&task)
}

// DeleteTask 删除任务
func DeleteTask(id string) {
	entity.GetDatabase().Delete(&entity.TaskTable{}, id)
}
