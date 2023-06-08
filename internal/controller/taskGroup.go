package controller

import (
	"broadcast_back_end/api/taskGroup"
	"broadcast_back_end/internal/dao"
	"broadcast_back_end/internal/service/mediaService"
	"broadcast_back_end/internal/service/taskGroupService"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type TaskGroup struct{}

// GetTaskGroupList 获取任务组列表
func (TaskGroup) GetTaskGroupList(ctx context.Context, req *taskGroup.GetTaskGroupListReq) (res *taskGroup.GetTaskGroupListRes, err error) {
	_ = req
	g.RequestFromCtx(ctx).Response.WriteJson(dao.GetTaskGroupList())
	return
}

// AddTaskGroup 添加任务组
func (TaskGroup) AddTaskGroup(ctx context.Context, req *taskGroup.AddTaskGroupReq) (res *taskGroup.AddTaskGroupRes, err error) {
	_ = ctx
	dao.AddTaskGroup(req.Name)
	res = &taskGroup.AddTaskGroupRes{
		Message: "success",
	}
	return
}

// DeleteTaskGroup 删除任务组
func (TaskGroup) DeleteTaskGroup(ctx context.Context, req *taskGroup.DeleteTaskGroupReq) (res *taskGroup.DeleteTaskGroupRes, err error) {
	_ = ctx
	dao.DeleteTaskGroup(req.Id)
	res = &taskGroup.DeleteTaskGroupRes{
		Message: "success",
	}
	return
}

// GetTask 获取任务
func (TaskGroup) GetTask(ctx context.Context, req *taskGroup.GetTaskReq) (res *taskGroup.GetTaskRes, err error) {
	_ = req
	g.RequestFromCtx(ctx).Response.WriteJson(dao.GetTaskList(req.GroupID))
	return
}

// GetAllTask 获取所有任务
func (TaskGroup) GetAllTask(ctx context.Context, req *taskGroup.GetAllTaskReq) (res *taskGroup.GetAllTaskRes, err error) {
	_ = req
	g.RequestFromCtx(ctx).Response.WriteJson(dao.GetAllTaskList())
	return
}

// AddTask 添加任务
func (TaskGroup) AddTask(ctx context.Context, req *taskGroup.AddTaskReq) (res *taskGroup.AddTaskRes, err error) {
	_ = ctx
	dao.AddTask(req.GroupID, req.MusicID, req.Loop)
	res = &taskGroup.AddTaskRes{
		Message: "success",
	}
	return
}

// DeleteTask 删除任务
func (TaskGroup) DeleteTask(ctx context.Context, req *taskGroup.DeleteTaskReq) (res *taskGroup.DeleteTaskRes, err error) {
	_ = ctx
	dao.DeleteTask(req.Id)
	res = &taskGroup.DeleteTaskRes{
		Message: "success",
	}
	return
}

// RunTaskGroup 运行任务组
func (TaskGroup) RunTaskGroup(ctx context.Context, req *taskGroup.RunTaskGroupReq) (res *taskGroup.RunTaskGroupRes, err error) {
	_ = ctx
	go taskGroupService.RunTaskGroup(req.Id)
	res = &taskGroup.RunTaskGroupRes{
		Message: "success",
	}
	return
}

// StopTaskGroup 停止任务组
func (TaskGroup) StopTaskGroup(ctx context.Context, req *taskGroup.StopTaskGroupReq) (res *taskGroup.StopTaskGroupRes, err error) {
	_ = ctx
	_ = req
	taskGroupService.StopTaskGroup()
	res = &taskGroup.StopTaskGroupRes{
		Message: "success",
	}
	return
}

// NextTask 下一个任务
func (TaskGroup) NextTask(ctx context.Context, req *taskGroup.NextTaskReq) (res *taskGroup.NextTaskRes, err error) {
	_ = ctx
	_ = req
	// 判断是否有执行中的任务组
	if taskGroupService.RunningTaskGroup == "" {
		res = &taskGroup.NextTaskRes{
			Message: "没有执行中的任务组",
		}
		return
	}
	// 判断是否是待命状态
	if taskGroupService.TaskName == "待命" {
		taskGroupService.Wait <- true
	} else {
		mediaService.NextSong()
	}

	res = &taskGroup.NextTaskRes{
		Message: "success",
	}
	return
}

// TaskGroupInfo 获取当前播放音乐信息 ws
func (TaskGroup) TaskGroupInfo(cxt context.Context, req *taskGroup.InfoReq) (res *taskGroup.InfoRes, err error) {
	_ = req
	ws, _ := g.RequestFromCtx(cxt).WebSocket()
	for {
		_ = ws.WriteJSON(g.Map{
			"taskGroupName": taskGroupService.RunningTaskGroup,
			"currentTask":   taskGroupService.TaskName,
			"nextTask":      taskGroupService.NextTaskName,
		})
		// 每隔700ms发送一次
		time.Sleep(time.Millisecond * 700)
	}
}
