package taskGroupService

import (
	"broadcast_back_end/internal/model/entity"
	"github.com/gogf/gf/v2/container/gqueue"
)

var RunningTaskGroup = ""
var TaskName = ""
var TaskGroupTaskList []entity.TaskTable
var TaskQueue = gqueue.New()
var Wait = make(chan bool)
var NextTaskName = ""
