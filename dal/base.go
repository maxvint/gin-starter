package dal

import (
	"github.com/yuwenhui/gin-starter/config"
	_ "github.com/yuwenhui/gin-starter/model"
	"github.com/yuwenhui/gin-starter/proto"
)

type ListModel []interface{}

func Pagination(list ListModel) ([]*proto.Pagination, error) {
	err := config.DBConfig.Where("status = 1").Limit(10).Find(&list).Error

	if err != nil {
		return nil, err
	}

	//results := make([]*proto.Pagination, 0, len(*list))
	// return results, err
}
