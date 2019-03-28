package config

import (
	"github.com/jinzhu/gorm"
	"github.com/yuwenhui/gopkg/goutil"
	"github.com/yuwenhui/gopkg/logs"
)

var (
	ApiConfig AppConfig
	DBConfig  *gorm.DB
)

func init() {
	err := goutil.LoadConfigFor(&ApiConfig, "app.yaml")
	if err != nil {
		logs.Panicf("Load App Config error: %v", err)
	}

	DBConfig, err = ApiConfig.DBAddress.DB()
	if err != nil {
		logs.Panicf("DBAddress error: %v", err)
	}

	// Redis = APIConfig.RedisAPI.RedisClient()
}
