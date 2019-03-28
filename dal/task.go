package dal

import (
	"github.com/yuwenhui/gin-starter/config"
	"github.com/yuwenhui/gin-starter/model"
	"github.com/yuwenhui/gin-starter/proto"
)

// GetTasks 获取任务列表
func GetTasks() ([]*proto.TaskItem, error) {
	var list []*model.Task
	err := config.DBConfig.Where("status = 1").Limit(10).Find(&list).Error

	// var test = Pagination()

	if err != nil {
		return nil, err
	}

	results := make([]*proto.TaskItem, 0, len(list))

	for _, item := range list {
		result := &proto.TaskItem{
			UserId:  item.UserId,
			Title:   item.Title,
			Summary: item.Summary,
		}
		results = append(results, result)
	}

	return results, nil
}

func GetTaskList() ([]*proto.TaskItem, error) {
	var list []*model.Task
	result, err := Pagination(list)
}
