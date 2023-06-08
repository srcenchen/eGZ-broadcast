package taskGroup

import (
	"broadcast_back_end/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

// GetTaskGroupListReq 获取任务组列表请求
type GetTaskGroupListReq struct {
	g.Meta `method:"GET" summary:"获取任务组列表" tags:"任务策略"`
}

// GetTaskGroupListRes 获取任务组列表响应
type GetTaskGroupListRes struct {
	entity.TaskGroupTable
}

// AddTaskGroupReq 添加任务组请求
type AddTaskGroupReq struct {
	g.Meta `method:"POST" summary:"添加任务组" tags:"任务策略"`
	Name   string `json:"name" v:"required#任务组名不能为空"`
}

// AddTaskGroupRes 添加任务组响应
type AddTaskGroupRes struct {
	Message string `json:"message" example:"success"`
}

// DeleteTaskGroupReq 删除任务组请求
type DeleteTaskGroupReq struct {
	g.Meta `method:"DELETE" summary:"删除任务组" tags:"任务策略"`
	Id     string `json:"id" v:"required#任务组id不能为空"`
}

// DeleteTaskGroupRes 删除任务组响应
type DeleteTaskGroupRes struct {
	Message string `json:"message" example:"success"`
}

// GetTaskReq 获取任务列表请求
type GetTaskReq struct {
	g.Meta  `method:"GET" summary:"获取任务列表" tags:"任务策略"`
	GroupID string `json:"groupId" v:"required#任务组id不能为空"`
}

// GetTaskRes 获取任务列表响应
type GetTaskRes struct {
	entity.TaskTable
}

// GetAllTaskReq 获取所有任务列表请求
type GetAllTaskReq struct {
	g.Meta `method:"GET" summary:"获取所有任务列表" tags:"任务策略"`
}

// GetAllTaskRes 获取所有任务列表响应
type GetAllTaskRes struct {
	entity.TaskTable
}

// AddTaskReq 添加任务请求
type AddTaskReq struct {
	g.Meta  `method:"POST" summary:"添加任务" tags:"任务策略"`
	GroupID string `json:"groupId" v:"required#任务组id不能为空"`
	MusicID string `json:"musicId" v:"required#音乐id不能为空"`
	Loop    bool   `json:"loop" v:"required#循环判断不能为空"`
}

// AddTaskRes 添加任务响应
type AddTaskRes struct {
	Message string `json:"message" example:"success"`
}

// DeleteTaskReq 删除任务请求
type DeleteTaskReq struct {
	g.Meta `method:"DELETE" summary:"删除任务" tags:"任务策略"`
	Id     string `json:"id" v:"required#任务id不能为空"`
}

// DeleteTaskRes 删除任务响应
type DeleteTaskRes struct {
	Message string `json:"message" example:"success"`
}

// RunTaskGroupReq 执行任务组请求
type RunTaskGroupReq struct {
	g.Meta `method:"POST" summary:"执行任务组" tags:"任务策略"`
	Id     string `json:"id" v:"required#任务组id不能为空"`
}

// RunTaskGroupRes 执行任务组响应
type RunTaskGroupRes struct {
	Message string `json:"message" example:"success"`
}

// StopTaskGroupReq 停止任务组请求
type StopTaskGroupReq struct {
	g.Meta `method:"PUT" summary:"停止任务组" tags:"任务策略"`
}

// StopTaskGroupRes 停止任务组响应
type StopTaskGroupRes struct {
	Message string `json:"message" example:"success"`
}

// NextTaskReq 下一个任务请求
type NextTaskReq struct {
	g.Meta `method:"PUT" summary:"下一个任务" tags:"任务策略"`
}

// NextTaskRes 下一个任务响应
type NextTaskRes struct {
	Message string `json:"message" example:"success"`
}

// InfoReq 获取任务组信息请求
type InfoReq struct {
	g.Meta `method:"ALL" summary:"获取任务组信息 ws" tags:"任务策略"`
}

// InfoRes 获取任务组信息响应
type InfoRes struct{}
