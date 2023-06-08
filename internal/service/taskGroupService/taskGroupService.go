package taskGroupService

import (
	"broadcast_back_end/internal/dao"
	"broadcast_back_end/internal/model/entity"
	"broadcast_back_end/internal/service/mediaService"
	"github.com/gogf/gf/v2/container/gqueue"
	"github.com/gogf/gf/v2/util/gconv"
)

// RunTaskGroup 执行任务组
func RunTaskGroup(groupId string) {
	TaskQueue = gqueue.New()
	// 获取详细任务
	TaskGroupTaskList = dao.GetTaskList(groupId)
	RunningTaskGroup = dao.GetTaskGroup(groupId).Name
	// 历遍任务
	for _, task := range TaskGroupTaskList {
		// 压入队列
		TaskQueue.Push(task)
	}
	var cursor = 0
	// 执行队列
	for TaskQueue.Len() > 0 {
		cursor++
		// 执行任务
		task := TaskQueue.Pop()
		// 检测是否是最后一个任务
		if cursor == len(TaskGroupTaskList) {
			NextTaskName = "结束任务组"
		} else {
			NextTaskName = TaskGroupTaskList[cursor].MusicName
		}
		RunTask(gconv.String(task.(entity.TaskTable).MusicID), task.(entity.TaskTable).Loop)
	}
	// 任务执行完毕
	RunningTaskGroup = ""
}

// RunTask 执行任务
func RunTask(musicId string, loop bool) {
	if musicId == "-1" { // 待命状态
		TaskName = "待命"
		<-Wait
	} else {
		// 根据musicID获取音乐信息
		var musicInfo entity.MusicResourceTable
		entity.GetDatabase().Find(&entity.MusicResourceTable{}, musicId).First(&musicInfo)
		// 播放音乐
		mediaService.MusicName = musicInfo.Title
		TaskName = musicInfo.Title
		mediaService.Play(musicInfo.MusicFile, true, loop)
		mediaService.Stop()
	}
}

// StopTaskGroup 停止任务组
func StopTaskGroup() {
	RunningTaskGroup = ""
	// 停止播放
	mediaService.Stop()
	// 清空队列
	TaskQueue.Close()
}
