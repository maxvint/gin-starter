package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yuwenhui/gin-starter/model"
)

type Pagination struct {
	Page     int
	PageSize int
}

func (p *Pagination) Pagination(c *gin.Context, *model) (*gorm.DB,
	error) {
}
