package controller

import (
	"github.com/yuwenhui/gin-starter/dal"
	"github.com/yuwenhui/gin-starter/proto"
	"github.com/yuwenhui/gopkg/apiutil"
	"github.com/yuwenhui/gopkg/errno"
)

func TaskList(c *apiutil.RequestContext) (apiutil.Message, *errno.ErrNo) {
	tasks, err := dal.GetTasks()

	if err != nil {
		return nil, errno.DBError
	}

	return &proto.TaskResponse{
		Tasks: tasks,
	}, nil
}
